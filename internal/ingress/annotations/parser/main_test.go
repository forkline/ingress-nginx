/*
Copyright 2016 The Kubernetes Authors.

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

package parser

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func buildIngress() *networking.Ingress {
	return &networking.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{},
	}
}

func TestGetBoolAnnotation(t *testing.T) {
	ing := buildIngress()

	_, err := GetBoolAnnotation("", nil, nil)
	if err == nil {
		t.Errorf("expected error but retuned nil")
	}

	tests := []struct {
		name   string
		field  string
		value  string
		exp    bool
		expErr bool
	}{
		{"valid - false", "bool", "false", false, false},
		{"valid - true", "bool", "true", true, false},
	}

	data := map[string]string{}
	ing.SetAnnotations(data)

	for _, test := range tests {
		data[GetAnnotationWithPrefix(test.field)] = test.value
		ing.SetAnnotations(data)
		u, err := GetBoolAnnotation(test.field, ing, nil)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but retuned nil", test.name)
			}
			continue
		}
		if u != test.exp {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned, %+v", test.name, test.exp, u, ing)
		}

		delete(data, test.field)
	}
}

func TestGetStringAnnotation(t *testing.T) {
	ing := buildIngress()

	_, err := GetStringAnnotation("", nil, nil)
	if err == nil {
		t.Errorf("expected error but none returned")
	}

	tests := []struct {
		name   string
		field  string
		value  string
		exp    string
		expErr bool
	}{
		{"valid - A", "string", "A ", "A", false},
		{"valid - B", "string", "	B", "B", false},
		{"empty", "string", " ", "", true},
		{
			"valid multiline", "string", `
		rewrite (?i)/arcgis/rest/services/Utilities/Geometry/GeometryServer(.*)$ /arcgis/rest/services/Utilities/Geometry/GeometryServer$1 break;
		rewrite (?i)/arcgis/services/Utilities/Geometry/GeometryServer(.*)$ /arcgis/services/Utilities/Geometry/GeometryServer$1 break;
		`, `
rewrite (?i)/arcgis/rest/services/Utilities/Geometry/GeometryServer(.*)$ /arcgis/rest/services/Utilities/Geometry/GeometryServer$1 break;
rewrite (?i)/arcgis/services/Utilities/Geometry/GeometryServer(.*)$ /arcgis/services/Utilities/Geometry/GeometryServer$1 break;
`,
			false,
		},
	}

	data := map[string]string{}
	ing.SetAnnotations(data)

	for _, test := range tests {
		data[GetAnnotationWithPrefix(test.field)] = test.value

		s, err := GetStringAnnotation(test.field, ing, nil)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but none returned", test.name)
			}
			continue
		}
		if !test.expErr {
			if err != nil {
				t.Errorf("%v: didn't expected error but error was returned: %v", test.name, err)
			}
			continue
		}
		if s != test.exp {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.name, test.exp, s)
		}

		delete(data, test.field)
	}
}

func TestGetFloatAnnotation(t *testing.T) {
	ing := buildIngress()

	_, err := GetFloatAnnotation("", nil, nil)
	if err == nil {
		t.Errorf("expected error but retuned nil")
	}

	tests := []struct {
		name   string
		field  string
		value  string
		exp    float32
		expErr bool
	}{
		{"valid - A", "string", "1.5", 1.5, false},
		{"valid - B", "string", "2", 2, false},
		{"valid - C", "string", "100.0", 100, false},
	}

	data := map[string]string{}
	ing.SetAnnotations(data)

	for _, test := range tests {
		data[GetAnnotationWithPrefix(test.field)] = test.value

		s, err := GetFloatAnnotation(test.field, ing, nil)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but retuned nil", test.name)
			}
			continue
		}
		if s != test.exp {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.name, test.exp, s)
		}

		delete(data, test.field)
	}
}

func TestGetIntAnnotation(t *testing.T) {
	ing := buildIngress()

	_, err := GetIntAnnotation("", nil, nil)
	if err == nil {
		t.Errorf("expected error but retuned nil")
	}

	tests := []struct {
		name   string
		field  string
		value  string
		exp    int
		expErr bool
	}{
		{"valid - A", "string", "1", 1, false},
		{"valid - B", "string", "2", 2, false},
	}

	data := map[string]string{}
	ing.SetAnnotations(data)

	for _, test := range tests {
		data[GetAnnotationWithPrefix(test.field)] = test.value

		s, err := GetIntAnnotation(test.field, ing, nil)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but retuned nil", test.name)
			}
			continue
		}
		if s != test.exp {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.name, test.exp, s)
		}

		delete(data, test.field)
	}
}

func TestStringToURL(t *testing.T) {
	validURL := "http://bar.foo.com/external-auth"
	validParsedURL, err := url.Parse(validURL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	tests := []struct {
		title   string
		url     string
		message string
		parsed  *url.URL
		expErr  bool
	}{
		{"empty", "", "url scheme is empty", nil, true},
		{"no scheme", "bar", "url scheme is empty", nil, true},
		{"invalid parse", "://lala.com", "://lala.com is not a valid URL: parse \"://lala.com\": missing protocol scheme", nil, true},
		{"invalid host", "http://", "url host is empty", nil, true},
		{"invalid host (multiple dots)", "http://foo..bar.com", "invalid url host", nil, true},
		{"valid URL", validURL, "", validParsedURL, false},
	}

	for _, test := range tests {
		i, err := StringToURL(test.url)
		if test.expErr {
			if err == nil {
				t.Fatalf("%v: expected error but none returned", test.title)
			}

			if err.Error() != test.message {
				t.Errorf("%v: expected error \"%v\" but \"%v\" was returned", test.title, test.message, err)
			}
			continue
		}

		if i.String() != test.parsed.String() {
			t.Errorf("%v: expected \"%v\" but \"%v\" was returned", test.title, test.parsed, i)
		}
	}
}

func TestAnnotationsReferencesConfigmap(t *testing.T) {
	tests := []struct {
		name   string
		ing    *networking.Ingress
		expect bool
	}{
		{
			name:   "nil ingress",
			ing:    nil,
			expect: false,
		},
		{
			name: "no annotations",
			ing: &networking.Ingress{
				ObjectMeta: v1.ObjectMeta{},
			},
			expect: false,
		},
		{
			name: "non configmap annotation",
			ing: &networking.Ingress{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						"nginx.ingress.kubernetes.io/rewrite-target": "/new",
					},
				},
			},
			expect: false,
		},
		{
			name: "auth-proxy-set-header annotation",
			ing: &networking.Ingress{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						"auth-proxy-set-header": "configmap",
					},
				},
			},
			expect: true,
		},
		{
			name: "fastcgi-params-configmap annotation",
			ing: &networking.Ingress{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						"fastcgi-params-configmap": "configmap",
					},
				},
			},
			expect: true,
		},
		{
			name: "mixed annotations with configmap",
			ing: &networking.Ingress{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						"nginx.ingress.kubernetes.io/rewrite-target": "/new",
						"auth-proxy-set-header":                      "configmap",
					},
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnnotationsReferencesConfigmap(tt.ing)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestNormalizeString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "no trimming needed",
			input:  "hello world",
			expect: "hello world",
		},
		{
			name:   "trims spaces on each line",
			input:  "  hello  \n  world  ",
			expect: "hello\nworld",
		},
		{
			name:   "single line with spaces",
			input:  "  trimmed  ",
			expect: "trimmed",
		},
		{
			name:   "empty string",
			input:  "",
			expect: "",
		},
		{
			name:   "multiple empty lines",
			input:  "\n\n\n",
			expect: "\n\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeString(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestTrimAnnotationPrefix(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "full prefix",
			input:  "nginx.ingress.kubernetes.io/rewrite-target",
			expect: "rewrite-target",
		},
		{
			name:   "no prefix",
			input:  "some-annotation",
			expect: "some-annotation",
		},
		{
			name:   "partial prefix",
			input:  "nginx.ingress.kubernetes.io",
			expect: "nginx.ingress.kubernetes.io",
		},
		{
			name:   "empty string",
			input:  "",
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimAnnotationPrefix(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestGetAnnotationWithPrefix(t *testing.T) {
	result := GetAnnotationWithPrefix("rewrite-target")
	assert.Equal(t, "nginx.ingress.kubernetes.io/rewrite-target", result)
}
