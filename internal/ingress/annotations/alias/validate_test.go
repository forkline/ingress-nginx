package alias

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func TestValidate(t *testing.T) {
	p := NewParser(&resolver.Mock{AnnotationsRiskLevel: "high"})

	tests := []struct {
		name    string
		anns    map[string]string
		wantErr bool
	}{
		{
			name:    "empty annotations pass",
			anns:    map[string]string{},
			wantErr: false,
		},
		{
			name:    "valid alias annotation passes with high risk",
			anns:    map[string]string{"nginx.ingress.kubernetes.io/server-alias": "example.com"},
			wantErr: false,
		},
		{
			name:    "valid alias annotation fails with low risk",
			anns:    map[string]string{"nginx.ingress.kubernetes.io/server-alias": "example.com"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				p = NewParser(&resolver.Mock{AnnotationsRiskLevel: "low"})
			}
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
	assert.Contains(t, fields, "server-alias")
}
