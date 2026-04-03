local lua_ingress = require("lua_ingress")

local unmocked_ngx = _G.ngx

local redirect_uri_result
local redirect_code_result

local function mock_ngx(mock)
  local _ngx = mock
  setmetatable(_ngx, { __index = ngx })
  _G.ngx = _ngx
  package.loaded["lua_ingress"] = nil
  lua_ingress = require("lua_ingress")
  lua_ingress.init_worker()
end

local function reset_ngx()
  _G.ngx = unmocked_ngx
  package.loaded["lua_ingress"] = nil
  lua_ingress = require("lua_ingress")
  lua_ingress.init_worker()
end

local function make_config(overrides)
  local config = {
    use_forwarded_headers = false,
    use_proxy_protocol = false,
    is_ssl_passthrough_enabled = false,
    listen_ports = {
      https = "443",
      ssl_proxy = "442",
    },
    http_redirect_code = 308,
  }
  for k, v in pairs(overrides or {}) do
    config[k] = v
  end
  return config
end

local function make_mock_ngx_var(overrides)
  local var = {
    force_ssl_redirect = "false",
    ssl_redirect = "false",
    force_no_ssl_redirect = "false",
    preserve_trailing_slash = "false",
    use_port_in_redirects = "false",
    scheme = "http",
    http_host = "example.com",
    host = "example.com",
    pass_server_port = "80",
    request_uri = "/path",
  }
  for k, v in pairs(overrides or {}) do
    var[k] = v
  end
  return var
end

local function setup_with_config(config, ngx_vars, redirect_mock)
  redirect_uri_result = nil
  redirect_code_result = nil

  lua_ingress.set_config(config)

  mock_ngx({
    var = make_mock_ngx_var(ngx_vars),
    ctx = {},
    redirect = function(uri, code)
      redirect_uri_result = uri
      redirect_code_result = code
    end,
  })

  lua_ingress.set_config(config)
end

describe("lua_ingress rewrite()", function()
  before_each(function()
    reset_ngx()
  end)

  after_each(function()
    _G.ngx = unmocked_ngx
    package.loaded["lua_ingress"] = nil
    lua_ingress = require("lua_ingress")
    lua_ingress.init_worker()
  end)

  describe("SSL redirect", function()
    it("redirects to HTTPS when force_ssl_redirect is true and scheme is http", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        scheme = "http",
      })

      lua_ingress.rewrite()

      assert.is_not_nil(redirect_uri_result)
      assert.equal("https://example.com/path", redirect_uri_result)
      assert.equal(308, redirect_code_result)
    end)

    it("does not redirect when force_ssl_redirect is false and scheme is http", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "false",
        scheme = "http",
      })

      lua_ingress.rewrite()

      assert.is_nil(redirect_uri_result)
    end)

    it("does not redirect when force_no_ssl_redirect is true", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        ssl_redirect = "true",
        force_no_ssl_redirect = "true",
        scheme = "http",
      })

      lua_ingress.rewrite()

      assert.is_nil(redirect_uri_result)
    end)

    it("does not redirect when scheme is https", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        ssl_redirect = "true",
        scheme = "https",
        pass_server_port = "443",
      })

      lua_ingress.rewrite()

      assert.is_nil(redirect_uri_result)
    end)
  end)

  describe("Forwarded headers", function()
    it("uses x-forwarded-proto as pass_access_scheme when forwarded headers enabled", function()
      setup_with_config(make_config({ use_forwarded_headers = true }), {
        http_x_forwarded_proto = "https",
        http_x_forwarded_port = "443",
      })

      lua_ingress.rewrite()

      assert.equal("https", ngx.var.pass_access_scheme)
      assert.equal("443", ngx.var.pass_server_port)
    end)

    it("uses x-forwarded-host as best_http_host when forwarded headers enabled", function()
      setup_with_config(make_config({ use_forwarded_headers = true }), {
        http_x_forwarded_host = "external.com",
      })

      lua_ingress.rewrite()

      assert.equal("external.com", ngx.var.best_http_host)
    end)

    it("uses first host from x-forwarded-host with comma-separated values", function()
      setup_with_config(make_config({ use_forwarded_headers = true }), {
        http_x_forwarded_host = "first.com, second.com",
      })

      lua_ingress.rewrite()

      assert.equal("first.com", ngx.var.best_http_host)
    end)

    it("does not use forwarded headers when use_forwarded_headers is false", function()
      setup_with_config(make_config({ use_forwarded_headers = false }), {
        http_x_forwarded_proto = "https",
        http_x_forwarded_port = "443",
        http_x_forwarded_host = "forwarded.com",
      })

      lua_ingress.rewrite()

      assert.equal("http", ngx.var.pass_access_scheme)
      assert.equal("example.com", ngx.var.best_http_host)
    end)
  end)

  describe("Proxy protocol", function()
    it("sets pass_access_scheme to https when proxy_protocol_server_port is 443", function()
      setup_with_config(make_config({ use_proxy_protocol = true }), {
        pass_server_port = "443",
        proxy_protocol_server_port = "443",
      })

      lua_ingress.rewrite()

      assert.equal("https", ngx.var.pass_access_scheme)
    end)

    it("does not set pass_access_scheme to https when proxy_protocol_server_port is not 443", function()
      setup_with_config(make_config({ use_proxy_protocol = true }), {
        pass_server_port = "80",
        proxy_protocol_server_port = "80",
      })

      lua_ingress.rewrite()

      assert.equal("http", ngx.var.pass_access_scheme)
    end)
  end)

  describe("Port handling", function()
    it("sets pass_port to 443 when pass_server_port matches https port", function()
      setup_with_config(make_config(), {
        scheme = "https",
        pass_server_port = "443",
      })

      lua_ingress.rewrite()

      assert.equal(443, ngx.var.pass_port)
    end)

    it("sets pass_port to 443 when ssl passthrough enabled and port matches ssl_proxy", function()
      setup_with_config(make_config({ is_ssl_passthrough_enabled = true }), {
        scheme = "https",
        pass_server_port = "442",
      })

      lua_ingress.rewrite()

      assert.equal(443, ngx.var.pass_port)
    end)

    it("uses pass_server_port as pass_port when it does not match https or ssl_proxy", function()
      setup_with_config(make_config(), {
        pass_server_port = "8080",
      })

      lua_ingress.rewrite()

      assert.equal("8080", ngx.var.pass_port)
    end)
  end)

  describe("Trailing slash handling", function()
    it("strips trailing slash when preserve_trailing_slash is false", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        scheme = "http",
        request_uri = "/path/",
        preserve_trailing_slash = "false",
      })

      lua_ingress.rewrite()

      assert.equal("https://example.com/path", redirect_uri_result)
    end)

    it("preserves trailing slash when preserve_trailing_slash is true", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        scheme = "http",
        request_uri = "/path/",
        preserve_trailing_slash = "true",
      })

      lua_ingress.rewrite()

      assert.equal("https://example.com/path/", redirect_uri_result)
    end)
  end)

  describe("Port in redirects", function()
    it("includes port in redirect URI when use_port_in_redirects is true", function()
      local config = make_config({
        listen_ports = {
          https = "8443",
          ssl_proxy = "442",
        },
      })
      setup_with_config(config, {
        force_ssl_redirect = "true",
        scheme = "http",
        use_port_in_redirects = "true",
      })

      lua_ingress.rewrite()

      assert.equal("https://example.com:8443/path", redirect_uri_result)
    end)
  end)

  describe("Host handling", function()
    it("uses http_host as best_http_host when present", function()
      setup_with_config(make_config(), {
        http_host = "example.com:8080",
        pass_server_port = "8080",
      })

      lua_ingress.rewrite()

      assert.equal("example.com:8080", ngx.var.best_http_host)
    end)

    it("uses host as best_http_host when http_host is nil", function()
      setup_with_config(make_config(), {
        http_host = nil,
      })

      lua_ingress.rewrite()

      assert.equal("example.com", ngx.var.best_http_host)
    end)

    it("strips port from host in redirect URL", function()
      setup_with_config(make_config(), {
        force_ssl_redirect = "true",
        scheme = "http",
        http_host = "example.com:8080",
        pass_server_port = "8080",
      })

      lua_ingress.rewrite()

      assert.equal("https://example.com/path", redirect_uri_result)
    end)
  end)
end)
