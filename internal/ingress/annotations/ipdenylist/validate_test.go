package ipdenylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/pkg/util/sets"
)

func TestSourceRangeEqual(t *testing.T) {
	tests := []struct {
		name   string
		sr1    *SourceRange
		sr2    *SourceRange
		expect bool
	}{
		{
			name:   "both nil",
			sr1:    nil,
			sr2:    nil,
			expect: true,
		},
		{
			name:   "one nil",
			sr1:    &SourceRange{CIDR: []string{"10.0.0.0/8"}},
			sr2:    nil,
			expect: false,
		},
		{
			name:   "identical",
			sr1:    &SourceRange{CIDR: []string{"10.0.0.0/8"}},
			sr2:    &SourceRange{CIDR: []string{"10.0.0.0/8"}},
			expect: true,
		},
		{
			name:   "different order same values",
			sr1:    &SourceRange{CIDR: []string{"10.0.0.0/8", "172.16.0.0/12"}},
			sr2:    &SourceRange{CIDR: []string{"172.16.0.0/12", "10.0.0.0/8"}},
			expect: true,
		},
		{
			name:   "different cidrs",
			sr1:    &SourceRange{CIDR: []string{"10.0.0.0/8"}},
			sr2:    &SourceRange{CIDR: []string{"192.168.0.0/16"}},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.sr1.Equal(tt.sr2))
		})
	}

	var sr1, sr2 *SourceRange
	assert.True(t, sr1.Equal(sr2), "both nil source ranges should be equal")

	sr := &SourceRange{CIDR: []string{"10.0.0.0/8"}}
	srCopy := &SourceRange{CIDR: []string{"10.0.0.0/8"}}
	assert.True(t, sr.Equal(srCopy), "source range should equal identical copy")
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
			anns:      map[string]string{"nginx.ingress.kubernetes.io/denylist-source-range": "10.0.0.0/8"},
			wantErr:   false,
		},
		{
			name:      "annotation fails with low risk",
			riskLevel: "low",
			anns:      map[string]string{"nginx.ingress.kubernetes.io/denylist-source-range": "10.0.0.0/8"},
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
	assert.Contains(t, fields, "denylist-source-range")
}

func TestStringElementsMatch(t *testing.T) {
	tests := []struct {
		name   string
		a      []string
		b      []string
		expect bool
	}{
		{
			name:   "both empty",
			a:      []string{},
			b:      []string{},
			expect: true,
		},
		{
			name:   "same elements different order",
			a:      []string{"b", "a"},
			b:      []string{"a", "b"},
			expect: true,
		},
		{
			name:   "different elements",
			a:      []string{"a"},
			b:      []string{"b"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, sets.StringElementsMatch(tt.a, tt.b))
		})
	}
}
