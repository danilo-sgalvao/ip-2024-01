package main

import (
	"fmt"
)

func lerConjunto(tamanho int) map[int]bool {
	conjunto := make(map[int]bool)

	for i := 0; i < tamanho; i++ {
		var elemento int
		fmt.Scan(&elemento)

		if _, existe := conjunto[elemento]; !existe {
			conjunto[elemento] = true
		} else {
			i-- // Permite que o usuário insira novamente um elemento inválido
		}
	}

	return conjunto
}

func imprimirConjunto(conjunto map[int]bool) {
	fmt.Print("(")
	i := 0
	for elemento := range conjunto {
		fmt.Print(elemento)
		i++
		if i < len(conjunto) {
			fmt.Print(",")
		}
	}
	fmt.Println(")")
}

func main() {
	var TA, TB int

	for {
		fmt.Scan(&TA)
		if TA >= 1 && TA <= 100 {
			break
		}
		fmt.Println("Tamanho de A inválido. Insira novamente.")
	}

	for {
		fmt.Scan(&TB)
		if TB >= 1 && TB <= 100 {
			break
		}
		fmt.Println("Tamanho de B inválido. Insira novamente.")
	}

	A := lerConjunto(TA)
	B := lerConjunto(TB)

	// União dos conjuntos
	uniao := make(map[int]bool)
	for elemento := range A {
		uniao[elemento] = true
	}
	for elemento := range B {
		uniao[elemento] = true
	}

	// Interseção dos conjuntos
	intersecao := make(map[int]bool)
	for elemento := range A {
		if _, existe := B[elemento]; existe {
			intersecao[elemento] = true
		}
	}

	// Imprime a união
	imprimirConjunto(uniao)

	// Imprime a interseção
	imprimirConjunto(intersecao)
}
