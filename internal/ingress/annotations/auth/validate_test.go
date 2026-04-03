package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func TestConfigEqual(t *testing.T) {
	tests := []struct {
		name   string
		c1     *Config
		c2     *Config
		expect bool
	}{
		{
			name:   "both nil",
			c1:     nil,
			c2:     nil,
			expect: true,
		},
		{
			name:   "one nil",
			c1:     &Config{Type: "basic"},
			c2:     nil,
			expect: false,
		},
		{
			name:   "other nil",
			c1:     nil,
			c2:     &Config{Type: "basic"},
			expect: false,
		},
		{
			name:   "identical",
			c1:     &Config{Type: "basic", Realm: "realm", File: "/path", Secured: true, FileSHA: "sha", Secret: "ns/secret", SecretType: "auth-file"},
			c2:     &Config{Type: "basic", Realm: "realm", File: "/path", Secured: true, FileSHA: "sha", Secret: "ns/secret", SecretType: "auth-file"},
			expect: true,
		},
		{
			name:   "different type",
			c1:     &Config{Type: "basic"},
			c2:     &Config{Type: "digest"},
			expect: false,
		},
		{
			name:   "different realm",
			c1:     &Config{Realm: "realm1"},
			c2:     &Config{Realm: "realm2"},
			expect: false,
		},
		{
			name:   "different file",
			c1:     &Config{File: "/path1"},
			c2:     &Config{File: "/path2"},
			expect: false,
		},
		{
			name:   "different secured",
			c1:     &Config{Secured: true},
			c2:     &Config{Secured: false},
			expect: false,
		},
		{
			name:   "different filesha",
			c1:     &Config{FileSHA: "sha1"},
			c2:     &Config{FileSHA: "sha2"},
			expect: false,
		},
		{
			name:   "different secret",
			c1:     &Config{Secret: "ns/sec1"},
			c2:     &Config{Secret: "ns/sec2"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "both nil" {
				var c1, c2 *Config
				assert.Equal(t, true, c1.Equal(c2))
				return
			}
			assert.Equal(t, tt.expect, tt.c1.Equal(tt.c2))
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		riskLevel string
		anns      map[string]string
		wantErr   bool
	}{
		{
			name:      "empty annotations pass",
			riskLevel: "high",
			anns:      map[string]string{},
			wantErr:   false,
		},
		{
			name:      "medium risk auth-realm passes with medium",
			riskLevel: "medium",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/auth-realm": "Restricted"},
			wantErr:   false,
		},
		{
			name:      "medium risk auth-realm fails with low",
			riskLevel: "low",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/auth-realm": "Restricted"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser("/tmp", &resolver.Mock{AnnotationsRiskLevel: tt.riskLevel})
			err := p.Validate(tt.anns)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetDocumentation(t *testing.T) {
	p := NewParser("/tmp", &resolver.Mock{})
	fields := p.GetDocumentation()
	assert.NotNil(t, fields)
	assert.Contains(t, fields, "auth-type")
}
