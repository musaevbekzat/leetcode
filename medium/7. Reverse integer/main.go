package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		pop := x % 10
		x /= 10

		if int64(result)*10 > math.MaxInt32 {
			return 0
		}
		if int64(result)*10 < math.MinInt32 {
			return 0
		}

		result = result*10 + pop
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
}
