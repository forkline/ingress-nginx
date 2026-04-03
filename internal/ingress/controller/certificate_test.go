/*
Copyright 2017 The Kubernetes Authors.

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

package controller

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLowerCaseASCII(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"already lowercase", "foobar", "foobar"},
		{"uppercase", "FOOBAR", "foobar"},
		{"mixed case", "FooBar", "foobar"},
		{"empty string", "", ""},
		{"with digits", "Foo123Bar", "foo123bar"},
		{"with hyphens", "FOO-BAR", "foo-bar"},
		{"with dots", "Example.COM", "example.com"},
		{"single char uppercase", "A", "a"},
		{"single char lowercase", "a", "a"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toLowerCaseASCII(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMatchHostnames(t *testing.T) {
	testCases := []struct {
		name     string
		pattern  string
		host     string
		expected bool
	}{
		{"exact match", "example.com", "example.com", true},
		{"no match", "example.com", "other.com", false},
		{"wildcard match", "*.example.com", "foo.example.com", true},
		{"wildcard no match different number of parts", "*.example.com", "example.com", false},
		{"wildcard no match subdomain", "*.example.com", "bar.foo.example.com", false},
		{"empty pattern", "", "example.com", false},
		{"empty host", "example.com", "", false},
		{"both empty", "", "", false},
		{"trailing dot on host", "example.com", "example.com.", true},
		{"trailing dot on pattern", "example.com.", "example.com", true},
		{"both trailing dots", "example.com.", "example.com.", true},
		{"wildcard not in first position", "foo.*.com", "foo.bar.com", false},
		{"single label pattern and host", "localhost", "localhost", true},
		{"single label mismatch", "localhost", "otherhost", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := matchHostnames(tc.pattern, tc.host)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestVerifyHostname(t *testing.T) {
	t.Run("exact DNS match via SAN", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"example.com", "other.com"},
		}
		err := verifyHostname("example.com", cert)
		assert.NoError(t, err)
	})

	t.Run("exact DNS match via CN when no SAN", func(t *testing.T) {
		cert := &x509.Certificate{
			Subject: pkix.Name{
				CommonName: "example.com",
			},
		}
		err := verifyHostname("example.com", cert)
		assert.NoError(t, err)
	})

	t.Run("wildcard DNS match via SAN", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"*.example.com"},
		}
		err := verifyHostname("foo.example.com", cert)
		assert.NoError(t, err)
	})

	t.Run("wildcard DNS no match via SAN", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"*.example.com"},
		}
		err := verifyHostname("example.com", cert)
		assert.Error(t, err)
	})

	t.Run("no match returns error", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"example.com"},
		}
		err := verifyHostname("other.com", cert)
		assert.Error(t, err)
	})

	t.Run("IP address match", func(t *testing.T) {
		cert := &x509.Certificate{
			IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		}
		err := verifyHostname("127.0.0.1", cert)
		assert.NoError(t, err)
	})

	t.Run("IP address in brackets match", func(t *testing.T) {
		cert := &x509.Certificate{
			IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		}
		err := verifyHostname("[127.0.0.1]", cert)
		assert.NoError(t, err)
	})

	t.Run("IP address no match", func(t *testing.T) {
		cert := &x509.Certificate{
			IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		}
		err := verifyHostname("192.168.0.1", cert)
		assert.Error(t, err)
	})

	t.Run("host with uppercase matches", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"example.com"},
		}
		err := verifyHostname("EXAMPLE.COM", cert)
		assert.NoError(t, err)
	})

	t.Run("SAN present ignores CN", func(t *testing.T) {
		cert := &x509.Certificate{
			DNSNames: []string{"example.com"},
			Subject: pkix.Name{
				CommonName: "other.com",
			},
		}
		err := verifyHostname("other.com", cert)
		assert.Error(t, err)
	})

	t.Run("IPv6 address match", func(t *testing.T) {
		cert := &x509.Certificate{
			IPAddresses: []net.IP{net.ParseIP("::1")},
		}
		err := verifyHostname("[::1]", cert)
		assert.NoError(t, err)
	})

	t.Run("empty DNS names and empty CN returns error", func(t *testing.T) {
		cert := &x509.Certificate{}
		err := verifyHostname("example.com", cert)
		assert.Error(t, err)
	})
}
