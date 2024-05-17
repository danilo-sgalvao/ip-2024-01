package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	V := make([]int, N)
	W := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
		W[N-1-i] = V[i]
	}

	maxV := V[0]
	for _, num := range V {
		if num > maxV {
			maxV = num
		}
	}

	minW := W[0]
	for _, num := range W {
		if num < minW {
			minW = num
		}
	}

	fmt.Println("V:", V)
	fmt.Println("W:", W)
	fmt.Println("Maior elemento de V:", maxV)
	fmt.Println("Menor elemento de W:", minW)
	fmt.Println()
}
