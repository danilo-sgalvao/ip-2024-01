package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, d int
	for {
		fmt.Scan(&n, &d)
		if n == 0 && d == 0 {
			break
		}

		var num string
		fmt.Scan(&num)

		result := getMaxPrize(num, d)
		fmt.Println(result)
	}
}

func getMaxPrize(num string, d int) string {
	stack := make([]rune, 0)

	for _, digit := range num {
		for d > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			d--
		}
		stack = append(stack, digit)
	}

	// Remove d digits if still remaining
	stack = stack[:len(stack)-d]

	// Convert stack back to string
	result := strings.TrimLeft(string(stack), "0")

	// If result is empty, return "0"
	if result == "" {
		return "0"
	}

	return result
}
