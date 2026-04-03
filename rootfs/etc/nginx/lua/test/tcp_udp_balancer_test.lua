local cjson = require("cjson.safe")
local util = require("util")

local tcp_udp_balancer, expected_implementations, backends
local original_ngx = ngx

local function reset_ngx()
  _G.ngx = original_ngx
  _G.ngx.ctx.balancer = nil
end

local function reset_balancer()
  package.loaded["tcp_udp_balancer"] = nil
  tcp_udp_balancer = require("tcp_udp_balancer")
end

local function mock_ngx(mock, after_mock_set)
  local _ngx = mock
  setmetatable(_ngx, { __index = ngx })
  _G.ngx = _ngx

  if after_mock_set then
    after_mock_set()
  end

  reset_balancer()
end

local function reset_expected_implementations()
  expected_implementations = {
    ["tcp-app-1"] = package.loaded["balancer.round_robin"],
    ["tcp-app-2"] = package.loaded["balancer.round_robin"],
    ["tcp-app-3"] = package.loaded["balancer.round_robin"],
  }
end

local function reset_backends()
  backends = {
    {
      name = "tcp-app-1", port = "5353",
      endpoints = {
        { address = "10.184.7.40", port = "5353", maxFails = 0, failTimeout = 0 },
        { address = "10.184.97.100", port = "5353", maxFails = 0, failTimeout = 0 },
      },
    },
    {
      name = "tcp-app-2",
      ["load-balance"] = "round_robin",
      endpoints = {
        { address = "10.184.7.40", port = "8080", maxFails = 0, failTimeout = 0 },
      },
    },
    {
      name = "tcp-app-3",
      ["load-balance"] = "ewma",
      endpoints = {
        { address = "10.184.7.40", port = "8080", maxFails = 0, failTimeout = 0 },
      },
    },
  }
end

describe("TCP/UDP Balancer", function()
  before_each(function()
    reset_balancer()
    reset_expected_implementations()
    reset_backends()
  end)

  after_each(function()
    reset_ngx()
  end)

  describe("get_implementation()", function()
    it("returns round_robin when no load-balance is specified", function()
      local backend = { name = "test-backend" }
      assert.equal(package.loaded["balancer.round_robin"],
        tcp_udp_balancer.get_implementation(backend))
    end)

    it("returns round_robin when load-balance is round_robin", function()
      local backend = { name = "test-backend", ["load-balance"] = "round_robin" }
      assert.equal(package.loaded["balancer.round_robin"],
        tcp_udp_balancer.get_implementation(backend))
    end)

    it("falls back to round_robin for unsupported algorithm", function()
      local backend = { name = "test-backend", ["load-balance"] = "ewma" }
      assert.equal(package.loaded["balancer.round_robin"],
        tcp_udp_balancer.get_implementation(backend))
    end)

    it("falls back to round_robin for unknown algorithm and logs warning", function()
      local backend = { name = "test-backend", ["load-balance"] = "invalid_alg" }
      local s = spy.on(ngx, "log")
      local implementation = tcp_udp_balancer.get_implementation(backend)
      assert.equal(package.loaded["balancer.round_robin"], implementation)
      assert.spy(s).was_called_with(ngx.WARN,
        "invalid_alg is not supported, falling back to round_robin")
    end)
  end)

  describe("sync_backend()", function()
    local backend, implementation

    before_each(function()
      backend = backends[1]
      implementation = expected_implementations[backend.name]
    end)

    it("initializes balancer for given backend", function()
      local s = spy.on(implementation, "new")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.spy(s).was_called_with(implementation, backend)
    end)

    it("does not create balancer when endpoints are empty", function()
      backend = { name = "empty-backend", endpoints = {} }
      local s = spy.on(implementation, "new")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.spy(s).was_not_called()
    end)

    it("does not create balancer when endpoints are nil", function()
      backend = { name = "nil-endpoints-backend", endpoints = nil }
      local s = spy.on(implementation, "new")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.spy(s).was_not_called()
    end)

    it("calls sync on existing balancer when load balancing config does not change", function()
      local mock_instance = { sync = function(...) end }
      setmetatable(mock_instance, implementation)
      implementation.new = function(self, backend) return mock_instance end
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)

      stub(mock_instance, "sync")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.stub(mock_instance.sync).was_called_with(mock_instance, backend)
    end)

    it("logs warning and falls back when unsupported load-balance is specified", function()
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)

      backend["load-balance"] = "ewma"

      local s_ngx_log = spy.on(ngx, "log")
      local s = spy.on(implementation, "new")

      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.spy(s_ngx_log).was_called_with(ngx.WARN,
        "ewma is not supported, falling back to round_robin")
      assert.spy(s).was_not_called()
    end)

    it("wraps IPv6 addresses into square brackets on sync", function()
      local backend = {
        name = "ipv6-backend",
        endpoints = {
          { address = "::1", port = "8080", maxFails = 0, failTimeout = 0 },
          { address = "192.168.1.1", port = "8080", maxFails = 0, failTimeout = 0 },
        }
      }
      local expected_backend = {
        name = "ipv6-backend",
        endpoints = {
          { address = "[::1]", port = "8080", maxFails = 0, failTimeout = 0 },
          { address = "192.168.1.1", port = "8080", maxFails = 0, failTimeout = 0 },
        }
      }

      local mock_instance = { sync = function(backend) end }
      setmetatable(mock_instance, implementation)
      implementation.new = function(self, backend) return mock_instance end
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(util.deepcopy(backend)) end)

      stub(mock_instance, "sync")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(util.deepcopy(backend)) end)
      assert.stub(mock_instance.sync).was_called_with(mock_instance, expected_backend)
    end)

    it("logs info when syncing a backend", function()
      local s = spy.on(ngx, "log")
      assert.has_no.errors(function() tcp_udp_balancer.sync_backend(backend) end)
      assert.spy(s).was_called_with(ngx.INFO, "sync tcp/udp backend: ", backend.name)
    end)
  end)

  describe("init_worker()", function()
    it("does not throw errors", function()
      mock_ngx({ var = {}, ctx = {} }, function()
        ngx.shared.tcp_udp_configuration_data:set("backends", cjson.encode(backends))
        ngx.shared.tcp_udp_configuration_data:set("raw_backends_last_synced_at", 1)
      end)
      assert.has_no.errors(function() tcp_udp_balancer.init_worker() end)
    end)
  end)
end)
