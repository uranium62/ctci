package ctci

func IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}

	var check uint = 0

	for _, c := range str {
		val := uint(c) - uint('a')
		if (check & (1 << val)) > 0 {
			return false
		}
		check |= (1 << val)
	}
	return true
}
