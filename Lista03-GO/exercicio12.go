package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanln(&N)

	numbers := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&numbers[i])
	}

	sort.Ints(numbers)

	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println()
}
