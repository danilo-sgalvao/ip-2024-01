package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	frequencies := make([]int, 101)
	maxFrequency := math.MinInt32
	maxValue := -1

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		frequencies[num]++
		if frequencies[num] > maxFrequency || (frequencies[num] == maxFrequency && num < maxValue) {
			maxFrequency = frequencies[num]
			maxValue = num
		}
	}

	fmt.Println(maxValue)
	fmt.Println(maxFrequency)
}
