local cjson = require("cjson")
local configuration = require("tcp_udp_configuration")

local unmocked_ngx = _G.ngx

local function get_backends()
  return {
    {
      name = "my-tcp-backend-1",
      endpoints = { { address = "10.183.7.40", port = "5353", maxFails = 0, failTimeout = 0 } },
    },
    {
      name = "my-udp-backend-2",
      endpoints = {
        { address = "10.184.7.40", port = "53", maxFails = 3, failTimeout = 2 },
        { address = "10.184.7.41", port = "53", maxFails = 2, failTimeout = 1 },
      }
    },
  }
end

local function get_mocked_ngx_env()
  local _ngx = {
    status = ngx.HTTP_OK,
    var = {},
    say = function(msg) end,
    log = function(...) end,
    req = {
      socket = function()
        return {
          receiveuntil = function(self, delim)
            local data = cjson.encode(get_backends())
            local reader_called = false
            return function()
              if reader_called then
                return nil
              end
              reader_called = true
              return data
            end
          end
        }
      end,
    },
  }
  setmetatable(_ngx, {__index = _G.ngx})
  return _ngx
end

describe("TCP/UDP Configuration", function()
  before_each(function()
    _G.ngx = get_mocked_ngx_env()
    package.loaded["tcp_udp_configuration"] = nil
    configuration = require("tcp_udp_configuration")
  end)

  after_each(function()
    _G.ngx = unmocked_ngx
  end)

  describe("get_backends_data()", function()
    it("returns nil when no backends data is stored", function()
      ngx.shared.tcp_udp_configuration_data:set("backends", nil)
      assert.is_nil(configuration.get_backends_data())
    end)

    it("returns stored backends data", function()
      local encoded_backends = cjson.encode(get_backends())
      ngx.shared.tcp_udp_configuration_data:set("backends", encoded_backends)
      assert.equal(encoded_backends, configuration.get_backends_data())
    end)
  end)

  describe("get_raw_backends_last_synced_at()", function()
    it("returns 1 when raw_backends_last_synced_at is nil", function()
      ngx.shared.tcp_udp_configuration_data:set("raw_backends_last_synced_at", nil)
      assert.equal(1, configuration.get_raw_backends_last_synced_at())
    end)

    it("returns stored value when set", function()
      ngx.shared.tcp_udp_configuration_data:set("raw_backends_last_synced_at", 12345)
      assert.equal(12345, configuration.get_raw_backends_last_synced_at())
    end)
  end)

  describe("call()", function()
    it("stores backends data in shared dict on valid input", function()
      assert.has_no.errors(configuration.call)
      local stored = ngx.shared.tcp_udp_configuration_data:get("backends")
      assert.equal(cjson.encode(get_backends()), stored)
    end)

    it("stores raw_backends_last_synced_at timestamp", function()
      assert.has_no.errors(configuration.call)
      local stored = ngx.shared.tcp_udp_configuration_data:get("raw_backends_last_synced_at")
      assert.is_not_nil(stored)
      assert.is_true(stored > 0)
    end)

    it("logs error when socket is not available", function()
      ngx.req.socket = function()
        return nil, "socket error"
      end
      local s = spy.on(ngx, "log")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called_with(ngx.ERR, "failed to get raw req socket: ", "socket error")
    end)

    it("says error when socket is not available", function()
      ngx.req.socket = function()
        return nil, "socket error"
      end
      local s = spy.on(ngx, "say")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called_with("error: ", "socket error")
    end)

    it("logs error when reader fails", function()
      ngx.req.socket = function()
        return {
          receiveuntil = function(self, delim)
            return function()
              return nil, "read error"
            end
          end
        }
      end
      local s = spy.on(ngx, "log")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called_with(ngx.ERR, "failed TCP/UDP dynamic-configuration:", "read error")
    end)

    it("returns when backends is nil", function()
      ngx.req.socket = function()
        return {
          receiveuntil = function(self, delim)
            return function()
              return nil
            end
          end
        }
      end
      assert.has_no.errors(configuration.call)
    end)

    it("returns when backends is empty string", function()
      ngx.req.socket = function()
        return {
          receiveuntil = function(self, delim)
            local called = false
            return function()
              if called then return nil end
              called = true
              return ""
            end
          end
        }
      end
      assert.has_no.errors(configuration.call)
    end)

    it("logs error and returns when backends is invalid JSON", function()
      ngx.req.socket = function()
        return {
          receiveuntil = function(self, delim)
            local called = false
            return function()
              if called then return nil end
              called = true
              return "not-valid-json"
            end
          end
        }
      end
      local s = spy.on(ngx, "log")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called()
    end)

    it("logs error when shared dict set fails", function()
      local original_set = ngx.shared.tcp_udp_configuration_data.set
      ngx.shared.tcp_udp_configuration_data.set = function(self, key, value)
        if key == "backends" then
          return false, "dict error"
        end
        return original_set(self, key, value)
      end

      local s = spy.on(ngx, "log")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called_with(ngx.ERR,
        "dynamic-configuration: error updating configuration: dict error")

      ngx.shared.tcp_udp_configuration_data.set = original_set
    end)

    it("logs error when raw_backends_last_synced_at set fails", function()
      local original_set = ngx.shared.tcp_udp_configuration_data.set
      ngx.shared.tcp_udp_configuration_data.set = function(self, key, value)
        if key == "raw_backends_last_synced_at" then
          return false, "timestamp error"
        end
        return original_set(self, key, value)
      end

      local s = spy.on(ngx, "log")
      assert.has_no.errors(configuration.call)
      assert.spy(s).was_called_with(ngx.ERR,
        "dynamic-configuration: error updating when backends sync, " ..
        "new upstream peers waiting for force syncing: timestamp error")

      ngx.shared.tcp_udp_configuration_data.set = original_set
    end)
  end)
end)
