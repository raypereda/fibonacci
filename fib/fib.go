package fib

// Fib1 returns of the fibonacci of input
func Fib1(n int) int {
	if n < 0 {
		panic("fibonacci input most be non-negative")
	}
	if n < 2 {
		return n
	}
	return Fib1(n-2) + Fib1(n-1)
}

