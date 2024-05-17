package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N)

	// Verificando se N está dentro do intervalo especificado
	for N < 1 || N > 1000 {
		fmt.Println("O valor de N deve pertencer ao intervalo 1 ≤ N ≤ 1000. Por favor, insira um valor válido para N:")
		fmt.Scan(&N)
	}

	// Criando o vetor V com tamanho N
	V := make([]int, N)

	// Lendo os elementos do vetor V
	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
	}

	// Lendo o inteiro K
	fmt.Scan(&K)

	// Contando quantos elementos de V são maiores ou iguais a K
	count := 0
	for _, num := range V {
		if num >= K {
			count++
		}
	}

	// Imprimindo o resultado
	fmt.Println(count)
}
