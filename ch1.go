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
