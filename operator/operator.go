package operator

func ToNegative(i int) int {
	return ^i + 1
}