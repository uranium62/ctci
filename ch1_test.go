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

		// Assert
		if res != it.expect {
			t.Errorf("Reverse: input %q, expect %q, but got %q\n", it.input, it.expect, res)
		}
	}
}

func TestIsPermutation(t *testing.T) {
	// Arrange
	cases := []struct {
		str1   string
		str2   string
		expect bool
	}{
		{"", "", true},
		{"", ".", false},
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "bac", true},
		{"abc", "abb", false},
	}

	for _, it := range cases {
		// Act
		res := IsPermutation(it.str1, it.str2)

		// Assert
		if res != it.expect {
			t.Errorf("IsPermutation: input %q %q, expect %q, but got %v\n", it.str1, it.str2, it.expect, res)
		}
	}
}

func TestReplaceSpace(t *testing.T) {
	// Arrange
	cases := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"a", "a"},
		{" ", "%20"},
		{"  ", "%20%20"},
		{" Hello world! ", "%20Hello%20world!%20"},
	}

	for _, it := range cases {
		// Act
		res := ReplaceSpace(it.input)

		// Assert
		if res != it.expect {
			t.Errorf("ReplaceSpace: input %q, expect %q, but got %q\n", it.input, it.expect, res)
		}
	}
}
