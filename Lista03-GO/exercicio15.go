package main

import (
	"fmt"
	"math"
)

func minDistance(arr []int) (int, int) {
	minDist := math.MaxInt32
	comparisons := 0

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			dist := int(math.Abs(float64(arr[i] - arr[j])))
			comparisons++
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist, comparisons
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		arr := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&arr[i])
		}

		minDist, comparisons := minDistance(arr)
		fmt.Printf("%d %d\n", minDist, comparisons)
	}
}
