package main

import (
	"fmt"
)

func countingSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	count := make([]int, max+1)

	for _, num := range arr {
		count[num]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	return sorted
}

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		sorted := countingSort(arr)

		for i, num := range sorted {
			fmt.Print(num)
			if i < n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
