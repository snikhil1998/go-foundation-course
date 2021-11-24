package main

import (
	"fmt"
)

func getFactors(n int) (int, int) {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return i, n/i;
		}
	}
	return 1, n
}

func main() {
	for p := 2; p < 26; p++ {
		f1, f2 := getFactors(p)
		if f1 != 1 {
			fmt.Println(p, "is a product of", f1, "*", f2)
		} else {
			fmt.Println(p, "is prime")
		}
	}
}
