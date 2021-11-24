package fizzbuzz

import (	
	"strconv"
)

func FizzBuzz(n int) string {
	var result string
	if n%3 == 0 {
		result += "Fizz"
	}
	if n%5 == 0 {
		result += "Buzz"
	}
	if len(result) == 0 {
		result = strconv.Itoa(n)
	}
	return result
}
