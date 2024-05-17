package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	// Crie um conjunto para armazenar elementos únicos
	uniqueSet := make(map[int]bool)

	// Leia os elementos e adicione-os ao conjunto
	for i := 0; i < N; i++ {
		var num int
		fmt.Scanln(&num)
		uniqueSet[num] = true
	}

	// Imprima os elementos únicos na ordem em que foram fornecidos
	for num := 1; num <= N; num++ {
		if uniqueSet[num] {
			fmt.Println(num)
		}
	}
}
