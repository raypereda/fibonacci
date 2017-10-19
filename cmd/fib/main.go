package main

import (
	"fmt"
	"github.com/raypereda/fibonacci/fib"
)

func main() {
	fmt.Println("Hi")
	for i := 0; i < 10; i++ {
		fmt.Println(i, "fibonacci number is", fib.Fib1(i))		
	}
}