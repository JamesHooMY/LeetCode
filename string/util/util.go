package util

func IsAlphanumeric[T byte | rune](c T) bool {
	// fix: (c >= 0 && c <= 9) --> (c >= '0' && c <= '9')
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}

	return false
}

func ToLowerCase[T byte | rune](c T) T {
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	return c
}

func Max[T int | int64 | float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T int | int64 | float64](a, b T) T {
	if a < b {
		return a
	}

	return b
}