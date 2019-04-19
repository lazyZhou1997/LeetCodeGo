package main

import "fmt"

func main() {
	var input = []int{2, 7, 11, 15}
	fmt.Println(twoSum(input, 9))

}

func twoSum(num []int, target int) []int {
	var result []int
	temp := make(map[int]int)

	for indices1, result1 := range num {
		indices2, ok := temp[target-result1]

		if ok {
			result = append(result, indices2)
			result = append(result, indices1)
		} else {
			temp[result1] = indices1
		}
	}

	return result
}
