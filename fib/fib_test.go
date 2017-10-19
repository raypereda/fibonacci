package fib_test

import (
	"testing"
	"github.com/raypereda/fibonacci/fib"
)

func TestFib(t *testing.T) {
	cases :=  []struct {
		in, want int
	} {
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, c := range cases {
		got := fib.Fib1(c.in)
		if got != c.want {
			t.Errorf("Fib1(%v) == %v, want %v\n", c.in, got, c.want)
		}
	}
}
