package main

import (
	"fmt"
)

func verificarAcertos(sena, aposta []int) (int, int, int) {
	acertos := 0
	for _, num := range aposta {
		for _, sorteado := range sena {
			if num == sorteado {
				acertos++
				break
			}
		}
	}

	switch acertos {
	case 6:
		return 1, 0, 0
	case 5:
		return 0, 1, 0
	case 4:
		return 0, 0, 1
	default:
		return 0, 0, 0
	}
}

func main() {
	var sorteadas [6]int
	for i := range sorteadas {
		fmt.Scan(&sorteadas[i])
	}

	var n int
	fmt.Scan(&n)

	var sena, quina, quadra int

	for i := 0; i < n; i++ {
		var aposta [6]int
		for j := range aposta {
			fmt.Scan(&aposta[j])
		}

		s, q, r := verificarAcertos(sorteadas[:], aposta[:])
		sena += s
		quina += q
		quadra += r
	}

	if sena > 0 {
		fmt.Printf("Houve %d acertador(es) da sena\n", sena)
	} else {
		fmt.Println("Nao houve acertador para sena")
	}

	if quina > 0 {
		fmt.Printf("Houve %d acertador(es) da quina\n", quina)
	} else {
		fmt.Println("Nao houve acertador para quina")
	}

	if quadra > 0 {
		fmt.Printf("Houve %d acertador(es) da quadra\n", quadra)
	} else {
		fmt.Println("Nao houve acertador para quadra")
	}
}
