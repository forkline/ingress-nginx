package snippet

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
			anns:      map[string]string{"nginx.ingress.kubernetes.io/configuration-snippet": "more_set_headers \"X-Test: test\";"},
			wantErr:   false,
		},
		{
			name:      "annotation fails with high risk",
			riskLevel: "high",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/configuration-snippet": "more_set_headers \"X-Test: test\";"},
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
	assert.Contains(t, fields, "configuration-snippet")
}
