package util

func Min[T int | int64 | float64](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T int | int64 | float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Abs[T int | int64 | float64](a T) T {
	if a < 0 {
		return -a
	}

	return a
}
