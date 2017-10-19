package fib

import (
	"testing"
	// "github.com/raypereda/fibonacci/fib"
)

func TestIsNonNegative(t *testing.T) {
	cases :=  []struct {
		in int
	    want bool
	} {
		{-2, false},
		{-1, false},
		{ 0, true},
		{ 1, true},
		{ 2, true},
		{ 3, true},
	}
	for _, c := range cases {
		got := isNonNegative(c.in)
		if got != c.want {
			t.Errorf("isNonNegative(%v) == %v, want %v\n", c.in, got, c.want)
		}
	}
}