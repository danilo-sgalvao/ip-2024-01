package main

import (
	"fmt"
	"sort"
)

func findCombination(numbers []int) []int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	diff := sum - 100

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == diff {
				combination := append(numbers[:i], numbers[i+1:j]...)
				combination = append(combination, numbers[j+1:]...)
				return combination
			}
		}
	}
	return nil
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		numbers := make([]int, 9)
		for i := 0; i < 9; i++ {
			fmt.Scan(&numbers[i])
		}

		sort.Ints(numbers)

		combination := findCombination(numbers)

		for _, num := range combination {
			fmt.Println(num)
		}
	}
}
