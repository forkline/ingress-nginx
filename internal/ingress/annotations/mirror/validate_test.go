package mirror

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
			c1:     &Config{Source: "/mirror", Target: "http://target"},
			c2:     nil,
			expect: false,
		},
		{
			name:   "identical",
			c1:     &Config{Source: "/mirror", RequestBody: "on", Target: "http://target", Host: "example.com"},
			c2:     &Config{Source: "/mirror", RequestBody: "on", Target: "http://target", Host: "example.com"},
			expect: true,
		},
		{
			name:   "different source",
			c1:     &Config{Source: "/mirror1"},
			c2:     &Config{Source: "/mirror2"},
			expect: false,
		},
		{
			name:   "different requestbody",
			c1:     &Config{RequestBody: "on"},
			c2:     &Config{RequestBody: "off"},
			expect: false,
		},
		{
			name:   "different target",
			c1:     &Config{Target: "http://a"},
			c2:     &Config{Target: "http://b"},
			expect: false,
		},
		{
			name:   "different host",
			c1:     &Config{Host: "a.com"},
			c2:     &Config{Host: "b.com"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "both nil" {
				var c1, c2 *Config
				assert.True(t, c1.Equal(c2))
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
			name:      "valid annotation passes with high risk",
			riskLevel: "high",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/mirror-target": "http://mirror"},
			wantErr:   false,
		},
		{
			name:      "annotation fails with low risk",
			riskLevel: "low",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/mirror-target": "http://mirror"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(&resolver.Mock{AnnotationsRiskLevel: tt.riskLevel})
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
	p := NewParser(&resolver.Mock{})
	fields := p.GetDocumentation()
	assert.NotNil(t, fields)
	assert.Contains(t, fields, "mirror-target")
}
