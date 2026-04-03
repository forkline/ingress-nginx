package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	RELEASE = "test"
	REPO = "test-repo"
	COMMIT = "abc123"

	result := String()

	if !strings.Contains(result, "test") {
		t.Errorf("String() should contain release version")
	}
	if !strings.Contains(result, "test-repo") {
		t.Errorf("String() should contain repo")
	}
	if !strings.Contains(result, "abc123") {
		t.Errorf("String() should contain commit")
	}
}
