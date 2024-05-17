package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break // Se houver um erro na entrada, como EOF, terminamos o loop
		}

		binary := decimalToBinary(n)
		fmt.Println(binary)
	}
}

// Função para converter um número decimal para binário
func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		n /= 2
	}
	return binary
}
