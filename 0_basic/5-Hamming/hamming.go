package main

import (
	"fmt"
)

func hammDist(dna1 string, dna2 string) int {
	if len(dna1) != len(dna2) {
		return -1
	}
	res := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			res++
		}
	}
	return res
}

func main() {
	var dna1, dna2 string
	fmt.Scanln(&dna1)
	fmt.Scanln(&dna2)
	fmt.Println(hammDist(dna1, dna2))
}
