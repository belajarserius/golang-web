package golang_web

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(4, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("Divide(4, 2) = %f; want %f", result, 2.0)
	}

	_, err = Divide(4, 0)
	if err == nil {
		t.Error("expected an error but got none")
	}
}
