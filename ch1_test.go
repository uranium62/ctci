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

func TestCompress(t *testing.T) {
	//Arrange
	cases := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "abc"},
		{"aaaaa", "a5"},
		{"abccccc", "a1b1c5"},
		{"aaabbbccc", "a3b3c3"},
	}

	for _, it := range cases {
		// Act
		res := Compress(it.input)

		// Assert
		if res != it.expect {
			t.Errorf("Compress: input %q, expect %q, but got %q\n", it.input, it.expect, res)
		}
	}
}

func TestMatrixNN_Rotate(t *testing.T) {
	cases := []struct {
		input  MatrixNN
		expect MatrixNN
	}{
		{MatrixNN{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			MatrixNN{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{MatrixNN{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			MatrixNN{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}},
	}

	for _, it := range cases {
		input := it.input
		it.input.Rotate()
		res := it.input
		if !res.Compare(it.expect) {
			t.Errorf("MatrixNN_Rotate: \ninput %v \nexpected %v\n but got %v \n", input, it.expect, res)
		}
	}
}

func TestMatrixMN_Zero(t *testing.T) {
	cases := []struct {
		input  MatrixMN
		expect MatrixMN
	}{
		{MatrixMN{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}},
			MatrixMN{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}}},
		{MatrixMN{{1, 2, 3, 0}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			MatrixMN{{0, 0, 0, 0}, {5, 6, 7, 0}, {9, 10, 11, 0}}},
	}

	for _, it := range cases {
		input := it.input
		it.input.Zero()
		res := it.input
		if !res.Compare(it.expect) {
			t.Errorf("MatrixMN_Zero: \ninput %v \nexpected %v\n but got %v \n", input, it.expect, res)
		}
	}
}

func TestIsRotation(t *testing.T) {
	cases := []struct {
		input  []string
		expect bool
	}{
		{[]string{"apple", "pleap"}, true},
		{[]string{"waterbottle", "erbottlewat"}, true},
		{[]string{"camera", "macera"}, false},
		{[]string{"santhosh", "@santhosh"}, false},
	}

	for _, it := range cases {
		input := it.input
		res := IsRotation(input[0], input[1])
		if res != it.expect {
			t.Errorf("IsRotation: input %q expected %v but got %v", input, it.expect, res)
		}
	}
}
