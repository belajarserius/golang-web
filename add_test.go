package golang_web

import "testing"

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5

// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
// 	}
// }

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add mix of positive and negative", 1, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}