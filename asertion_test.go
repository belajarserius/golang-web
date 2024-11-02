package golang_web

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestAdd1(t *testing.T) {
// 	result := Add1(2, 3)
// 	expected := 5

// 	// Assertion manual
// 	if result != expected {
// 		t.Errorf("Expected %d but got %d", expected, result)
// 	}

// }

func TestAdd1(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	// Assertion menggunakan testify
	assert.Equal(t, expected, result, "The result should be 5")
}
