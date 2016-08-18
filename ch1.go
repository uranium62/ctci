package ctci


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
