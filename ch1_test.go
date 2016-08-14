package ctci

import "testing"

func TestIsUnique(t *testing.T) {

	// Arrange
	cases := map[string]bool{
		"":      true,
		"a":     true,
		"abc":   true,
		"./,":   true,
		"aa":    false,
		"abvda": false,
	}

	for input, expect := range cases {
		// Act
		result := IsUnique(input)

		// Assert
		if result != expect {
			t.Errorf("IsUnique: input %q, expect %v, but got %v\n", input, expect, result)
		}
	}
}
