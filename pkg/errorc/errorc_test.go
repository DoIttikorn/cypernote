//go:build unit

package errorc

import "testing"

func TestNewError(t *testing.T) {
	err := New(400, "test")
	if err == nil {
		t.Errorf("error is nil")
	}
}

func TestNewErrorWithNil(t *testing.T) {
	err := New(400, "")
	if err == nil {
		t.Errorf("error is nil")
	}
}

func TestResponseErr(t *testing.T) {
	got := "test"
	err := ResponseErr(got)
	if err.Error() != got {
		t.Errorf("error is not equal %s", got)
	}
}

func TestErrorMessage(t *testing.T) {
	got := "test"
	err := New(400, got)
	if err.Error() != got {
		t.Errorf("error is not equal %s", got)
	}
}
