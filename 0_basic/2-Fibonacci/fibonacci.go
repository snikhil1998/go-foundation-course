package main

import (
	"fmt"
)

func main() {
	var fib[47]int
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < 47; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(fib[i])
	}
}
