package main

import "fmt"

func twoSum(nums []int, target int) []int {
	valueToIndex := make(map[int]int)

	for i, num := range nums {
		if j, ok := valueToIndex[target-num]; ok {
			return []int{j, i}
		}
		valueToIndex[num] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
