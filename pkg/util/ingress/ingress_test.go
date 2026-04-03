/*
Copyright 2022 The Kubernetes Authors.

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

package ingress

import (
	"testing"

	"github.com/stretchr/testify/assert"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/pkg/apis/ingress"
)

func TestIsDynamicConfigurationEnough(t *testing.T) {
	backends := []*ingress.Backend{{
		Name: "fakenamespace-myapp-80",
		Endpoints: []ingress.Endpoint{
			{
				Address: "10.0.0.1",
				Port:    "8080",
			},
			{
				Address: "10.0.0.2",
				Port:    "8080",
			},
		},
	}}

	servers := []*ingress.Server{{
		Hostname: "myapp.fake",
		Locations: []*ingress.Location{
			{
				Path:    "/",
				Backend: "fakenamespace-myapp-80",
			},
		},
		SSLCert: &ingress.SSLCert{
			PemCertKey: "fake-certificate",
		},
	}}

	commonConfig := &ingress.Configuration{
		Backends: backends,
		Servers:  servers,
	}

	runningConfig := &ingress.Configuration{
		Backends: backends,
		Servers:  servers,
	}

	newConfig := commonConfig
	if !IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("When new config is same as the running config it should be deemed as dynamically configurable")
	}

	newConfig = &ingress.Configuration{
		Backends: []*ingress.Backend{{Name: "another-backend-8081"}},
		Servers:  []*ingress.Server{{Hostname: "myapp1.fake"}},
	}
	if IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("Expected to not be dynamically configurable when there's more than just backends change")
	}

	newConfig = &ingress.Configuration{
		Backends: []*ingress.Backend{{Name: "a-backend-8080"}},
		Servers:  servers,
	}

	if !IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("Expected to be dynamically configurable when only backends change")
	}

	newServers := []*ingress.Server{{
		Hostname: "myapp1.fake",
		Locations: []*ingress.Location{
			{
				Path:    "/",
				Backend: "fakenamespace-myapp-80",
			},
		},
		SSLCert: &ingress.SSLCert{
			PemCertKey: "fake-certificate",
		},
	}}

	newConfig = &ingress.Configuration{
		Backends: backends,
		Servers:  newServers,
	}
	if IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("Expected to not be dynamically configurable when dynamic certificates is enabled and a non-certificate field in servers is updated")
	}

	newServers[0].Hostname = "myapp.fake"
	newServers[0].SSLCert.PemCertKey = "new-fake-certificate"

	newConfig = &ingress.Configuration{
		Backends: backends,
		Servers:  newServers,
	}
	if !IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("Expected to be dynamically configurable when only SSLCert changes")
	}

	newConfig = &ingress.Configuration{
		Backends: []*ingress.Backend{{Name: "a-backend-8080"}},
		Servers:  newServers,
	}
	if !IsDynamicConfigurationEnough(newConfig, runningConfig) {
		t.Errorf("Expected to be dynamically configurable when backend and SSLCert changes")
	}

	if !runningConfig.Equal(commonConfig) {
		t.Errorf("Expected running config to not change")
	}

	if !newConfig.Equal(&ingress.Configuration{Backends: []*ingress.Backend{{Name: "a-backend-8080"}}, Servers: newServers}) {
		t.Errorf("Expected new config to not change")
	}
}

func TestGetRemovedHosts(t *testing.T) {
	tests := []struct {
		name     string
		oldCfg   *ingress.Configuration
		newCfg   *ingress.Configuration
		expected []string
	}{
		{
			name:     "no servers in either config",
			oldCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			newCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			expected: []string{},
		},
		{
			name: "same servers",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}},
			},
			expected: []string{},
		},
		{
			name: "server removed",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}, {Hostname: "old.example.com"}},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}},
			},
			expected: []string{"old.example.com"},
		},
		{
			name: "server added",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com"}, {Hostname: "new.example.com"}},
			},
			expected: []string{},
		},
		{
			name: "multiple servers removed",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "a.example.com"}, {Hostname: "b.example.com"}, {Hostname: "c.example.com"}},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "b.example.com"}},
			},
			expected: []string{"a.example.com", "c.example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRemovedHosts(tt.oldCfg, tt.newCfg)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestGetRemovedIngresses(t *testing.T) {
	pathPrefix := networking.PathTypePrefix

	tests := []struct {
		name     string
		oldCfg   *ingress.Configuration
		newCfg   *ingress.Configuration
		expected []string
	}{
		{
			name:     "no ingresses",
			oldCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			newCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			expected: []string{},
		},
		{
			name: "same ingress",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{
						Hostname: "example.com",
						Locations: []*ingress.Location{
							{
								Path:     "/",
								PathType: &pathPrefix,
								Ingress: &ingress.Ingress{
									Ingress: networking.Ingress{
										ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
									},
								},
							},
						},
					},
				},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{
						Hostname: "example.com",
						Locations: []*ingress.Location{
							{
								Path:     "/",
								PathType: &pathPrefix,
								Ingress: &ingress.Ingress{
									Ingress: networking.Ingress{
										ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
									},
								},
							},
						},
					},
				},
			},
			expected: []string{},
		},
		{
			name: "ingress removed",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{
						Hostname: "example.com",
						Locations: []*ingress.Location{
							{Path: "/", PathType: &pathPrefix, Ingress: &ingress.Ingress{Ingress: networking.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "old", Namespace: "default"}}}},
							{Path: "/kept", PathType: &pathPrefix, Ingress: &ingress.Ingress{Ingress: networking.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "kept", Namespace: "default"}}}},
						},
					},
				},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{
						Hostname: "example.com",
						Locations: []*ingress.Location{
							{Path: "/kept", PathType: &pathPrefix, Ingress: &ingress.Ingress{Ingress: networking.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "kept", Namespace: "default"}}}},
						},
					},
				},
			},
			expected: []string{"default/old"},
		},
		{
			name: "nil ingress in location",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{Hostname: "example.com", Locations: []*ingress.Location{{Ingress: nil}}},
				},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{
					{Hostname: "example.com", Locations: []*ingress.Location{{Ingress: nil}}},
				},
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRemovedIngresses(tt.oldCfg, tt.newCfg)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestGetRemovedCertificateSerialNumbers(t *testing.T) {
	tests := []struct {
		name     string
		oldCfg   *ingress.Configuration
		newCfg   *ingress.Configuration
		expected []string
	}{
		{
			name:     "no certificates",
			oldCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			newCfg:   &ingress.Configuration{Servers: []*ingress.Server{}},
			expected: []string{},
		},
		{
			name: "nil cert skipped",
			oldCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com", SSLCert: nil}},
			},
			newCfg: &ingress.Configuration{
				Servers: []*ingress.Server{{Hostname: "example.com", SSLCert: nil}},
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRemovedCertificateSerialNumbers(tt.oldCfg, tt.newCfg)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestBuildRedirects(t *testing.T) {
	tests := []struct {
		name          string
		servers       []*ingress.Server
		expectedCount int
	}{
		{
			name:          "no servers",
			servers:       []*ingress.Server{},
			expectedCount: 0,
		},
		{
			name: "server without redirect",
			servers: []*ingress.Server{
				{Hostname: "example.com", RedirectFromToWWW: false},
			},
			expectedCount: 0,
		},
		{
			name: "server with redirect from www",
			servers: []*ingress.Server{
				{Hostname: "example.com", RedirectFromToWWW: true},
			},
			expectedCount: 1,
		},
		{
			name: "server with redirect to www",
			servers: []*ingress.Server{
				{Hostname: "www.example.com", RedirectFromToWWW: true},
			},
			expectedCount: 1,
		},
		{
			name: "multiple redirects",
			servers: []*ingress.Server{
				{Hostname: "example.com", RedirectFromToWWW: true},
				{Hostname: "other.com", RedirectFromToWWW: true},
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildRedirects(tt.servers)
			assert.Equal(t, tt.expectedCount, len(result))
		})
	}
}

func TestBuildRedirectsFromTo(t *testing.T) {
	servers := []*ingress.Server{
		{Hostname: "example.com", RedirectFromToWWW: true},
	}
	result := BuildRedirects(servers)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "www.example.com", result[0].From)
	assert.Equal(t, "example.com", result[0].To)

	servers2 := []*ingress.Server{
		{Hostname: "www.example.com", RedirectFromToWWW: true},
	}
	result2 := BuildRedirects(servers2)
	assert.Equal(t, 1, len(result2))
	assert.Equal(t, "example.com", result2[0].From)
	assert.Equal(t, "www.example.com", result2[0].To)
}

func TestClearL4ServiceEndpoints(t *testing.T) {
	config := &ingress.Configuration{
		TCPEndpoints: []ingress.L4Service{
			{Port: 9000, Endpoints: []ingress.Endpoint{{Address: "10.0.0.1", Port: "80"}}},
		},
		UDPEndpoints: []ingress.L4Service{
			{Port: 5000, Endpoints: []ingress.Endpoint{{Address: "10.0.0.2", Port: "53"}}},
		},
	}

	clearL4serviceEndpoints(config)

	assert.Empty(t, config.TCPEndpoints[0].Endpoints)
	assert.Empty(t, config.UDPEndpoints[0].Endpoints)
	assert.Equal(t, 9000, config.TCPEndpoints[0].Port)
	assert.Equal(t, 5000, config.UDPEndpoints[0].Port)
}

func TestClearCertificates(t *testing.T) {
	config := &ingress.Configuration{
		Servers: []*ingress.Server{
			{Hostname: "example.com", SSLCert: &ingress.SSLCert{}},
			{Hostname: "secure.example.com", SSLCert: &ingress.SSLCert{}},
		},
	}

	clearCertificates(config)

	assert.Nil(t, config.Servers[0].SSLCert)
	assert.Nil(t, config.Servers[1].SSLCert)
	assert.Equal(t, "example.com", config.Servers[0].Hostname)
	assert.Equal(t, "secure.example.com", config.Servers[1].Hostname)
}
