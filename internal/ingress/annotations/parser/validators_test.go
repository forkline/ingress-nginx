/*
Copyright 2023 The Kubernetes Authors.

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
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	networking "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidateArrayOfServerName(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "should accept common name",
			value:   "something.com,anything.com",
			wantErr: false,
		},
		{
			name:    "should accept wildcard name",
			value:   "*.something.com,otherthing.com",
			wantErr: false,
		},
		{
			name:    "should allow names with spaces between array and some regexes",
			value:   `~^www\d+\.example\.com$,something.com`,
			wantErr: false,
		},
		{
			name:    "should allow names with regexes",
			value:   `http://some.test.env.com:2121/$someparam=1&$someotherparam=2`,
			wantErr: false,
		},
		{
			name:    "should allow names with wildcard in middle common name",
			value:   "*.so*mething.com,bla.com",
			wantErr: false,
		},
		{
			name:    "should allow comma separated query params",
			value:   "https://oauth.example/oauth2/auth?allowed_groups=gid1,gid2",
			wantErr: false,
		},
		{
			name:    "should deny names with weird characters",
			value:   "something.com,lolo;xpto.com,nothing.com",
			wantErr: true,
		},
		{
			name:    "should deny names with malicous chars",
			value:   "http://something.com/#;\nournewinjection",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateArrayOfServerName(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateArrayOfServerName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkAnnotation(t *testing.T) {
	type args struct {
		name   string
		ing    *networking.Ingress
		fields AnnotationFields
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "null ingress should error",
			want: "",
			args: args{
				name: "some-random-annotation",
			},
			wantErr: true,
		},
		{
			name: "not having a validator for a specific annotation is a bug",
			want: "",
			args: args{
				name: "some-new-invalid-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							GetAnnotationWithPrefix("some-new-invalid-annotation"): "xpto",
						},
					},
				},
				fields: AnnotationFields{
					"otherannotation": AnnotationConfig{
						Validator: func(_ string) error { return nil },
					},
				},
			},
			wantErr: true,
		},
		{
			name: "annotationconfig found and no validation func defined on annotation is a bug",
			want: "",
			args: args{
				name: "some-new-invalid-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							GetAnnotationWithPrefix("some-new-invalid-annotation"): "xpto",
						},
					},
				},
				fields: AnnotationFields{
					"some-new-invalid-annotation": AnnotationConfig{},
				},
			},
			wantErr: true,
		},
		{
			name: "no annotation can turn into a null pointer and should fail",
			want: "",
			args: args{
				name: "some-new-invalid-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{},
				},
				fields: AnnotationFields{
					"some-new-invalid-annotation": AnnotationConfig{},
				},
			},
			wantErr: true,
		},
		{
			name: "no AnnotationField config should bypass validations",
			want: GetAnnotationWithPrefix("some-valid-annotation"),
			args: args{
				name: "some-valid-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							GetAnnotationWithPrefix("some-valid-annotation"): "xpto",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "annotation with invalid value should fail",
			want: "",
			args: args{
				name: "some-new-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							GetAnnotationWithPrefix("some-new-annotation"): "xpto1",
						},
					},
				},
				fields: AnnotationFields{
					"some-new-annotation": AnnotationConfig{
						Validator: func(value string) error {
							if value != "xpto" {
								return fmt.Errorf("this is an error")
							}
							return nil
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "annotation with valid value should pass",
			want: GetAnnotationWithPrefix("some-other-annotation"),
			args: args{
				name: "some-other-annotation",
				ing: &networking.Ingress{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							GetAnnotationWithPrefix("some-other-annotation"): "xpto",
						},
					},
				},
				fields: AnnotationFields{
					"some-other-annotation": AnnotationConfig{
						Validator: func(value string) error {
							if value != "xpto" {
								return fmt.Errorf("this is an error")
							}
							return nil
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkAnnotation(tt.args.name, tt.args.ing, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkAnnotation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckAnnotationRisk(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		maxrisk     AnnotationRisk
		config      AnnotationFields
		wantErr     bool
	}{
		{
			name:    "high risk should not be accepted with maximum medium",
			maxrisk: AnnotationRiskMedium,
			annotations: map[string]string{
				"nginx.ingress.kubernetes.io/bla": "blo",
				"nginx.ingress.kubernetes.io/bli": "bl3",
			},
			config: AnnotationFields{
				"bla": {
					Risk: AnnotationRiskHigh,
				},
				"bli": {
					Risk: AnnotationRiskMedium,
				},
			},
			wantErr: true,
		},
		{
			name:    "high risk should  be accepted with maximum critical",
			maxrisk: AnnotationRiskCritical,
			annotations: map[string]string{
				"nginx.ingress.kubernetes.io/bla": "blo",
				"nginx.ingress.kubernetes.io/bli": "bl3",
			},
			config: AnnotationFields{
				"bla": {
					Risk: AnnotationRiskHigh,
				},
				"bli": {
					Risk: AnnotationRiskMedium,
				},
			},
			wantErr: false,
		},
		{
			name:    "low risk should  be accepted with maximum low",
			maxrisk: AnnotationRiskLow,
			annotations: map[string]string{
				"nginx.ingress.kubernetes.io/bla": "blo",
				"nginx.ingress.kubernetes.io/bli": "bl3",
			},
			config: AnnotationFields{
				"bla": {
					Risk: AnnotationRiskLow,
				},
				"bli": {
					Risk: AnnotationRiskLow,
				},
			},
			wantErr: false,
		},
		{
			name:    "critical risk should  be accepted with maximum critical",
			maxrisk: AnnotationRiskCritical,
			annotations: map[string]string{
				"nginx.ingress.kubernetes.io/bla": "blo",
				"nginx.ingress.kubernetes.io/bli": "bl3",
			},
			config: AnnotationFields{
				"bla": {
					Risk: AnnotationRiskCritical,
				},
				"bli": {
					Risk: AnnotationRiskCritical,
				},
			},
			wantErr: false,
		},
		{
			name:    "annotation aliases should be considered in risk evaluation",
			maxrisk: AnnotationRiskLow,
			annotations: map[string]string{
				"nginx.ingress.kubernetes.io/alias": "value",
			},
			config: AnnotationFields{
				"annotation": {
					Risk:              AnnotationRiskCritical,
					AnnotationAliases: []string{"alias"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAnnotationRisk(tt.annotations, tt.maxrisk, tt.config); (err != nil) != tt.wantErr {
				t.Errorf("CheckAnnotationRisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommonNameAnnotationValidator(t *testing.T) {
	tests := []struct {
		name       string
		annotation string
		wantErr    bool
	}{
		{
			name:       "correct example",
			annotation: `CN=(my\.common\.name)`,
			wantErr:    false,
		},
		{
			name:       "no CN= prefix",
			annotation: `(my\.common\.name)`,
			wantErr:    true,
		},
		{
			name:       "invalid prefix",
			annotation: `CN(my\.common\.name)`,
			wantErr:    true,
		},
		{
			name:       "invalid regex",
			annotation: `CN=(my\.common\.name]`,
			wantErr:    true,
		},
		{
			name:       "wildcard regex",
			annotation: `CN=(my\..*\.name)`,
			wantErr:    false,
		},
		{
			name:       "somewhat complex regex",
			annotation: "CN=(my\\.app\\.dev|.*\\.bbb\\.aaaa\\.tld)",
			wantErr:    false,
		},
		{
			name:       "another somewhat complex regex",
			annotation: `CN=(my-app.*\.c\.defg\.net|other.app.com)`,
			wantErr:    false,
		},
		{
			name:       "nested parenthesis regex",
			annotation: `CN=(api-one\.(asdf)?qwer\.webpage\.organization\.org)`,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CommonNameAnnotationValidator(tt.annotation); (err != nil) != tt.wantErr {
				t.Errorf("CommonNameAnnotationValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"plain alphanumeric", "abc123", true},
		{"with dash", "my-host", true},
		{"with dot", "example.com", true},
		{"with underscore", "my_host", true},
		{"with tilde", "~^www", true},
		{"with slash and colon", "http://host", true},
		{"regex star", ".*", true},
		{"regex plus", "a+", true},
		{"regex question mark", "a?", true},
		{"regex brackets", "[0-9]", true},
		{"regex parens", "(abc)", true},
		{"regex curly braces", "{1}", true},
		{"regex pipe", "a|b", true},
		{"regex caret and dollar", "^abc$", true},
		{"regex backslash", `\.`, true},
		{"regex ampersand equals", "&=", true},
		{"empty string", "", true},
		{"with space", "hello world", false},
		{"with comma", "a,b", false},
		{"with at sign", "user@host", false},
		{"with exclamation", "a!b", false},
		{"with hash", "a#b", false},
		{"with percent", "100%", false},
		{"with semicolon", "a;b", false},
		{"with newline", "a\nb", false},
		{"with carriage return", "a\rb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, IsValidRegex.MatchString(tt.input))
		})
	}
}

func TestSizeRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"plain number", "100", true},
		{"kilobyte lower", "100k", true},
		{"kilobyte upper", "100K", true},
		{"megabyte lower", "50m", true},
		{"megabyte upper", "50M", true},
		{"gigabyte lower", "1g", true},
		{"gigabyte upper", "1G", true},
		{"byte suffix", "2048b", true},
		{"byte upper suffix", "2048B", true},
		{"large number", "999999", true},
		{"single digit", "1", true},
		{"empty string", "", false},
		{"only letter", "k", false},
		{"double suffix", "100kk", false},
		{"wrong suffix", "100x", false},
		{"decimal number", "1.5m", false},
		{"negative number", "-1", false},
		{"with space", "100 m", false},
		{"letters only", "abc", false},
		{"suffix kb", "100kb", false},
		{"zero with suffix", "0k", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, SizeRegex.MatchString(tt.input))
		})
	}
}

func TestURLIsValidRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"simple url", "http://example.com", true},
		{"url with path", "http://example.com/path", true},
		{"url with port", "http://example.com:8080", true},
		{"url with query", "http://example.com?key=value", true},
		{"url with multiple queries", "http://example.com?a=1&b=2", true},
		{"url with comma", "http://example.com?a=1,2", true},
		{"url with colon slash", "http://host:port/path", true},
		{"url with tilde", "http://example.com/~user", true},
		{"url with dash and underscore", "http://my-host_example.com", true},
		{"url with dot", "http://sub.example.com", true},
		{"empty string", "", true},
		{"with at sign", "http://user@host.com", false},
		{"with exclamation", "http://host.com/a!b", false},
		{"with hash", "http://host.com/#frag", false},
		{"with square bracket", "http://host[0].com", false},
		{"with newline", "http://host.com/\npath", false},
		{"with carriage return", "http://host.com/\rpath", false},
		{"with asterisk", "http://host.com/*", false},
		{"with space", "http://host.com/a b", false},
		{"with percent encoding", "http://host.com/%20", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, URLIsValidRegex.MatchString(tt.input))
		})
	}
}

func TestBasicCharsRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"alphanumeric", "abc123", true},
		{"with dash", "my-host", true},
		{"with dot", "example.com", true},
		{"with underscore", "my_host", true},
		{"with tilde", "~value", true},
		{"with slash", "namespace/name", true},
		{"with colon", "host:port", true},
		{"empty string", "", true},
		{"mixed valid chars", "ns/my-app_v2:8080", true},
		{"path with tilde", "/~user/path", true},
		{"all special valid chars", "a.b-c_d~e/f:g", true},
		{"with space", "hello world", false},
		{"with comma", "a,b", false},
		{"with at sign", "user@host", false},
		{"with exclamation", "a!b", false},
		{"with hash", "a#b", false},
		{"with dollar", "$var", false},
		{"with asterisk", "a*", false},
		{"with question mark", "a?", false},
		{"with semicolon", "a;b", false},
		{"with newline", "a\nb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, BasicCharsRegex.MatchString(tt.input))
		})
	}
}

func TestExtendedCharsRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"alphanumeric", "abc123", true},
		{"with dash", "my-host", true},
		{"with dot", "example.com", true},
		{"with comma", "host1,host2", true},
		{"with space", "hello world", true},
		{"with comma and space", "a, b, c", true},
		{"mixed valid", "ns/my-app, other", true},
		{"with colon", "host:port, other", true},
		{"with slash", "ns/name, other/name", true},
		{"empty string", "", true},
		{"with at sign", "user@host", false},
		{"with exclamation", "a!b", false},
		{"with hash", "a#b", false},
		{"with dollar", "$var", false},
		{"with asterisk", "a*", false},
		{"with newline", "a\nb", false},
		{"with carriage return", "a\rb", false},
		{"with semicolon", "a;b", false},
		{"with percent", "100%", false},
		{"with ampersand", "a&b", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, ExtendedCharsRegex.MatchString(tt.input))
		})
	}
}

func TestCharsWithSpace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"alphanumeric", "abc123", true},
		{"with dash", "my-host", true},
		{"with dot", "example.com", true},
		{"with underscore", "my_host", true},
		{"with slash", "ns/name", true},
		{"with colon", "host:port", true},
		{"with space", "hello world", true},
		{"multiple spaces", "a b c d", true},
		{"with tilde", "~ value", true},
		{"empty string", "", true},
		{"with comma", "a,b", false},
		{"with at sign", "user@host", false},
		{"with exclamation", "a!b", false},
		{"with hash", "a#b", false},
		{"with dollar", "$var", false},
		{"with asterisk", "a*", false},
		{"with newline", "a\nb", false},
		{"with semicolon", "a;b", false},
		{"with ampersand", "a&b", false},
		{"with percent", "100%", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, CharsWithSpace.MatchString(tt.input))
		})
	}
}

func TestNGINXVariable(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"simple variable", "$host", true},
		{"braced variable", "${host}", true},
		{"http header variable", "$http_header", true},
		{"alphanumeric only", "abc123", true},
		{"with dash", "my-var", true},
		{"with underscore", "my_var", true},
		{"dollar at end", "value$", true},
		{"multiple variables", "${var1}${var2}", true},
		{"complex braced", "${http_x_forwarded_for}", true},
		{"empty string", "", true},
		{"with slash", "ns/name", false},
		{"with space", "a b", false},
		{"with dot", "a.b", false},
		{"with at sign", "a@b", false},
		{"with colon", "host:port", false},
		{"with tilde", "~var", false},
		{"with comma", "a,b", false},
		{"with newline", "a\nb", false},
		{"with asterisk", "a*", false},
		{"with hash", "#var", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, NGINXVariable.MatchString(tt.input))
		})
	}
}

func TestRegexPathWithCapture(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"simple path", "/path/to/resource", true},
		{"path with capture", "/something/$1", true},
		{"path with multiple captures", "/a/$1/b/$2", true},
		{"path with dollar only", "/path/$", true},
		{"no leading slash", "path/to/resource", true},
		{"with dot", "/file.html", true},
		{"with dash", "/my-path", true},
		{"with underscore", "/my_path", true},
		{"with tilde", "/~user/path", true},
		{"with colon", "/path:name", true},
		{"empty string", "", true},
		{"with space", "/path to/file", false},
		{"with comma", "/path,file", false},
		{"with at sign", "/path@host", false},
		{"with exclamation", "/path!", false},
		{"with hash", "/path#frag", false},
		{"with percent", "/path%20", false},
		{"with newline", "/path\nfile", false},
		{"with semicolon", "/path;cmd", false},
		{"with asterisk and brackets", "/*.html", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, RegexPathWithCapture.MatchString(tt.input))
		})
	}
}

func TestHeadersVariable(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"single header", "Content-Type", true},
		{"multiple headers", "X-Header-1, X-Header-2", true},
		{"headers with space", "Accept, Content-Type", true},
		{"alphanumeric header", "X123", true},
		{"header with underscore", "X_Custom_Header", true},
		{"simple name", "Authorization", true},
		{"single char", "A", true},
		{"empty string", "", true},
		{"header with spaces around comma", "H1 , H2 , H3", true},
		{"all lowercase", "accept, content-type", true},
		{"with slash", "Content/Type", false},
		{"with dollar", "$header", false},
		{"with colon", "Content:Type", false},
		{"with at sign", "X@Header", false},
		{"with dot", "Header.Name", false},
		{"with newline", "Header\nValue", false},
		{"with carriage return", "Header\rValue", false},
		{"with hash", "Header#1", false},
		{"with exclamation", "Header!", false},
		{"with percent", "Header%", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, HeadersVariable.MatchString(tt.input))
		})
	}
}

func TestURLWithNginxVariableRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"plain url", "http://example.com", true},
		{"url with variable", "http://$host/path", true},
		{"url with braced variable", "http://host/$path", true},
		{"url with query", "http://host?key=$val", true},
		{"url with comma", "http://host?a=1,2", true},
		{"url with space", "http://host/path to", true},
		{"url with equals", "http://host?key=value", true},
		{"url with ampersand", "http://host?a=1&b=2", true},
		{"url with port", "http://host:8080", true},
		{"empty string", "", true},
		{"with at sign", "http://user@host", false},
		{"with exclamation", "http://host!", false},
		{"with hash", "http://host#frag", false},
		{"with asterisk", "http://host/*", false},
		{"with newline", "http://host\npath", false},
		{"with carriage return", "http://host\rpath", false},
		{"with percent", "http://host/%20", false},
		{"with semicolon", "http://host;cmd", false},
		{"with square bracket", "http://host[0]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, URLWithNginxVariableRegex.MatchString(tt.input))
		})
	}
}

func TestMaliciousRegex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{"carriage return", "value\r", true},
		{"newline", "value\n", true},
		{"carriage return newline", "value\r\n", true},
		{"embedded carriage return", "val\rue", true},
		{"embedded newline", "val\nue", true},
		{"only carriage return", "\r", true},
		{"only newline", "\n", true},
		{"clean string", "clean-value", false},
		{"clean url", "http://example.com", false},
		{"clean alphanumeric", "abc123", false},
		{"empty string", "", false},
		{"clean with special chars", "host:port/path", false},
		{"clean with comma", "a,b,c", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.match, MaliciousRegex.MatchString(tt.input))
		})
	}
}

func TestValidateRegex(t *testing.T) {
	validator := ValidateRegex(BasicCharsRegex, false)

	t.Run("valid basic chars", func(t *testing.T) {
		assert.NoError(t, validator("host:port"))
	})

	t.Run("invalid chars", func(t *testing.T) {
		assert.Error(t, validator("value!invalid"))
	})

	t.Run("malicious with carriage return", func(t *testing.T) {
		assert.Error(t, validator("value\rinjection"))
	})

	t.Run("malicious with newline", func(t *testing.T) {
		assert.Error(t, validator("value\ninjection"))
	})

	t.Run("empty string passes regex match", func(t *testing.T) {
		assert.NoError(t, validator(""))
	})

	validatorRemoveSpace := ValidateRegex(CharsWithSpace, true)

	t.Run("remove space enabled strips spaces before match", func(t *testing.T) {
		assert.NoError(t, validatorRemoveSpace("a b c"))
	})

	validatorNoRemove := ValidateRegex(BasicCharsRegex, false)

	t.Run("spaces not removed when removeSpace false", func(t *testing.T) {
		assert.Error(t, validatorNoRemove("a b c"))
	})
}

func TestValidateOptions(t *testing.T) {
	options := []string{"enabled", "disabled", "auto"}

	t.Run("case sensitive match", func(t *testing.T) {
		validator := ValidateOptions(options, true, false)
		assert.NoError(t, validator("enabled"))
	})

	t.Run("case sensitive no match", func(t *testing.T) {
		validator := ValidateOptions(options, true, false)
		assert.Error(t, validator("Enabled"))
	})

	t.Run("case insensitive match", func(t *testing.T) {
		validator := ValidateOptions(options, false, false)
		assert.NoError(t, validator("Enabled"))
	})

	t.Run("case insensitive uppercase match", func(t *testing.T) {
		validator := ValidateOptions(options, false, false)
		assert.NoError(t, validator("DISABLED"))
	})

	t.Run("no match returns error", func(t *testing.T) {
		validator := ValidateOptions(options, false, false)
		assert.Error(t, validator("unknown"))
	})

	t.Run("trim space enabled", func(t *testing.T) {
		validator := ValidateOptions(options, true, true)
		assert.NoError(t, validator("  enabled  "))
	})

	t.Run("trim space disabled with spaces", func(t *testing.T) {
		validator := ValidateOptions(options, true, false)
		assert.Error(t, validator(" enabled "))
	})

	t.Run("empty string no match", func(t *testing.T) {
		validator := ValidateOptions(options, false, false)
		assert.Error(t, validator(""))
	})

	t.Run("case insensitive and trim space", func(t *testing.T) {
		validator := ValidateOptions(options, false, true)
		assert.NoError(t, validator("  ENABLED  "))
	})
}

func TestValidateBool(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"true", "true", false},
		{"false", "false", false},
		{"True", "True", false},
		{"False", "False", false},
		{"TRUE", "TRUE", false},
		{"FALSE", "FALSE", false},
		{"one", "1", false},
		{"zero", "0", false},
		{"yes", "yes", true},
		{"no", "no", true},
		{"Yes", "Yes", true},
		{"No", "No", true},
		{"T", "T", false},
		{"F", "F", false},
		{"t", "t", false},
		{"f", "f", false},
		{"empty string", "", true},
		{"random string", "maybe", true},
		{"number 2", "2", true},
		{"string truee", "truee", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBool(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateInt(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"zero", "0", false},
		{"positive", "123", false},
		{"negative", "-456", false},
		{"large number", "999999999", false},
		{"single digit", "5", false},
		{"empty string", "", true},
		{"decimal", "1.5", true},
		{"letters", "abc", true},
		{"with space", "12 34", true},
		{"hex prefix", "0xff", true},
		{"float notation", "1e10", true},
		{"boolean string", "true", true},
		{"special chars", "12!3", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInt(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateCIDRs(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"single ipv4 cidr", "10.0.0.0/8", false},
		{"single ipv4 address", "192.168.1.1", false},
		{"multiple cidrs", "10.0.0.0/8,172.16.0.0/12", false},
		{"ipv6 cidr", "2001:db8::/32", false},
		{"ipv6 address", "::1", false},
		{"mixed ipv4 and ipv6", "10.0.0.0/8,2001:db8::/32", false},
		{"single host cidr", "10.0.0.1/32", false},
		{"empty string", "", false},
		{"invalid cidr", "999.0.0.0/8", true},
		{"invalid notation", "not-a-cidr", true},
		{"garbage mixed", "10.0.0.0/8,garbage", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCIDRs(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateDuration(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"seconds", "30s", false},
		{"minutes", "5m", false},
		{"hours", "1h", false},
		{"milliseconds", "500ms", false},
		{"microseconds", "100us", false},
		{"nanoseconds", "100ns", false},
		{"combined", "1h30m", false},
		{"combined full", "2h30m10s", false},
		{"zero", "0s", false},
		{"zero no suffix", "0", false},
		{"empty string", "", true},
		{"no unit", "100", true},
		{"invalid unit", "100x", true},
		{"negative", "-5m", false},
		{"float without unit", "1.5", true},
		{"spaces", "5 m", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDuration(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateServiceName(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"simple name", "myservice", false},
		{"with dash", "my-service", false},
		{"with numbers", "service123", false},
		{"starts with letter", "a123", false},
		{"single char", "a", false},
		{"starts with number", "1service", true},
		{"starts with dash", "-service", true},
		{"ends with dash", "service-", true},
		{"contains dot", "my.service", true},
		{"contains underscore", "my_service", true},
		{"empty string", "", true},
		{"too long name", "a1234567890123456789012345678901234567890123456789012345678901234", true},
		{"uppercase letters", "MyService", true},
		{"special chars", "my$service", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateServiceName(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringRiskToRisk(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect AnnotationRisk
	}{
		{"low", "low", AnnotationRiskLow},
		{"medium", "medium", AnnotationRiskMedium},
		{"high", "high", AnnotationRiskHigh},
		{"critical", "critical", AnnotationRiskCritical},
		{"uppercase low", "Low", AnnotationRiskLow},
		{"uppercase medium", "Medium", AnnotationRiskMedium},
		{"uppercase high", "High", AnnotationRiskHigh},
		{"uppercase critical", "Critical", AnnotationRiskCritical},
		{"all caps low", "LOW", AnnotationRiskLow},
		{"all caps medium", "MEDIUM", AnnotationRiskMedium},
		{"all caps high", "HIGH", AnnotationRiskHigh},
		{"all caps critical", "CRITICAL", AnnotationRiskCritical},
		{"unknown defaults to low", "unknown", AnnotationRiskLow},
		{"empty defaults to low", "", AnnotationRiskLow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, StringRiskToRisk(tt.input))
		})
	}
}

func TestValidateNull(t *testing.T) {
	t.Run("always returns nil for any input", func(t *testing.T) {
		assert.NoError(t, ValidateNull(""))
		assert.NoError(t, ValidateNull("anything"))
		assert.NoError(t, ValidateNull("malicious\r\ninput"))
	})
}

func TestValidateServerName(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"simple hostname", "example.com", false},
		{"with wildcard star", "*.example.com", false},
		{"regex characters", `~^www\d+\.example\.com$`, false},
		{"with path", "example.com/path", false},
		{"with port", "example.com:8080", false},
		{"with spaces around gets trimmed", "  example.com  ", false},
		{"malicious with newline", "example.com\ninjection", true},
		{"malicious with semicolon", "example.com;injection", true},
		{"with at sign", "user@host.com", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateServerName(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAnnotationRiskToString(t *testing.T) {
	tests := []struct {
		name   string
		risk   AnnotationRisk
		expect string
	}{
		{"low", AnnotationRiskLow, "Low"},
		{"medium", AnnotationRiskMedium, "Medium"},
		{"high", AnnotationRiskHigh, "High"},
		{"critical", AnnotationRiskCritical, "Critical"},
		{"unknown value", AnnotationRisk(99), "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.risk.ToString())
		})
	}
}

func TestValidateRegexWithVariousRegexes(t *testing.T) {
	t.Run("SizeRegex valid", func(t *testing.T) {
		validator := ValidateRegex(SizeRegex, false)
		assert.NoError(t, validator("100m"))
	})

	t.Run("SizeRegex invalid", func(t *testing.T) {
		validator := ValidateRegex(SizeRegex, false)
		assert.Error(t, validator("abc"))
	})

	t.Run("NGINXVariable valid", func(t *testing.T) {
		validator := ValidateRegex(NGINXVariable, false)
		assert.NoError(t, validator("$host"))
	})

	t.Run("NGINXVariable with malicious input", func(t *testing.T) {
		validator := ValidateRegex(NGINXVariable, false)
		assert.Error(t, validator("$host\r"))
	})

	t.Run("HeadersVariable valid", func(t *testing.T) {
		validator := ValidateRegex(HeadersVariable, false)
		assert.NoError(t, validator("X-Custom-Header"))
	})

	t.Run("HeadersVariable invalid", func(t *testing.T) {
		validator := ValidateRegex(HeadersVariable, false)
		assert.Error(t, validator("X-Custom:Header"))
	})

	t.Run("removeSpace with CharsWithSpace", func(t *testing.T) {
		validator := ValidateRegex(BasicCharsRegex, true)
		assert.NoError(t, validator("a b c"))
	})

	t.Run("compiled regex passed through", func(t *testing.T) {
		custom := regexp.MustCompile(`^[a-z]+$`)
		validator := ValidateRegex(custom, false)
		assert.NoError(t, validator("abc"))
		assert.Error(t, validator("ABC"))
	})
}
