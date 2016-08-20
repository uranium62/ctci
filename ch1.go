package ctci

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}

	var c_set [128]bool

	for _, val := range str {
		if c_set[val] {
			return false
		}
		c_set[val] = true
	}
	return true
}

func Reverse(str string) string {
	lst := len(str) - 1
	mid := len(str) / 2
	arr := []byte(str)

	for i := 0; i < mid; i++ {
		arr[i], arr[lst-i] = arr[lst-i], arr[i]
	}

	return string(arr)
}

func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make([]int, 256)

	for _, c := range str1 {
		count[rune(c)]++
	}
	for _, c := range str2 {
		ch := rune(c)
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func ReplaceSpace(str string) string {
	src := []byte(str)
	length := len(src)
	count := 0

	for i := 0; i < length; i++ {
		if src[i] == ' ' {
			count++
		}
	}

	des := make([]byte, length+count*2)

	for i, j := 0, 0; i < length; {
		if src[i] == ' ' {
			des[j] = '%'
			j++
			des[j] = '2'
			j++
			des[j] = '0'
		} else {
			des[j] = src[i]
		}
		i++
		j++
	}

	return string(des)
}

func Compress(str string) string {
	var buf bytes.Buffer
	src := []rune(str)

	for i := 0; i < len(src); {
		buf.WriteRune(src[i])

		c, j := 1, 0
		for j = i + 1; j < len(src); {
			if src[i] == src[j] {
				c++
			} else {
				break
			}
			j++
		}

		buf.WriteString(strconv.Itoa(c))
		i = j
	}

	res := buf.String()

	if len(res) > len(str) {
		return str
	}
	return res
}

type MatrixNN [][]int32

func (m MatrixNN) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\n%v x %v Matrix\n", len(m), len(m)))

	for i := 0; i < len(m); i++ {
		buf.WriteString(" ")
		for j := 0; j < len(m); j++ {
			buf.WriteString(fmt.Sprint("%v ", m[i][j]))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (m MatrixNN) Compare(mat MatrixNN) bool {
	if len(m) != len(mat) {
		return false
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] != mat[i][j] {
				return false
			}
		}
	}
	return true
}

func (m MatrixNN) Rotate() {
	l := len(m) / 2
	n := len(m) - 1

	for i := 0; i < l; i++ {
		for j := i; j < n-i; j++ {
			top := m[i][j]
			// left -> top
			m[i][j] = m[n-j][i]
			// bottom -> left
			m[n-j][i] = m[n-i][n-j]
			// right -> bottom
			m[n-i][n-j] = m[j][n-i]
			// top -> right
			m[j][n-i] = top
		}
	}
}

type MatrixMN [][]int32

func (a MatrixMN) Zero() {
	m := len(a)
	n := len(a[0])

	bm := make([]bool, m)
	bn := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 0 {
				bm[i] = true
				bn[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if bm[i] || bn[j] {
				a[i][j] = 0
			}
		}
	}
}

func (a MatrixMN) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\n%v x %v Matrix\n", len(a), len(a[0])))

	for i := 0; i < len(a); i++ {
		buf.WriteString(" ")
		for j := 0; j < len(a[0]); j++ {
			buf.WriteString(fmt.Sprint("%v ", a[i][j]))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (a MatrixMN) Compare(mat MatrixMN) bool {
	if len(a) != len(mat) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != mat[i][j] {
				return false
			}
		}
	}
	return true
}

func IsRotation(str1, str2 string) bool {
	len1 := len(str1)
	if len1 == len(str2) && len1 > 0 {
		s1s1 := str1 + str1
		if strings.Contains(s1s1, str2) {
			return true
		}
	}
	return false
}
