package main

import (
	"fmt"
)

func main() {
	var n, i, K, s float64

	// Leitura dos valores de entrada
	fmt.Println("Digite um número n de 0 a 9:")
	fmt.Scan(&n)
	fmt.Println("Digite o valor inicial i:")
	fmt.Scan(&i)
	fmt.Println("Digite o número de valores K:")
	fmt.Scan(&K)
	fmt.Println("Digite o incremento s:")
	fmt.Scan(&s)

	// Tabuada de soma
	fmt.Println("Tabuada de soma:")
	for j := 0.0; j < K; j++ {
		B := i + j*s
		R := n + B
		fmt.Printf("%.2f + %.2f = %.2f\n", n, B, R)
	}

	// Tabuada de subtração
	fmt.Println("Tabuada de subtração:")
	for j := 0.0; j < K; j++ {
		B := i + j*s
		R := n - B
		fmt.Printf("%.2f - %.2f = %.2f\n", n, B, R)
	}

	// Tabuada de multiplicação
	fmt.Println("Tabuada de multiplicação:")
	for j := 0.0; j < K; j++ {
		B := i + j*s
		R := n * B
		fmt.Printf("%.2f * %.2f = %.2f\n", n, B, R)
	}

	// Tabuada de divisão
	fmt.Println("Tabuada de divisão:")
	for j := 0.0; j < K; j++ {
		B := i + j*s
		if B != 0 {
			R := n / B
			fmt.Printf("%.2f / %.2f = %.2f\n", n, B, R)
		} else {
			fmt.Printf("%.2f / %.2f = Indefinido\n", n, B)
		}
	}
}
