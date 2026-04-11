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

package template

import (
	"os"
	"path"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/ingress-nginx/internal/ingress/annotations/cors"
	"k8s.io/ingress-nginx/internal/ingress/annotations/proxy"
	"k8s.io/ingress-nginx/internal/ingress/annotations/rewrite"
	"k8s.io/ingress-nginx/internal/ingress/controller/config"
	"k8s.io/ingress-nginx/internal/nginx"
	"k8s.io/ingress-nginx/pkg/apis/ingress"
)

func goldenDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(pwd, "../../../../test/data/golden"), nil
}

func mustRenderTemplate(t *testing.T, tc *config.TemplateConfig) string {
	t.Helper()

	if tc.ListenPorts == nil {
		tc.ListenPorts = &config.ListenPorts{}
	}
	if tc.Cfg.DefaultSSLCertificate == nil {
		tc.Cfg.DefaultSSLCertificate = &ingress.SSLCert{}
	}
	tc.Cfg.WorkerProcesses = "8"

	ngxTpl, err := NewTemplate(nginx.TemplatePath)
	assert.NoError(t, err, "loading NGINX template")

	out, err := ngxTpl.Write(tc)
	assert.NoError(t, err, "rendering NGINX template")

	return string(out)
}

func TestGoldenTemplates(t *testing.T) {
	pathPrefix := networking.PathTypePrefix
	pathExact := networking.PathTypeExact

	scenarios := []struct {
		name string
		data func() *config.TemplateConfig
	}{
		{
			name: "basic_http",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "example.com",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "default-example-svc-80",
								},
							},
						},
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{
							Name: "default-example-svc-80",
						},
						{
							Name: "upstream-default-backend",
						},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
				}
			},
		},
		{
			name: "https_tls",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "secure.example.com",
							SSLCert: &ingress.SSLCert{
								PemFileName: "/etc/ingress-controller/ssl/default-tls.pem",
								PemSHA:      "abc123",
							},
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "default-secure-svc-443",
									Rewrite: rewrite.Config{
										SSLRedirect: true,
									},
								},
							},
						},
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{
							Name: "default-secure-svc-443",
						},
						{
							Name: "upstream-default-backend",
						},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
				}
			},
		},
		{
			name: "multi_host",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "app1.example.com",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "default-app1-svc-80",
								},
							},
						},
						{
							Hostname: "app2.example.com",
							Locations: []*ingress.Location{
								{
									Path:     "/api",
									PathType: &pathPrefix,
									Backend:  "default-app2-api-svc-8080",
								},
								{
									Path:     "/web",
									PathType: &pathPrefix,
									Backend:  "default-app2-web-svc-80",
								},
							},
						},
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{Name: "default-app1-svc-80"},
						{Name: "default-app2-api-svc-8080"},
						{Name: "default-app2-web-svc-80"},
						{Name: "upstream-default-backend"},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
				}
			},
		},
		{
			name: "canary",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "canary.example.com",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "default-main-svc-80",
								},
							},
						},
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{
							Name:                "default-main-svc-80",
							AlternativeBackends: []string{"default-canary-svc-80"},
							TrafficShapingPolicy: ingress.TrafficShapingPolicy{
								Weight:      80,
								WeightTotal: 100,
							},
						},
						{
							Name:     "default-canary-svc-80",
							NoServer: true,
							TrafficShapingPolicy: ingress.TrafficShapingPolicy{
								Weight:      20,
								WeightTotal: 100,
							},
						},
						{Name: "upstream-default-backend"},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
				}
			},
		},
		{
			name: "tcp_stream",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{Name: "upstream-default-backend"},
					},
					TCPBackends: []ingress.L4Service{
						{
							Port: 9000,
							Backend: ingress.L4Backend{
								Name:      "tcp-echo-svc",
								Namespace: "default",
								Port:      intstr.FromInt(2701),
								Protocol:  apiv1.ProtocolTCP,
							},
							Endpoints: []ingress.Endpoint{
								{
									Address: "10.0.0.1",
									Port:    "2701",
								},
							},
						},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
					StreamPort: 10255,
				}
			},
		},
		{
			name: "custom_annotations",
			data: func() *config.TemplateConfig {
				return &config.TemplateConfig{
					BacklogSize: 32768,
					Cfg:         config.NewDefault(),
					Servers: []*ingress.Server{
						{
							Hostname: "annotated.example.com",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "default-annotated-svc-80",
									Rewrite: rewrite.Config{
										Target:   "/app",
										UseRegex: true,
									},
									Proxy: proxy.Config{
										BodySize:            "2m",
										ConnectTimeout:      10,
										ReadTimeout:         120,
										SendTimeout:         120,
										BuffersNumber:       8,
										BufferSize:          "16k",
										NextUpstream:        "error timeout http_502",
										NextUpstreamTimeout: 0,
										NextUpstreamTries:   5,
										RequestBuffering:    "on",
									},
									CorsConfig: cors.Config{
										CorsEnabled:          true,
										CorsAllowOrigin:      []string{"https://origin.example.com"},
										CorsAllowMethods:     "GET, PUT, POST, DELETE, PATCH, OPTIONS",
										CorsAllowHeaders:     "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,Content-Type,Authorization",
										CorsAllowCredentials: true,
										CorsMaxAge:           86400,
									},
								},
								{
									Path:     "/exact",
									PathType: &pathExact,
									Backend:  "default-annotated-svc-80",
								},
							},
						},
						{
							Hostname: "_",
							Locations: []*ingress.Location{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend:  "upstream-default-backend",
								},
							},
						},
					},
					Backends: []*ingress.Backend{
						{Name: "default-annotated-svc-80"},
						{Name: "upstream-default-backend"},
					},
					ListenPorts: &config.ListenPorts{
						HTTP:     80,
						HTTPS:    443,
						Health:   10254,
						Default:  8181,
						SSLProxy: 442,
					},
				}
			},
		},
	}

	dir, err := goldenDir()
	assert.NoError(t, err, "resolving golden directory")

	update := os.Getenv("UPDATE_GOLDEN") == "1"

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			tc := sc.data()
			rendered := mustRenderTemplate(t, tc)

			goldenPath := path.Join(dir, sc.name+".conf")

			if update {
				err := os.MkdirAll(dir, 0o755)
				assert.NoError(t, err, "creating golden directory")
				err = os.WriteFile(goldenPath, []byte(rendered), 0o644)
				assert.NoError(t, err, "writing golden file")
				t.Logf("updated golden file: %s", goldenPath)
				return
			}

			expected, err := os.ReadFile(goldenPath)
			if os.IsNotExist(err) {
				t.Fatalf("golden file %s does not exist. Run with UPDATE_GOLDEN=1 to create it.", goldenPath)
			}
			assert.NoError(t, err, "reading golden file")

			assert.Equal(t, string(expected), rendered, "rendered output does not match golden file %s", goldenPath)
		})
	}
}

func TestGoldenTemplateFromConfigJSON(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err)

	data, err := os.ReadFile(path.Join(pwd, "../../../../test/data/config.json"))
	assert.NoError(t, err)

	var dat config.TemplateConfig
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &dat)
	assert.NoError(t, err)

	if dat.ListenPorts == nil {
		dat.ListenPorts = &config.ListenPorts{}
	}
	dat.Cfg.DefaultSSLCertificate = &ingress.SSLCert{}

	ngxTpl, err := NewTemplate(nginx.TemplatePath)
	assert.NoError(t, err)

	rendered, err := ngxTpl.Write(&dat)
	assert.NoError(t, err)

	dir, err := goldenDir()
	assert.NoError(t, err)

	goldenPath := path.Join(dir, "config_json.conf")

	if os.Getenv("UPDATE_GOLDEN") == "1" {
		err = os.MkdirAll(dir, 0o755)
		assert.NoError(t, err)
		err = os.WriteFile(goldenPath, rendered, 0o644)
		assert.NoError(t, err)
		t.Logf("updated golden file: %s", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if os.IsNotExist(err) {
		t.Fatalf("golden file %s does not exist. Run with UPDATE_GOLDEN=1 to create it.", goldenPath)
	}
	assert.NoError(t, err)

	assert.Equal(t, string(expected), string(rendered), "rendered output does not match golden file %s", goldenPath)
}
