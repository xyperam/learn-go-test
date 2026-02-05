package integers

import "testing"

func TestAdder(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive numbers", 2, 2, 4},
		{"with zero", 5, 0, 5},
		{"negative numbers", -2, -3, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := Add(tt.x, tt.y)

			if sum != tt.expected {
				t.Errorf("expected '%d' but got '%d'", tt.expected, sum)
			}
		})
	}
}
