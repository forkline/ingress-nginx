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

package flags

import (
	"os"
	"testing"
	"time"

	"k8s.io/ingress-nginx/internal/ingress/controller"
	"k8s.io/ingress-nginx/internal/ingress/controller/config"
)

const (
	testCmd                      = "cmd"
	testHTTPPortFlag             = "--http-port"
	testHTTPSPortFlag            = "--https-port"
	testPortZero                 = "0"
	testEnableSSLPassthrough     = "--enable-ssl-passthrough" //nolint:gosec // This is a flag name, not credentials
	testElectionTTLFlag          = "--election-ttl"
	testDefaultBackendSvcFlag    = "--default-backend-service"
	testSSLPassthroughProxyFlag  = "--ssl-passthrough-proxy-port" //nolint:gosec // This is a flag name, not credentials
	testPublishSvcFlag           = "--publish-service"
	testPublishStatusAddressFlag = "--publish-status-address"
	testMaxmindEditionIDsFlag    = "--maxmind-edition-ids"
	testMaxmindLicenseKeyFlag    = "--maxmind-license-key"
	testMaxmindMirrorFlag        = "--maxmind-mirror"
	testNamespaceTest            = "namespace/test"
)

func TestNoMandatoryFlag(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{
		testCmd,
		testHTTPPortFlag, testPortZero,
		testHTTPSPortFlag, testPortZero,
		"--default-server-port", testPortZero,
		"--status-port", testPortZero,
		"--stream-port", testPortZero,
		"--profiler-port", testPortZero,
	}

	_, _, err := ParseFlags()
	if err != nil {
		t.Fatalf("Expected no error but got: %s", err)
	}
}

func TestDefaults(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{
		testCmd,
		testDefaultBackendSvcFlag, testNamespaceTest,
		testHTTPPortFlag, testPortZero,
		testHTTPSPortFlag, testPortZero,
	}

	showVersion, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if showVersion {
		t.Fatal("Expected flag \"show-version\" to be false")
	}

	if conf == nil {
		t.Fatal("Expected a controller Configuration")
	}
}

func TestSetupSSLProxy(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectError    bool
		description    string
		validateConfig func(t *testing.T, _ bool, cfg *controller.Configuration)
	}{
		{
			name:        "valid SSL proxy configuration with passthrough enabled",
			args:        []string{testCmd, testEnableSSLPassthrough, testSSLPassthroughProxyFlag, "9999", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should accept valid SSL proxy port with passthrough enabled",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
				if cfg.ListenPorts.SSLProxy != 9999 {
					t.Errorf("Expected SSLProxy port to be 9999, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
		{
			name:        "SSL proxy port without explicit passthrough enabling",
			args:        []string{testCmd, testSSLPassthroughProxyFlag, "8443", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should accept SSL proxy port configuration without explicit passthrough enable",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if cfg.ListenPorts.SSLProxy != 8443 {
					t.Errorf("Expected SSLProxy port to be 8443, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
		{
			name:        "SSL proxy with default backend service",
			args:        []string{testCmd, testEnableSSLPassthrough, testDefaultBackendSvcFlag, "default/backend", testSSLPassthroughProxyFlag, "9000", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should work with default backend service and SSL passthrough",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
				if cfg.DefaultService != "default/backend" {
					t.Errorf("Expected DefaultService to be 'default/backend', got %s", cfg.DefaultService)
				}
				if cfg.ListenPorts.SSLProxy != 9000 {
					t.Errorf("Expected SSLProxy port to be 9000, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
		{
			name:        "SSL proxy with default SSL certificate",
			args:        []string{testCmd, testEnableSSLPassthrough, "--default-ssl-certificate", "default/tls-cert", testSSLPassthroughProxyFlag, "8444", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should work with default SSL certificate and passthrough",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
				if cfg.DefaultSSLCertificate != "default/tls-cert" {
					t.Errorf("Expected DefaultSSLCertificate to be 'default/tls-cert', got %s", cfg.DefaultSSLCertificate)
				}
				if cfg.ListenPorts.SSLProxy != 8444 {
					t.Errorf("Expected SSLProxy port to be 8444, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
		{
			name:        "SSL proxy with chain completion enabled",
			args:        []string{testCmd, testEnableSSLPassthrough, "--enable-ssl-chain-completion", testSSLPassthroughProxyFlag, "7443", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should work with SSL chain completion and passthrough",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
				if !config.EnableSSLChainCompletion {
					t.Error("Expected EnableSSLChainCompletion to be true")
				}
				if cfg.ListenPorts.SSLProxy != 7443 {
					t.Errorf("Expected SSLProxy port to be 7443, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
		{
			name:        "SSL proxy with minimal configuration",
			args:        []string{testCmd, testEnableSSLPassthrough, testSSLPassthroughProxyFlag, "0", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should work with minimal SSL passthrough configuration",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
			},
		},
		{
			name:        "SSL proxy with comprehensive configuration",
			args:        []string{testCmd, testEnableSSLPassthrough, "--enable-ssl-chain-completion", "--default-ssl-certificate", "kube-system/default-cert", testDefaultBackendSvcFlag, "kube-system/default-backend", testSSLPassthroughProxyFlag, "10443", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero},
			expectError: false,
			description: "Should work with comprehensive SSL proxy configuration",
			validateConfig: func(t *testing.T, _ bool, cfg *controller.Configuration) {
				if !cfg.EnableSSLPassthrough {
					t.Error("Expected EnableSSLPassthrough to be true")
				}
				if !config.EnableSSLChainCompletion {
					t.Error("Expected EnableSSLChainCompletion to be true")
				}
				if cfg.DefaultSSLCertificate != "kube-system/default-cert" {
					t.Errorf("Expected DefaultSSLCertificate to be 'kube-system/default-cert', got %s", cfg.DefaultSSLCertificate)
				}
				if cfg.DefaultService != "kube-system/default-backend" {
					t.Errorf("Expected DefaultService to be 'kube-system/default-backend', got %s", cfg.DefaultService)
				}
				if cfg.ListenPorts.SSLProxy != 10443 {
					t.Errorf("Expected SSLProxy port to be 10443, got %d", cfg.ListenPorts.SSLProxy)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetForTesting(func() { t.Fatal("Parsing failed") })

			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args

			showVersion, cfg, err := ParseFlags()
			if tt.expectError && err == nil {
				t.Fatalf("Expected error for %s, but got none", tt.description)
			}
			if !tt.expectError && err != nil {
				t.Fatalf("Expected no error for %s, got: %v", tt.description, err)
			}

			if !tt.expectError && tt.validateConfig != nil {
				tt.validateConfig(t, showVersion, cfg)
			}
		})
	}
}

func TestFlagConflict(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testPublishSvcFlag, testNamespaceTest, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testPublishStatusAddressFlag, "1.1.1.1"}

	_, _, err := ParseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMaxmindEdition(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testPublishSvcFlag, testNamespaceTest, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testMaxmindLicenseKeyFlag, "0000000", testMaxmindEditionIDsFlag, "GeoLite2-City, TestCheck"}

	_, _, err := ParseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMaxmindMirror(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testPublishSvcFlag, testNamespaceTest, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testMaxmindMirrorFlag, "http://geoip.local", testMaxmindLicenseKeyFlag, "0000000", testMaxmindEditionIDsFlag, "GeoLite2-City, TestCheck"}

	_, _, err := ParseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMaxmindRetryDownload(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testPublishSvcFlag, testNamespaceTest, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testMaxmindMirrorFlag, "http://127.0.0.1", testMaxmindLicenseKeyFlag, "0000000", testMaxmindEditionIDsFlag, "GeoLite2-City", "--maxmind-retries-timeout", "1s", "--maxmind-retries-count", "3"}

	_, _, err := ParseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestDisableLeaderElectionFlag(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, "--disable-leader-election", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if !conf.DisableLeaderElection {
		t.Fatalf("Expected --disable-leader-election and conf.DisableLeaderElection as true, but found: %v", conf.DisableLeaderElection)
	}
}

func TestIfLeaderElectionDisabledFlagIsFalse(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if conf.DisableLeaderElection {
		t.Fatalf("Expected --disable-leader-election and conf.DisableLeaderElection as false, but found: %v", conf.DisableLeaderElection)
	}
}

func TestLeaderElectionTTLDefaultValue(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if conf.ElectionTTL != 30*time.Second {
		t.Fatalf("Expected --election-ttl and conf.ElectionTTL as 30s, but found: %v", conf.ElectionTTL)
	}
}

func TestLeaderElectionTTLParseValueInSeconds(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testElectionTTLFlag, "10s"}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if conf.ElectionTTL != 10*time.Second {
		t.Fatalf("Expected --election-ttl and conf.ElectionTTL as 10s, but found: %v", conf.ElectionTTL)
	}
}

func TestLeaderElectionTTLParseValueInMinutes(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testElectionTTLFlag, "10m"}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if conf.ElectionTTL != 10*time.Minute {
		t.Fatalf("Expected --election-ttl and conf.ElectionTTL as 10m, but found: %v", conf.ElectionTTL)
	}
}

func TestLeaderElectionTTLParseValueInHours(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero, testElectionTTLFlag, "1h"}

	_, conf, err := ParseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if conf.ElectionTTL != 1*time.Hour {
		t.Fatalf("Expected --election-ttl and conf.ElectionTTL as 1h, but found: %v", conf.ElectionTTL)
	}
}

func TestMetricsPerUndefinedHost(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, "--metrics-per-undefined-host=true", testHTTPPortFlag, testPortZero, testHTTPSPortFlag, testPortZero}

	_, _, err := ParseFlags()
	if err != nil {
		t.Fatalf("Expected no error but got: %s", err)
	}
}

func TestMetricsPerUndefinedHostWithMetricsPerHostFalse(t *testing.T) {
	ResetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{testCmd, "--metrics-per-host=false", "--metrics-per-undefined-host=true"}

	_, _, err := ParseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}
