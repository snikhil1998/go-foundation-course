package main

import (
	"fmt"
)

func main() {
	var fac[13]int
	fac[0] = 1
	fac[1] = 1
	for i := 2; i < 13; i++ {
		fac[i] = i*fac[i-1]
	}
	var n int
	fmt.Scanln(&n)
	fmt.Println(fac[n])
}
