package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)

	// Criando o vetor V com tamanho N
	V := make([]int, N)

	// Lendo os elementos do vetor V
	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
	}

	// Lendo a quantidade de buscas M
	fmt.Scan(&M)

	// Criando um mapa para armazenar os elementos do vetor V
	// Isso torna a busca mais eficiente
	// A chave será o número e o valor será true se o número estiver presente no vetor
	numbersMap := make(map[int]bool)
	for _, num := range V {
		numbersMap[num] = true
	}

	// Realizando as buscas
	for i := 0; i < M; i++ {
		var num int
		fmt.Scan(&num)
		if numbersMap[num] {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NAO ACHEI")
		}
	}
}
