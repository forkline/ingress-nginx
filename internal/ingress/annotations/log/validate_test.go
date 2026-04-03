package log

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
			name:   "identical",
			c1:     &Config{Access: true, Rewrite: false},
			c2:     &Config{Access: true, Rewrite: false},
			expect: true,
		},
		{
			name:   "different access",
			c1:     &Config{Access: true},
			c2:     &Config{Access: false},
			expect: false,
		},
		{
			name:   "different rewrite",
			c1:     &Config{Rewrite: true},
			c2:     &Config{Rewrite: false},
			expect: false,
		},
		{
			name:   "both false",
			c1:     &Config{Access: false, Rewrite: false},
			c2:     &Config{Access: false, Rewrite: false},
			expect: true,
		},
		{
			name:   "both true",
			c1:     &Config{Access: true, Rewrite: true},
			c2:     &Config{Access: true, Rewrite: true},
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			anns:      map[string]string{"nginx.ingress.kubernetes.io/enable-access-log": "true"},
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
	assert.Contains(t, fields, "enable-access-log")
}
