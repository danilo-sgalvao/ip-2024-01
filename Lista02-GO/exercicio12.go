package main

import "fmt"

func main() {
	var valorIngresso float64
	var valorInicial float64
	var valorFinal float64

	fmt.Scanln(&valorIngresso)
	fmt.Scanln(&valorInicial)
	fmt.Scanln(&valorFinal)

	lucro := 0.0
	lucroatual := 0.0
	melhorValor :=0.0
	ingressos:=0

	for i := valorInicial; i < valorIngresso; i++ {
		if valorInicial >= valorFinal {
			break
		}
		diferencaMenor := valorIngresso - i
		ingressosMenor := diferencaMenor/0.5*25 + 120
		lucroatual = ingressosMenor*i - 200 - 0.05*ingressosMenor

		fmt.Printf("V: %.2f, N: %.0f, L: %.2f\n", i, ingressosMenor, lucroatual)
		if lucroatual > lucro {
			lucro = lucroatual
			melhorValor = i
			ingressos = int(ingressosMenor) 
		}
	}

	if lucroatual > lucro {
		lucro = lucroatual
	}

	for i := valorIngresso; i <= valorFinal; i++ {
		if valorInicial >= valorFinal {
			break
		}
		diferencaMaior := i-valorIngresso
		ingressosMaior := 120 - diferencaMaior/0.5*30
		lucroatual = ingressosMaior*i - 200 - 0.05*ingressosMaior

		fmt.Printf("V: %.2f, N: %.0f, L: %.2f\n", i, ingressosMaior, lucroatual)

		if lucroatual > lucro {
			lucro = lucroatual
			melhorValor = i
			ingressos = int(ingressosMaior) 
		}

	}
	if lucroatual > lucro {
		lucro = lucroatual
	}
	if valorInicial >= valorFinal {
		fmt.Println("invalido")
	}

	fmt.Printf("Melhor valor final: %.2f\nLucro: %.2f\nNumero de ingressos:%d",melhorValor,lucro,ingressos)

}