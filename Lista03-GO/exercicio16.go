package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	arrivalTimes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arrivalTimes[i])
	}

	sort.Ints(arrivalTimes)

	// Verifica se a aula é cancelada
	if arrivalTimes[K-1] > 0 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		for i := 0; i < K; i++ {
			if arrivalTimes[i] <= 0 {
				fmt.Println(i + 1)
			}
		}
	}
}
