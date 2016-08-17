package ctci

import "testing"

func TestIsUnique(t *testing.T) {
	// Arrange
	cases := []struct {
		input  string
		expect bool
	}{
		{"", true},
		{"a", true},
		{"abc", true},
		{"./,", true},
		{"aa", false},
		{"abvda", false},
	}

	for _, it := range cases {
		// Act
		res := IsUnique(it.input)

		// Assert
		if res != it.expect {
			t.Errorf("IsUnique: input %q, expect %v, but got %v\n", it.input, it.expect, res)
		}
	}
}

func TestReverse(t *testing.T) {
	//Arrange
	cases := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"abcde", "edcba"},
	}

	for _, it := range cases {
		// Act
		res := Reverse(it.input)

		if res != it.expect {
			t.Errorf("Reverse: input %q, expect %q, but got %q\n", it.input, it.expect, res)
		}
	}
}
