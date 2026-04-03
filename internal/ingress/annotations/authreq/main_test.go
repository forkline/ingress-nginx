/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package authreq

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		Service: &networking.IngressServiceBackend{
			Name: "default-backend",
			Port: networking.ServiceBackendPort{
				Number: 80,
			},
		},
	}

	return &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &networking.IngressBackend{
				Service: &networking.IngressServiceBackend{
					Name: "default-backend",
					Port: networking.ServiceBackendPort{
						Number: 80,
					},
				},
			},
			Rules: []networking.IngressRule{
				{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:    "/foo",
									Backend: defaultBackend,
								},
							},
						},
					},
				},
			},
		},
	}
}

func boolToString(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func TestAnnotations(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	ing.SetAnnotations(data)

	tests := []struct {
		title                  string
		url                    string
		signinURL              string
		signinURLRedirectParam string
		method                 string
		requestRedirect        string
		authSnippet            string
		authCacheKey           string
		authAlwaysSetCookie    bool
		expErr                 bool
	}{
		{"empty", "", "", "", "", "", "", "", false, true},
		{"no scheme", "bar", "bar", "", "", "", "", "", false, true},
		{"invalid host", "http://", "http://", "", "", "", "", "", false, true},
		{"invalid host (multiple dots)", "http://foo..bar.com", "http://foo..bar.com", "", "", "", "", "", false, true},
		{"valid URL", "http://bar.foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://bar.foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "", "", "", "", false, false},
		{"valid URL - send body", "http://foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "POST", "", "", "", false, false},
		{"valid URL - send body", "http://foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "GET", "", "", "", false, false},
		{"valid URL - request redirect", "http://foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "GET", "http://foo.com/redirect-me", "", "", false, false},
		{"auth snippet", "http://foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "", "", "proxy_set_header My-Custom-Header 42;", "", false, false},
		{"auth cache ", "http://foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "", "", "", "", "$foo$bar", false, false},
		{"redirect param", "http://bar.foo.com/external-auth/auth?allowed_groups=snow-group,rain-group", "http://bar.foo.com/external-auth/start?rd=https://$host$escaped_request_uri", "origUrl", "", "", "", "", true, false},
	}

	for _, test := range tests {
		data[parser.GetAnnotationWithPrefix("auth-url")] = test.url
		data[parser.GetAnnotationWithPrefix("auth-signin")] = test.signinURL
		data[parser.GetAnnotationWithPrefix("auth-signin-redirect-param")] = test.signinURLRedirectParam
		data[parser.GetAnnotationWithPrefix("auth-method")] = test.method
		data[parser.GetAnnotationWithPrefix("auth-request-redirect")] = test.requestRedirect
		data[parser.GetAnnotationWithPrefix("auth-snippet")] = test.authSnippet
		data[parser.GetAnnotationWithPrefix("auth-cache-key")] = test.authCacheKey
		data[parser.GetAnnotationWithPrefix("auth-always-set-cookie")] = boolToString(test.authAlwaysSetCookie)

		i, err := NewParser(&resolver.Mock{}).Parse(ing)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but returned nil", test.title)
			}
			continue
		}
		if err != nil {
			t.Errorf("%v: unexpected error: %v", test.title, err)
		}

		u, ok := i.(*Config)
		if !ok {
			t.Errorf("%v: expected an External type", test.title)
		}
		if u.URL != test.url {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.url, u.URL)
		}
		if u.SigninURL != test.signinURL {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.signinURL, u.SigninURL)
		}
		if u.SigninURLRedirectParam != test.signinURLRedirectParam {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.signinURLRedirectParam, u.SigninURLRedirectParam)
		}
		if u.Method != test.method {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.method, u.Method)
		}
		if u.RequestRedirect != test.requestRedirect {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.requestRedirect, u.RequestRedirect)
		}
		if u.AuthSnippet != test.authSnippet {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.authSnippet, u.AuthSnippet)
		}
		if u.AuthCacheKey != test.authCacheKey {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.authCacheKey, u.AuthCacheKey)
		}

		if u.AlwaysSetCookie != test.authAlwaysSetCookie {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.authAlwaysSetCookie, u.AlwaysSetCookie)
		}
	}
}

func TestHeaderAnnotations(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	ing.SetAnnotations(data)

	tests := []struct {
		title         string
		url           string
		headers       string
		parsedHeaders []string
		expErr        bool
	}{
		{"single header", "http://goog.url", "h1", []string{"h1"}, false},
		{"nothing", "http://goog.url", "", []string{}, false},
		{"spaces", "http://goog.url", "  ", []string{}, false},
		{"two headers", "http://goog.url", "1,2", []string{"1", "2"}, false},
		{"two headers and empty entries", "http://goog.url", ",1,,2,", []string{"1", "2"}, false},
		{"header with spaces", "http://goog.url", "1 2", []string{}, true},
		{"header with other bad symbols", "http://goog.url", "1+2", []string{}, true},
	}

	for _, test := range tests {
		data[parser.GetAnnotationWithPrefix("auth-url")] = test.url
		data[parser.GetAnnotationWithPrefix("auth-response-headers")] = test.headers
		data[parser.GetAnnotationWithPrefix("auth-method")] = "GET"

		i, err := NewParser(&resolver.Mock{}).Parse(ing)
		if test.expErr {
			if err == nil {
				t.Errorf("%v expected error but retuned nil", test.title)
			}
			continue
		}
		if err != nil {
			t.Errorf("no error was expected but %v happened in %s", err, test.title)
		}
		u, ok := i.(*Config)
		if !ok {
			t.Errorf("%v: expected an External type", test.title)
			continue
		}

		if !reflect.DeepEqual(u.ResponseHeaders, test.parsedHeaders) {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.headers, u.ResponseHeaders)
		}
	}
}

func TestCacheDurationAnnotations(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	ing.SetAnnotations(data)

	tests := []struct {
		title          string
		url            string
		duration       string
		parsedDuration []string
		expErr         bool
	}{
		{"nothing", "http://goog.url", "", []string{DefaultCacheDuration}, false},
		{"spaces", "http://goog.url", "  ", []string{DefaultCacheDuration}, false},
		{"one duration", "http://goog.url", "5m", []string{"5m"}, false},
		{"two durations", "http://goog.url", "200 202 10m, 401 5m", []string{"200 202 10m", "401 5m"}, false},
		{"two durations and empty entries", "http://goog.url", ",5m,,401 10m,", []string{"5m", "401 10m"}, false},
		{"only status code provided", "http://goog.url", "200", []string{DefaultCacheDuration}, true},
		{"mixed valid/invalid", "http://goog.url", "5m, xaxax", []string{DefaultCacheDuration}, true},
		{"code after duration", "http://goog.url", "5m 200", []string{DefaultCacheDuration}, true},
	}

	for _, test := range tests {
		data[parser.GetAnnotationWithPrefix("auth-url")] = test.url
		data[parser.GetAnnotationWithPrefix("auth-cache-duration")] = test.duration

		i, err := NewParser(&resolver.Mock{}).Parse(ing)
		if test.expErr {
			if err == nil {
				t.Errorf("expected error but retuned nil")
			}
			continue
		}

		u, ok := i.(*Config)
		if !ok {
			t.Errorf("%v: expected an External type", test.title)
			continue
		}

		if !reflect.DeepEqual(u.AuthCacheDuration, test.parsedDuration) {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.duration, u.AuthCacheDuration)
		}
	}
}

func TestKeepaliveAnnotations(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	ing.SetAnnotations(data)

	tests := []struct {
		title                string
		url                  string
		keepaliveConnections string
		keepaliveShareVars   string
		keepaliveRequests    string
		keepaliveTimeout     string
		expectedConnections  int
		expectedShareVars    bool
		expectedRequests     int
		expectedTimeout      int
	}{
		{"all set", "http://goog.url", "5", "false", "500", "50", 5, false, 500, 50},
		{"no annotation", "http://goog.url", "", "", "", "", defaultKeepaliveConnections, defaultKeepaliveShareVars, defaultKeepaliveRequests, defaultKeepaliveTimeout},
		{"default for connections", "http://goog.url", "x", "true", "500", "50", defaultKeepaliveConnections, true, 500, 50},
		{"default for requests", "http://goog.url", "5", "x", "dummy", "50", 5, defaultKeepaliveShareVars, defaultKeepaliveRequests, 50},
		{"default for invalid timeout", "http://goog.url", "5", "t", "500", "x", 5, true, 500, defaultKeepaliveTimeout},
		{"variable in host", "http://$host:5000/a/b", "5", "1", "", "", 0, true, defaultKeepaliveRequests, defaultKeepaliveTimeout},
		{"variable in path", "http://goog.url:5000/$path", "5", "t", "", "", 5, true, defaultKeepaliveRequests, defaultKeepaliveTimeout},
		{"negative connections", "http://goog.url", "-2", "f", "", "", 0, false, defaultKeepaliveRequests, defaultKeepaliveTimeout},
		{"negative requests", "http://goog.url", "5", "True", "-1", "", 0, true, -1, defaultKeepaliveTimeout},
		{"negative timeout", "http://goog.url", "5", "0", "", "-1", 0, false, defaultKeepaliveRequests, -1},
		{"negative request and timeout", "http://goog.url", "5", "False", "-2", "-3", 0, false, -2, -3},
	}

	for _, test := range tests {
		data[parser.GetAnnotationWithPrefix("auth-url")] = test.url
		data[parser.GetAnnotationWithPrefix("auth-keepalive")] = test.keepaliveConnections
		data[parser.GetAnnotationWithPrefix("auth-keepalive-share-vars")] = test.keepaliveShareVars
		data[parser.GetAnnotationWithPrefix("auth-keepalive-timeout")] = test.keepaliveTimeout
		data[parser.GetAnnotationWithPrefix("auth-keepalive-requests")] = test.keepaliveRequests

		i, err := NewParser(&resolver.Mock{}).Parse(ing)
		if err != nil {
			t.Errorf("%v: unexpected error: %v", test.title, err)
			continue
		}

		u, ok := i.(*Config)
		if !ok {
			t.Errorf("%v: expected an External type", test.title)
			continue
		}

		if u.URL != test.url {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.url, u.URL)
		}

		if u.KeepaliveConnections != test.expectedConnections {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.expectedConnections, u.KeepaliveConnections)
		}

		if u.KeepaliveShareVars != test.expectedShareVars {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.expectedShareVars, u.KeepaliveShareVars)
		}

		if u.KeepaliveRequests != test.expectedRequests {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.expectedRequests, u.KeepaliveRequests)
		}

		if u.KeepaliveTimeout != test.expectedTimeout {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.expectedTimeout, u.KeepaliveTimeout)
		}
	}
}

func TestParseStringToCacheDurations(t *testing.T) {
	tests := []struct {
		title             string
		duration          string
		expectedDurations []string
		expErr            bool
	}{
		{"empty", "", []string{DefaultCacheDuration}, false},
		{"invalid", ",200,", []string{DefaultCacheDuration}, true},
		{"single", ",200 5m,", []string{"200 5m"}, false},
		{"multiple with duration", ",5m,,401 10m,", []string{"5m", "401 10m"}, false},
		{"multiple durations", "200 202 401 5m, 418 30m", []string{"200 202 401 5m", "418 30m"}, false},
	}

	for _, test := range tests {
		dur, err := ParseStringToCacheDurations(test.duration)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but nil was returned", test.title)
			}
			continue
		}

		if !reflect.DeepEqual(dur, test.expectedDurations) {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.expectedDurations, dur)
		}
	}
}

func TestProxySetHeaders(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	ing.SetAnnotations(data)

	tests := []struct {
		title   string
		url     string
		headers map[string]string
		expErr  bool
	}{
		{"single header", "http://goog.url", map[string]string{"header": "h1"}, false},
		{"no header map", "http://goog.url", nil, true},
		{"header with spaces", "http://goog.url", map[string]string{"header": "bad value"}, false},
		{"header with other bad symbols", "http://goog.url", map[string]string{"header": "bad+value"}, false},
	}

	for _, test := range tests {
		data[parser.GetAnnotationWithPrefix("auth-url")] = test.url
		data[parser.GetAnnotationWithPrefix("auth-proxy-set-headers")] = "proxy-headers-map"
		data[parser.GetAnnotationWithPrefix("auth-method")] = "GET"

		configMapResolver := &resolver.Mock{
			ConfigMaps: map[string]*api.ConfigMap{},
		}

		if test.headers != nil {
			configMapResolver.ConfigMaps["proxy-headers-map"] = &api.ConfigMap{Data: test.headers}
		}

		i, err := NewParser(configMapResolver).Parse(ing)
		if test.expErr {
			if err == nil {
				t.Errorf("expected error but retuned nil")
			}
			continue
		}

		u, ok := i.(*Config)
		if !ok {
			t.Errorf("%v: expected an External type", test.title)
			continue
		}

		if !reflect.DeepEqual(u.ProxySetHeaders, test.headers) {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.headers, u.ProxySetHeaders)
		}
	}
}

func TestValidCacheDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid single status and duration", "200 5m", true},
		{"valid multiple statuses and duration", "200 202 401 5m", true},
		{"valid duration only", "5m", true},
		{"valid compound duration", "200 1h 30m", true},
		{"valid milliseconds", "200 500ms", true},
		{"valid seconds", "200 30s", true},
		{"valid hours", "200 2h", true},
		{"valid days", "200 1d", true},
		{"valid weeks", "200 1w", true},
		{"valid months", "200 1M", true},
		{"valid years", "200 1y", true},
		{"only status code", "200", false},
		{"empty string", "", false},
		{"code after duration", "5m 200", false},
		{"invalid duration unit", "200 5x", false},
		{"non-standard status code passes", "99 5m", true},
		{"four digit code ignores code validates duration", "2000 5m", true},
		{"random text", "something", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ValidCacheDuration(tt.input))
		})
	}
}

func TestParseStringToCacheDurationsExtended(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		expErr   bool
	}{
		{"whitespace only", "   ", []string{DefaultCacheDuration}, false},
		{"single valid duration", "10m", []string{"10m"}, false},
		{"multiple codes one duration", "200 202 10m", []string{"200 202 10m"}, false},
		{"comma separated valid", "200 5m, 401 10m", []string{"200 5m", "401 10m"}, false},
		{"trailing comma valid", "200 5m,", []string{"200 5m"}, false},
		{"leading comma valid", ",200 5m", []string{"200 5m"}, false},
		{"invalid duration returns default", "badvalue", []string{DefaultCacheDuration}, true},
		{"code only returns error", "200", []string{DefaultCacheDuration}, true},
		{"compound duration valid", "200 1h 30m", []string{"200 1h 30m"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur, err := ParseStringToCacheDurations(tt.input)
			if tt.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, dur)
		})
	}
}

func TestValidMethod(t *testing.T) {
	tests := []struct {
		name   string
		method string
		valid  bool
	}{
		{"GET", "GET", true},
		{"HEAD", "HEAD", true},
		{"POST", "POST", true},
		{"PUT", "PUT", true},
		{"PATCH", "PATCH", true},
		{"DELETE", "DELETE", true},
		{"CONNECT", "CONNECT", true},
		{"OPTIONS", "OPTIONS", true},
		{"TRACE", "TRACE", true},
		{"lowercase get", "get", false},
		{"empty string", "", false},
		{"invalid method", "INVALID", false},
		{"method with space", "GE T", false},
		{"mixed case", "Get", false},
		{"partial match", "GETMORE", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, ValidMethod(tt.method))
		})
	}
}

func TestValidHeader(t *testing.T) {
	tests := []struct {
		name   string
		header string
		valid  bool
	}{
		{"simple header", "X-Custom-Header", true},
		{"with digits", "X-Header123", true},
		{"with underscore", "X_Custom_Header", true},
		{"single char", "A", true},
		{"all digits", "123", true},
		{"empty string", "", false},
		{"with space", "X Header", false},
		{"with colon", "X:Header", false},
		{"with at sign", "X@Header", false},
		{"with hash", "X#Header", false},
		{"with slash", "X/Header", false},
		{"with dot", "X.Header", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, ValidHeader(tt.header))
		})
	}
}

func TestAuthreqConfigEqual(t *testing.T) {
	tests := []struct {
		name   string
		e1     *Config
		e2     *Config
		expect bool
	}{
		{
			"both nil",
			nil,
			nil,
			true,
		},
		{
			"one nil",
			&Config{},
			nil,
			false,
		},
		{
			"same identity",
			nil,
			nil,
			true,
		},
		{
			"equal configs",
			&Config{
				URL:                    "http://foo.com",
				Host:                   "foo.com",
				SigninURL:              "http://foo.com/signin",
				SigninURLRedirectParam: "rd",
				Method:                 "GET",
				ResponseHeaders:        []string{"X-Auth", "X-User"},
				RequestRedirect:        "http://redirect.com",
				AuthSnippet:            "snippet",
				AuthCacheKey:           "$uri",
				AuthCacheDuration:      []string{"200 5m"},
				KeepaliveConnections:   5,
				KeepaliveShareVars:     true,
				KeepaliveRequests:      1000,
				KeepaliveTimeout:       60,
				ProxySetHeaders:        map[string]string{"X-Header": "value"},
				AlwaysSetCookie:        true,
			},
			&Config{
				URL:                    "http://foo.com",
				Host:                   "foo.com",
				SigninURL:              "http://foo.com/signin",
				SigninURLRedirectParam: "rd",
				Method:                 "GET",
				ResponseHeaders:        []string{"X-User", "X-Auth"},
				RequestRedirect:        "http://redirect.com",
				AuthSnippet:            "snippet",
				AuthCacheKey:           "$uri",
				AuthCacheDuration:      []string{"200 5m"},
				KeepaliveConnections:   5,
				KeepaliveShareVars:     true,
				KeepaliveRequests:      1000,
				KeepaliveTimeout:       60,
				ProxySetHeaders:        map[string]string{"X-Header": "value"},
				AlwaysSetCookie:        true,
			},
			true,
		},
		{
			"different URL",
			&Config{URL: "http://a.com"},
			&Config{URL: "http://b.com"},
			false,
		},
		{
			"different Host",
			&Config{Host: "a.com"},
			&Config{Host: "b.com"},
			false,
		},
		{
			"different SigninURL",
			&Config{SigninURL: "http://a.com"},
			&Config{SigninURL: "http://b.com"},
			false,
		},
		{
			"different SigninURLRedirectParam",
			&Config{SigninURLRedirectParam: "a"},
			&Config{SigninURLRedirectParam: "b"},
			false,
		},
		{
			"different Method",
			&Config{Method: "GET"},
			&Config{Method: "POST"},
			false,
		},
		{
			"different ResponseHeaders",
			&Config{ResponseHeaders: []string{"X-A"}},
			&Config{ResponseHeaders: []string{"X-B"}},
			false,
		},
		{
			"different RequestRedirect",
			&Config{RequestRedirect: "http://a.com"},
			&Config{RequestRedirect: "http://b.com"},
			false,
		},
		{
			"different AuthSnippet",
			&Config{AuthSnippet: "a"},
			&Config{AuthSnippet: "b"},
			false,
		},
		{
			"different AuthCacheKey",
			&Config{AuthCacheKey: "a"},
			&Config{AuthCacheKey: "b"},
			false,
		},
		{
			"different KeepaliveConnections",
			&Config{KeepaliveConnections: 1},
			&Config{KeepaliveConnections: 2},
			false,
		},
		{
			"different KeepaliveShareVars",
			&Config{KeepaliveShareVars: true},
			&Config{KeepaliveShareVars: false},
			false,
		},
		{
			"different KeepaliveRequests",
			&Config{KeepaliveRequests: 1},
			&Config{KeepaliveRequests: 2},
			false,
		},
		{
			"different KeepaliveTimeout",
			&Config{KeepaliveTimeout: 1},
			&Config{KeepaliveTimeout: 2},
			false,
		},
		{
			"different AlwaysSetCookie",
			&Config{AlwaysSetCookie: true},
			&Config{AlwaysSetCookie: false},
			false,
		},
		{
			"different ProxySetHeaders",
			&Config{ProxySetHeaders: map[string]string{"a": "1"}},
			&Config{ProxySetHeaders: map[string]string{"b": "1"}},
			false,
		},
		{
			"different AuthCacheDuration",
			&Config{AuthCacheDuration: []string{"200 5m"}},
			&Config{AuthCacheDuration: []string{"200 10m"}},
			false,
		},
		{
			"response headers order independent",
			&Config{ResponseHeaders: []string{"X-A", "X-B"}},
			&Config{ResponseHeaders: []string{"X-B", "X-A"}},
			true,
		},
		{
			"cache duration order independent",
			&Config{AuthCacheDuration: []string{"200 5m", "401 10m"}},
			&Config{AuthCacheDuration: []string{"401 10m", "200 5m"}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.e1.Equal(tt.e2))
		})
	}
}
