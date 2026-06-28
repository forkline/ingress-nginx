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

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"k8s.io/ingress-nginx/internal/nginx"
	"k8s.io/ingress-nginx/pkg/metrics"
)

func TestMetricsEndpointRegistration(t *testing.T) {
	tests := []struct {
		name          string
		enableMetrics bool
		expectMetrics bool
	}{
		{
			name:          "metrics enabled",
			enableMetrics: true,
			expectMetrics: true,
		},
		{
			name:          "metrics disabled",
			enableMetrics: false,
			expectMetrics: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := prometheus.NewRegistry()
			mux := http.NewServeMux()

			metrics.RegisterHealthz(nginx.HealthPath, mux)
			if tt.enableMetrics {
				metrics.RegisterMetrics(reg, mux)
			}

			req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)

			if tt.expectMetrics && w.Code != http.StatusOK {
				t.Errorf("expected /metrics to be registered (status %d), got %d", http.StatusOK, w.Code)
			}
			if !tt.expectMetrics && w.Code != http.StatusNotFound {
				t.Errorf("expected /metrics to NOT be registered (status %d), got %d", http.StatusNotFound, w.Code)
			}
		})
	}
}

func TestHealthzEndpointAlwaysAvailable(t *testing.T) {
	tests := []struct {
		name          string
		enableMetrics bool
	}{
		{
			name:          "healthz available when metrics enabled",
			enableMetrics: true,
		},
		{
			name:          "healthz available when metrics disabled",
			enableMetrics: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := prometheus.NewRegistry()
			mux := http.NewServeMux()

			metrics.RegisterHealthz(nginx.HealthPath, mux)
			if tt.enableMetrics {
				metrics.RegisterMetrics(reg, mux)
			}

			req := httptest.NewRequest(http.MethodGet, nginx.HealthPath, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("expected /healthz to always be available (status %d), got %d", http.StatusOK, w.Code)
			}
		})
	}
}
