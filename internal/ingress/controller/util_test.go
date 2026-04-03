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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	networking "k8s.io/api/networking/v1"
)

func TestRlimitMaxNumFiles(t *testing.T) {
	i := rlimitMaxNumFiles()
	if i < 1 {
		t.Errorf("returned %v but expected > 0", i)
	}
}

func TestSysctlSomaxconn(t *testing.T) {
	i := sysctlSomaxconn()
	if i < 511 {
		t.Errorf("returned %v but expected >= 511", i)
	}
}

func TestUpstreamNameSimple(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		service   *networking.IngressServiceBackend
		expect    string
	}{
		{
			name:      "service with port number",
			namespace: "default",
			service: &networking.IngressServiceBackend{
				Name: "my-service",
				Port: networking.ServiceBackendPort{Number: 80},
			},
			expect: "default-my-service-80",
		},
		{
			name:      "service with port name",
			namespace: "kube-system",
			service: &networking.IngressServiceBackend{
				Name: "my-service",
				Port: networking.ServiceBackendPort{Name: "http"},
			},
			expect: "kube-system-my-service-http",
		},
		{
			name:      "nil service",
			namespace: "default",
			service:   nil,
			expect:    "default-INVALID",
		},
		{
			name:      "service with no port",
			namespace: "default",
			service: &networking.IngressServiceBackend{
				Name: "my-service",
			},
			expect: "default-INVALID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := upstreamName(tt.namespace, tt.service)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestUpstreamServiceNameAndPortSimple(t *testing.T) {
	tests := []struct {
		name          string
		service       *networking.IngressServiceBackend
		expectName    string
		expectPort    int
		expectPortStr string
	}{
		{
			name: "service with port number",
			service: &networking.IngressServiceBackend{
				Name: "svc",
				Port: networking.ServiceBackendPort{Number: 443},
			},
			expectName: "svc",
			expectPort: 443,
		},
		{
			name: "service with port name",
			service: &networking.IngressServiceBackend{
				Name: "svc",
				Port: networking.ServiceBackendPort{Name: "https"},
			},
			expectName:    "svc",
			expectPortStr: "https",
		},
		{
			name:    "nil service",
			service: nil,
		},
		{
			name: "service with no port",
			service: &networking.IngressServiceBackend{
				Name: "svc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, port := upstreamServiceNameAndPort(tt.service)
			assert.Equal(t, tt.expectName, name)
			if tt.expectPort != 0 {
				assert.Equal(t, tt.expectPort, port.IntValue())
			}
			if tt.expectPortStr != "" {
				assert.Equal(t, tt.expectPortStr, port.StrVal)
			}
		})
	}
}

func TestNewNginxCommand(t *testing.T) {
	t.Run("uses default binary", func(t *testing.T) {
		nc := NewNginxCommand()
		assert.Equal(t, "/usr/bin/nginx", nc.Binary)
	})

	t.Run("uses env var when set", func(t *testing.T) {
		os.Setenv("NGINX_BINARY", "/custom/nginx")
		defer os.Unsetenv("NGINX_BINARY")
		nc := NewNginxCommand()
		assert.Equal(t, "/custom/nginx", nc.Binary)
	})
}

func TestNginxCommandExecCommand(t *testing.T) {
	nc := NginxCommand{Binary: "/usr/bin/nginx"}
	cmd := nc.ExecCommand("-v")
	assert.Contains(t, cmd.Args, "/usr/bin/nginx")
	assert.Contains(t, cmd.Args, "-c")
	assert.Contains(t, cmd.Args, "/etc/nginx/nginx.conf")
	assert.Contains(t, cmd.Args, "-v")
}

func TestGetSysctl(t *testing.T) {
	val, err := getSysctl("net/core/somaxconn")
	if err != nil {
		t.Skipf("could not read sysctl: %v", err)
	}
	if val < 0 {
		t.Errorf("expected positive value but got %d", val)
	}

	_, err = getSysctl("nonexistent/sysctl")
	assert.Error(t, err)
}
