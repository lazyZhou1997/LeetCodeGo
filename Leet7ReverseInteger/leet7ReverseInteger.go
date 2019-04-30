package main

import (
	"fmt"
)

func main() {
	var input = -123
	result := reverse(input)

	fmt.Println(result)
}

func reverse(x int) int {

	if x == 0 {
		return 0
	}

	var result = 0

	for ; x != 0; x /= 10 {
		result = result*10 + x%10
	}

	if result >= 1<<31 || result < -1<<31 {
		return 0
	}

	return result
}
