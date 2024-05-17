package main

import (
	"fmt"
)

func countUniqueElements(arr []int) int {
	// Crie um mapa para contar a frequência de cada elemento
	freqMap := make(map[int]int)

	// Percorra o vetor e conte a frequência de cada elemento
	for _, num := range arr {
		freqMap[num]++
	}

	// Inicialize uma variável para contar o número de elementos únicos
	uniqueCount := 0

	// Percorra o mapa e conte quantos elementos têm frequência 1
	for _, freq := range freqMap {
		if freq == 1 {
			uniqueCount++
		}
	}

	return uniqueCount
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	uniqueCount := countUniqueElements(arr)
	fmt.Println(uniqueCount)
}
