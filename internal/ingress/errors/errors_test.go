/*
Copyright 2017 The Kubernetes Authors.

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

package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocationDenied(t *testing.T) {
	err := NewLocationDenied("demo")
	if !IsLocationDenied(err) {
		t.Error("expected true")
	}
	if IsLocationDenied(nil) {
		t.Error("expected false")
	}
}

func TestIsMissingAnnotations(t *testing.T) {
	if !IsMissingAnnotations(ErrMissingAnnotations) {
		t.Error("expected true")
	}
}

func TestInvalidContent(t *testing.T) {
	if IsInvalidContent(ErrMissingAnnotations) {
		t.Error("expected false")
	}
	err := NewInvalidAnnotationContent("demo", "")
	if !IsInvalidContent(err) {
		t.Error("expected true")
	}
	if IsInvalidContent(nil) {
		t.Error("expected false")
	}
	err = NewLocationDenied("demo")
	if IsInvalidContent(err) {
		t.Error("expected false")
	}
}

func TestNew(t *testing.T) {
	err := New("test error")
	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())
}

func TestErrorf(t *testing.T) {
	err := Errorf("error with %s", "format")
	assert.NotNil(t, err)
	assert.Equal(t, "error with format", err.Error())
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("some-annotation")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "some-annotation")
	assert.Contains(t, err.Error(), "invalid value")
}

func TestIsValidationError(t *testing.T) {
	err := NewValidationError("test-annotation")
	assert.True(t, IsValidationError(err))
	assert.False(t, IsValidationError(nil))
	assert.False(t, IsValidationError(New("plain error")))
}

func TestValidationError_Error(t *testing.T) {
	err := NewValidationError("ann")
	assert.Equal(t, "annotation ann contains invalid value", err.Error())
}

func TestNewRiskyAnnotations(t *testing.T) {
	err := NewRiskyAnnotations("test-group")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test-group")
	assert.Contains(t, err.Error(), "risky")
}

func TestIsRiskyAnnotationError(t *testing.T) {
	assert.False(t, IsRiskyAnnotationError(nil))
	assert.False(t, IsRiskyAnnotationError(New("plain error")))
	validateErr := NewValidationError("test")
	assert.True(t, IsRiskyAnnotationError(validateErr))
}

func TestRiskyAnnotationError_Error(t *testing.T) {
	err := NewRiskyAnnotations("group")
	assert.Contains(t, err.Error(), "group")
}

func TestInvalidConfigurationError(t *testing.T) {
	err := NewInvalidAnnotationConfiguration("test-ann", "bad value")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test-ann")
	assert.Contains(t, err.Error(), "bad value")
}

func TestInvalidContentError(t *testing.T) {
	err := NewInvalidAnnotationContent("test-ann", "bad-val")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test-ann")
	assert.Contains(t, err.Error(), "bad-val")
}

func TestLocationDeniedError(t *testing.T) {
	err := NewLocationDenied("reason")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "reason")
}
