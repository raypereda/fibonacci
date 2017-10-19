package fib

func isNonNegative(n int) bool {
	if 0 <= n {
		return true
	}
	return false
}

// Fib1 returns of the fibonacci of input
func Fib1(n int) int {
	if !isNonNegative(n) {
		panic("fibonacci input must be non-negative")
	}
	if n < 2 {
		return n
	}
	return Fib1(n-2) + Fib1(n-1)
}

