package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanln(&str)
	for _, ch := range str {
		fmt.Print(string(ch), string(ch))
	}
}
