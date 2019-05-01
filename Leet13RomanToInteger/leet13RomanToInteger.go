package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MXIV"))
}

func romanToInt(s string) int {
	var roman2alabo = map[rune]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}

	before := 1000
	result := 0
	value := 0
	for _, ch := range s {
		value = roman2alabo[ch]

		if before < value {
			result += value - before*2
		} else {
			result += value
		}

		before = value
	}

	return result
}
