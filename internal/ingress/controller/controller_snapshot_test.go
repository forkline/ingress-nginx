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

package controller

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/internal/ingress/annotations"
	"k8s.io/ingress-nginx/internal/ingress/annotations/authreq"
	"k8s.io/ingress-nginx/internal/ingress/annotations/modsecurity"
	ngx_config "k8s.io/ingress-nginx/internal/ingress/controller/config"
	"k8s.io/ingress-nginx/internal/ingress/defaults"
	"k8s.io/ingress-nginx/internal/ingress/metric"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/pkg/apis/ingress"
)

func goldenControllerDir() string {
	_, filename, unused1, unused2 := runtime.Caller(0)
	_ = unused1
	_ = unused2
	repoRoot := filepath.Join(filepath.Dir(filename), "..", "..", "..")
	return filepath.Join(repoRoot, "test", "data", "golden", "controller")
}

func snapshotJSON(t *testing.T, name string, data interface{}) {
	t.Helper()
	dir := goldenControllerDir()
	path := filepath.Join(dir, name+".json")

	if os.Getenv("UPDATE_GOLDEN") == "1" {
		err := os.MkdirAll(dir, 0o755)
		require.NoError(t, err)
		raw, err := json.MarshalIndent(data, "", "  ")
		require.NoError(t, err)
		err = os.WriteFile(path, append(raw, '\n'), 0o644)
		require.NoError(t, err)
		return
	}

	raw, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		t.Fatalf("golden file %s does not exist. Run with UPDATE_GOLDEN=1 to create it.", path)
	}
	require.NoError(t, err)

	actual, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)
	assert.Equal(t, string(raw), string(actual)+"\n", "snapshot mismatch for %s", name)
}

type snapshotStore struct {
	services       map[string]*apiv1.Service
	backendCfg     ngx_config.Configuration
	defaultBackend defaults.Backend
	ingresses      []*ingress.Ingress
}

func (s *snapshotStore) GetBackendConfiguration() ngx_config.Configuration {
	return s.backendCfg
}

func (s *snapshotStore) GetSecurityConfiguration() defaults.SecurityConfiguration {
	return defaults.SecurityConfiguration{
		AllowCrossNamespaceResources: s.backendCfg.AllowCrossNamespaceResources,
		AnnotationsRiskLevel:         s.backendCfg.AnnotationsRiskLevel,
	}
}

func (s *snapshotStore) GetConfigMap(_ string) (*apiv1.ConfigMap, error) {
	return nil, nil
}

func (s *snapshotStore) GetSecret(_ string) (*apiv1.Secret, error) {
	return nil, nil
}

func (s *snapshotStore) GetService(key string) (*apiv1.Service, error) {
	if svc, ok := s.services[key]; ok {
		return svc, nil
	}
	return nil, nil
}

func (s *snapshotStore) GetServiceEndpointsSlices(_ string) ([]*networking.Ingress, error) {
	return nil, nil
}

func (s *snapshotStore) ListIngresses() []*ingress.Ingress {
	return s.ingresses
}

func (s *snapshotStore) FilterIngresses(ingresses []*ingress.Ingress, _ func(*ingress.Ingress) bool) []*ingress.Ingress {
	return ingresses
}

func (s *snapshotStore) GetLocalSSLCert(_ string) (*ingress.SSLCert, error) {
	return nil, nil
}

func (s *snapshotStore) ListLocalSSLCerts() []*ingress.SSLCert {
	return nil
}

func (s *snapshotStore) GetAuthCertificate(_ string) (*resolver.AuthSSLCert, error) {
	return nil, nil
}

func (s *snapshotStore) GetDefaultBackend() defaults.Backend {
	return s.defaultBackend
}

func (s *snapshotStore) Run(_ chan struct{}) {}

func (s *snapshotStore) GetIngressClass(_ *networking.Ingress, _ *interface{}) (string, error) {
	return "nginx", nil
}

func TestDropSnippetDirectives(t *testing.T) {
	tests := []struct {
		name string
		anns *annotations.Ingress
	}{
		{
			name: "nil_annotations",
			anns: nil,
		},
		{
			name: "all_snippets",
			anns: &annotations.Ingress{
				ConfigurationSnippet: "config snippet",
				ServerSnippet:        "server snippet",
				ModSecurity: modsecurity.Config{
					Snippet: "modsec snippet",
				},
				ExternalAuth: authreq.Config{
					AuthSnippet: "auth snippet",
				},
				StreamSnippet: "stream snippet",
			},
		},
		{
			name: "some_snippets",
			anns: &annotations.Ingress{
				ConfigurationSnippet: "config snippet",
				ServerSnippet:        "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dropSnippetDirectives(tt.anns, "test-ns/test-ingress")
			if tt.anns != nil {
				snapshotJSON(t, "drop_snippet_"+tt.name, tt.anns)
			}
		})
	}
}

func TestGetIngressPodZone(t *testing.T) {
	svc := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-svc",
			Namespace: "default",
		},
	}
	result := getIngressPodZone(svc)
	assert.Equal(t, "", result)
}

func TestGetDefaultUpstream(t *testing.T) {
	tests := []struct {
		name           string
		defaultService string
		expectedName   string
	}{
		{
			name:           "empty_default_service",
			defaultService: "",
			expectedName:   defUpstreamName,
		},
		{
			name:           "service_not_found",
			defaultService: "default/missing-svc",
			expectedName:   defUpstreamName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &ingress.Backend{
				Name: defUpstreamName,
				Endpoints: []ingress.Endpoint{
					{Address: "127.0.0.1", Port: "8181"},
				},
			}

			snapshotJSON(t, "default_upstream_"+tt.name, map[string]interface{}{
				"name":      du.Name,
				"endpoints": du.Endpoints,
			})
		})
	}
}

func TestGetStreamServices(t *testing.T) {
	tests := []struct {
		name        string
		cmName      string
		cmData      map[string]string
		proto       apiv1.Protocol
		listenPorts *ngx_config.ListenPorts
	}{
		{
			name:        "empty_configmap_name",
			cmName:      "",
			cmData:      nil,
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{},
		},
		{
			name:   "single_tcp_service",
			cmName: "default/tcp-services",
			cmData: map[string]string{
				"9000": "default/tcp-echo:2701",
			},
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{},
		},
		{
			name:   "tcp_service_with_proxy_protocol",
			cmName: "default/tcp-services",
			cmData: map[string]string{
				"9000": "default/tcp-echo:2701:PROXY:PROXY",
			},
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{},
		},
		{
			name:   "udp_service",
			cmName: "default/udp-services",
			cmData: map[string]string{
				"5000": "default/udp-echo:5000",
			},
			proto:       apiv1.ProtocolUDP,
			listenPorts: &ngx_config.ListenPorts{},
		},
		{
			name:   "reserved_port_skipped",
			cmName: "default/tcp-services",
			cmData: map[string]string{
				"80": "default/tcp-echo:80",
			},
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{HTTP: 80},
		},
		{
			name:   "invalid_port_number",
			cmName: "default/tcp-services",
			cmData: map[string]string{
				"not-a-number": "default/tcp-echo:80",
			},
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{},
		},
		{
			name:   "invalid_service_reference",
			cmName: "default/tcp-services",
			cmData: map[string]string{
				"9000": "invalid-reference",
			},
			proto:       apiv1.ProtocolTCP,
			listenPorts: &ngx_config.ListenPorts{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NGINXController{
				store: &fakeIngressStore{
					configuration: ngx_config.Configuration{},
				},
				cfg: &Configuration{
					ListenPorts: tt.listenPorts,
				},
				metricCollector: metric.DummyCollector{},
			}

			services := n.getStreamServices(tt.cmName, tt.proto)

			snapshotJSON(t, "stream_services_"+tt.name, map[string]interface{}{
				"configmap": tt.cmName,
				"protocol":  string(tt.proto),
				"services":  services,
			})
		})
	}
}

func TestCanMergeBackend(t *testing.T) {
	tests := []struct {
		name        string
		primary     *ingress.Backend
		alternative *ingress.Backend
		expected    bool
	}{
		{
			name:        "alternative is nil",
			primary:     &ingress.Backend{Name: "primary"},
			alternative: nil,
			expected:    false,
		},
		{
			name:        "same name",
			primary:     &ingress.Backend{Name: "same"},
			alternative: &ingress.Backend{Name: "same"},
			expected:    false,
		},
		{
			name:        "primary is default upstream",
			primary:     &ingress.Backend{Name: defUpstreamName},
			alternative: &ingress.Backend{Name: "alt"},
			expected:    false,
		},
		{
			name:        "primary has no server",
			primary:     &ingress.Backend{Name: "primary", NoServer: true},
			alternative: &ingress.Backend{Name: "alt"},
			expected:    false,
		},
		{
			name:        "valid merge",
			primary:     &ingress.Backend{Name: "primary"},
			alternative: &ingress.Backend{Name: "alt"},
			expected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canMergeBackend(tt.primary, tt.alternative)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUpstreamName(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		service   *networking.IngressServiceBackend
		expected  string
	}{
		{
			name:      "numeric port",
			namespace: "default",
			service: &networking.IngressServiceBackend{
				Name: "my-svc",
				Port: networking.ServiceBackendPort{Number: 80},
			},
			expected: "default-my-svc-80",
		},
		{
			name:      "named port",
			namespace: "kube-system",
			service: &networking.IngressServiceBackend{
				Name: "dns-svc",
				Port: networking.ServiceBackendPort{Name: "udp"},
			},
			expected: "kube-system-dns-svc-udp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := upstreamName(tt.namespace, tt.service)
			assert.Equal(t, tt.expected, result)
		})
	}
}
