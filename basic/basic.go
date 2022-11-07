package basic

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Int | Float
}

type Basic interface {
	~bool | ~string | Number
}

// Max returns the largest of x or y.
func Max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the smallest of x or y.
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Abs returns the absolute value of x.
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
