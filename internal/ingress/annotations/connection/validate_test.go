package connection

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
			c1:     &Config{Header: "close", Enabled: true},
			c2:     nil,
			expect: false,
		},
		{
			name:   "identical",
			c1:     &Config{Header: "close", Enabled: true},
			c2:     &Config{Header: "close", Enabled: true},
			expect: true,
		},
		{
			name:   "different header",
			c1:     &Config{Header: "close"},
			c2:     &Config{Header: "keep-alive"},
			expect: false,
		},
		{
			name:   "different enabled",
			c1:     &Config{Enabled: true},
			c2:     &Config{Enabled: false},
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
			name:      "valid annotation passes with low risk",
			riskLevel: "low",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/connection-proxy-header": "close"},
			wantErr:   false,
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
	assert.Contains(t, fields, "connection-proxy-header")
}
