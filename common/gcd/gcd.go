package gcd

// 求两个数的最大公约数 - 欧几里得算法
func Euclid(a, b int) int {
	if b == 0 {
		return a
	}

	return Euclid(b, a%b)
}

// Stein算法
func Stein(a, b int) int {
	switch {
	// 0和n，最大公约数为n
	case a == 0:
		return b
	case b == 0:
		return a
	case a%2 == 0 && b%2 == 0:
		return 2 * Stein(a>>1, b>>1)
	case a%2 == 0:
		return Stein(a>>1, b)
	case b%2 == 0:
		return Stein(a, b>>1)
	default:
		if a < b {
			a, b = b, a
		}
		return Stein(a-b, b)
	}
}

// 求n个数的最大公约数
func NGCD(numbers ...int) int {
	l := len(numbers)
	if l == 1 {
		return numbers[0]
	}
	return Stein(numbers[0], NGCD(numbers[1:]...))
}
