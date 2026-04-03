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

package cors

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

const enableAnnotation = "true"

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

func TestIngressCorsConfigValid(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}

	// Valid
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = enableAnnotation
	data[parser.GetAnnotationWithPrefix(corsAllowHeadersAnnotation)] = "DNT,X-CustomHeader, Keep-Alive,User-Agent"
	data[parser.GetAnnotationWithPrefix(corsAllowCredentialsAnnotation)] = "false"
	data[parser.GetAnnotationWithPrefix(corsAllowMethodsAnnotation)] = "GET, PATCH"
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "null, https://origin123.test.com:4443"
	data[parser.GetAnnotationWithPrefix(corsExposeHeadersAnnotation)] = "*, X-CustomResponseHeader"
	data[parser.GetAnnotationWithPrefix(corsMaxAgeAnnotation)] = "600"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("error parsing annotations: %v", err)
	}

	nginxCors, ok := corst.(*Config)
	if !ok {
		t.Errorf("expected a Config type but returned %t", corst)
	}

	if !nginxCors.CorsEnabled {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)], nginxCors.CorsEnabled)
	}

	if nginxCors.CorsAllowCredentials {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsAllowCredentialsAnnotation)], nginxCors.CorsAllowCredentials)
	}

	if nginxCors.CorsAllowHeaders != "DNT,X-CustomHeader, Keep-Alive,User-Agent" {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsAllowHeadersAnnotation)], nginxCors.CorsAllowHeaders)
	}

	if nginxCors.CorsAllowMethods != "GET, PATCH" {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsAllowMethodsAnnotation)], nginxCors.CorsAllowMethods)
	}

	if !reflect.DeepEqual(nginxCors.CorsAllowOrigin, []string{"null", "https://origin123.test.com:4443"}) {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)], nginxCors.CorsAllowOrigin)
	}

	if nginxCors.CorsExposeHeaders != "*, X-CustomResponseHeader" {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsExposeHeadersAnnotation)], nginxCors.CorsExposeHeaders)
	}

	if nginxCors.CorsMaxAge != 600 {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsMaxAgeAnnotation)], nginxCors.CorsMaxAge)
	}
}

func TestIngressCorsConfigInvalid(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}

	// Valid
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = "yes"
	data[parser.GetAnnotationWithPrefix(corsAllowHeadersAnnotation)] = "@alright, #ingress"
	data[parser.GetAnnotationWithPrefix(corsAllowCredentialsAnnotation)] = "no"
	data[parser.GetAnnotationWithPrefix(corsAllowMethodsAnnotation)] = "GET, PATCH, $nginx"
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "origin123.test.com:4443"
	data[parser.GetAnnotationWithPrefix(corsExposeHeadersAnnotation)] = "@alright, #ingress"
	data[parser.GetAnnotationWithPrefix(corsMaxAgeAnnotation)] = "abcd"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("error parsing annotations: %v", err)
	}

	nginxCors, ok := corst.(*Config)
	if !ok {
		t.Errorf("expected a Config type but returned %t", corst)
	}

	if nginxCors.CorsEnabled {
		t.Errorf("expected %v but returned %v", false, nginxCors.CorsEnabled)
	}

	if !nginxCors.CorsAllowCredentials {
		t.Errorf("expected %v but returned %v", true, nginxCors.CorsAllowCredentials)
	}

	if nginxCors.CorsAllowHeaders != defaultCorsHeaders {
		t.Errorf("expected %v but returned %v", defaultCorsHeaders, nginxCors.CorsAllowHeaders)
	}

	if nginxCors.CorsAllowMethods != defaultCorsMethods {
		t.Errorf("expected %v but returned %v", defaultCorsHeaders, nginxCors.CorsAllowMethods)
	}

	if nginxCors.CorsExposeHeaders != "" {
		t.Errorf("expected %v but returned %v", "", nginxCors.CorsExposeHeaders)
	}

	if nginxCors.CorsMaxAge != defaultCorsMaxAge {
		t.Errorf("expected %v but returned %v", defaultCorsMaxAge, nginxCors.CorsMaxAge)
	}
}

func TestIngressCorsConfigAllowOriginWithTrailingComma(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = enableAnnotation

	// Include a trailing comma and an empty value between the commas.
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "https://origin123.test.com:4443,    ,https://origin321.test.com:4443,"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("error parsing annotations: %v", err)
	}

	nginxCors, ok := corst.(*Config)
	if !ok {
		t.Errorf("expected a Config type but returned %t", corst)
	}

	if !nginxCors.CorsEnabled {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)], nginxCors.CorsEnabled)
	}

	expectedCorsAllowOrigins := []string{"https://origin123.test.com:4443", "https://origin321.test.com:4443"}
	if !reflect.DeepEqual(nginxCors.CorsAllowOrigin, expectedCorsAllowOrigins) {
		t.Errorf("expected %v but returned %v", expectedCorsAllowOrigins, nginxCors.CorsAllowOrigin)
	}
}

func TestIngressCorsConfigAllowOriginNull(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = enableAnnotation

	// Include a trailing comma and an empty value between the commas.
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "https://origin123.test.com:4443,null,https://origin321.test.com:4443"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("error parsing annotations: %v", err)
	}

	nginxCors, ok := corst.(*Config)
	if !ok {
		t.Errorf("expected a Config type but returned %t", corst)
	}

	if !nginxCors.CorsEnabled {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)], nginxCors.CorsEnabled)
	}

	expectedCorsAllowOrigins := []string{"https://origin123.test.com:4443", "null", "https://origin321.test.com:4443"}
	if !reflect.DeepEqual(nginxCors.CorsAllowOrigin, expectedCorsAllowOrigins) {
		t.Errorf("expected %v but returned %v", expectedCorsAllowOrigins, nginxCors.CorsAllowOrigin)
	}
}

func TestIngressCorsConfigAllowOriginWithNonHttpProtocol(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = enableAnnotation

	// Include a trailing comma and an empty value between the commas.
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "test://localhost"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("error parsing annotations: %v", err)
	}

	nginxCors, ok := corst.(*Config)
	if !ok {
		t.Errorf("expected a Config type but returned %t", corst)
	}

	if !nginxCors.CorsEnabled {
		t.Errorf("expected %v but returned %v", data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)], nginxCors.CorsEnabled)
	}

	expectedCorsAllowOrigins := []string{"test://localhost"}
	if !reflect.DeepEqual(nginxCors.CorsAllowOrigin, expectedCorsAllowOrigins) {
		t.Errorf("expected %v but returned %v", expectedCorsAllowOrigins, nginxCors.CorsAllowOrigin)
	}
}

func TestCorsOriginRegexValidator(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"wildcard subdomain", "https://*.foo.bar", true},
		{"wildcard subdomain with port", "http://*.bar.foo:8080", true},
		{"custom protocol wildcard", "myprotocol://*.abc.bar.foo:9000", true},
		{"multiple origins mixed protocols", "https://origin1.com,http://origin2.com:8080,null", true},
		{"multiple origins with wildcards", "https://*.foo.bar,http://*.baz.qux:443", true},
		{"single origin with port", "https://origin.test.com:4443", true},
		{"just star", "*", true},
		{"just null", "null", true},
		{"empty string", "", true},
		{"origin without protocol", "origin123.test.com:4443", true},
		{"origin with only port is valid", ":8080", true},
		{"invalid special chars", "https://$upstream", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, corsOriginRegexValidator.MatchString(tt.input))
		})
	}
}

func TestCorsMethodsRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"valid methods", "GET, PUT, POST", true},
		{"single method", "GET", true},
		{"lowercase methods", "get, post", true},
		{"invalid with special chars", "GET, PATCH, $nginx", false},
		{"invalid with at sign", "GET@POST", false},
		{"invalid with dollar", "$GET", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, corsMethodsRegex.MatchString(tt.input))
		})
	}
}

func TestCorsExposeHeadersRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"valid headers", "X-CustomHeader, X-Another", true},
		{"star wildcard", "*", true},
		{"mixed with star", "*, X-CustomResponseHeader", true},
		{"headers with underscore", "X_Custom_Header", true},
		{"headers with digits", "X-Header123", true},
		{"invalid with at sign", "@alright", false},
		{"invalid with hash", "#ingress", false},
		{"space separated valid headers", "X Header", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, corsExposeHeadersRegex.MatchString(tt.input))
		})
	}
}

func TestCorsWildcardOriginStopsProcessing(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(corsEnableAnnotation)] = enableAnnotation
	data[parser.GetAnnotationWithPrefix(corsAllowOriginAnnotation)] = "*, https://origin.test.com"
	ing.SetAnnotations(data)

	corst, err := NewParser(&resolver.Mock{}).Parse(ing)
	assert.NoError(t, err)

	nginxCors, ok := corst.(*Config)
	assert.True(t, ok)
	assert.Equal(t, []string{"*"}, nginxCors.CorsAllowOrigin)
}

func TestCorsConfigEqual(t *testing.T) {
	tests := []struct {
		name   string
		c1     *Config
		c2     *Config
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
			"equal configs",
			&Config{
				CorsEnabled:          true,
				CorsAllowOrigin:      []string{"https://foo.com"},
				CorsAllowMethods:     "GET, POST",
				CorsAllowHeaders:     "X-Header",
				CorsAllowCredentials: false,
				CorsExposeHeaders:    "X-Exposed",
				CorsMaxAge:           600,
			},
			&Config{
				CorsEnabled:          true,
				CorsAllowOrigin:      []string{"https://foo.com"},
				CorsAllowMethods:     "GET, POST",
				CorsAllowHeaders:     "X-Header",
				CorsAllowCredentials: false,
				CorsExposeHeaders:    "X-Exposed",
				CorsMaxAge:           600,
			},
			true,
		},
		{
			"different max age",
			&Config{CorsMaxAge: 100},
			&Config{CorsMaxAge: 200},
			false,
		},
		{
			"different expose headers",
			&Config{CorsExposeHeaders: "X-A"},
			&Config{CorsExposeHeaders: "X-B"},
			false,
		},
		{
			"different credentials",
			&Config{CorsAllowCredentials: true},
			&Config{CorsAllowCredentials: false},
			false,
		},
		{
			"different allow headers",
			&Config{CorsAllowHeaders: "X-A"},
			&Config{CorsAllowHeaders: "X-B"},
			false,
		},
		{
			"different methods",
			&Config{CorsAllowMethods: "GET"},
			&Config{CorsAllowMethods: "POST"},
			false,
		},
		{
			"different enabled",
			&Config{CorsEnabled: true},
			&Config{CorsEnabled: false},
			false,
		},
		{
			"different origin lengths",
			&Config{CorsAllowOrigin: []string{"a", "b"}},
			&Config{CorsAllowOrigin: []string{"a"}},
			false,
		},
		{
			"different origin values",
			&Config{CorsAllowOrigin: []string{"a"}},
			&Config{CorsAllowOrigin: []string{"b"}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.c1.Equal(tt.c2))
		})
	}
}
