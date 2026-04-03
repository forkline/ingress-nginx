package streamsnippet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		riskLevel string
		anns      map[string]string
		wantErr   bool
	}{
		{
			name:      "empty annotations pass",
			riskLevel: "critical",
			anns:      map[string]string{},
			wantErr:   false,
		},
		{
			name:      "valid annotation passes with critical risk",
			riskLevel: "critical",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/stream-snippet": "return 200"},
			wantErr:   false,
		},
		{
			name:      "annotation fails with high risk",
			riskLevel: "high",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/stream-snippet": "return 200"},
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
	assert.Contains(t, fields, "stream-snippet")
}
