package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	data := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	sort.Ints(data)

	median := calculateMedian(data)
	fmt.Printf("%.2f\n", median)
}

func calculateMedian(data []int) float64 {
	N := len(data)
	if N%2 == 0 {
		// Média dos dois elementos centrais
		middle1 := data[N/2]
		middle2 := data[N/2-1]
		return float64(middle1+middle2) / 2.0
	} else {
		// Elemento central
		return float64(data[N/2])
	}
}
