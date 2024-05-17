package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i]%2 == nums[j]%2 {
			if nums[i]%2 == 0 {
				return nums[i] < nums[j]
			} else {
				return nums[i] > nums[j]
			}
		} else {
			return nums[i]%2 == 0
		}
	})

	for _, num := range nums {
		fmt.Println(num)
	}
}
