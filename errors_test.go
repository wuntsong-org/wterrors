package errors

import "testing"

func TestIs(t *testing.T) {
	err := Errorf("New error: %d", 10)
	_ = err.SetCode("Code")

	if err.Code() != "Code" {
		t.Errorf("Error code is not set")
	}
}
