package main

import (
	"fmt"
)

func calcularRendaTotal(ingressos int, popular, geral, arquibancada, cadeiras float64) float64 {
	// Calcula o número de ingressos vendidos para cada categoria
	ingressosPopular := float64(ingressos) * (popular / 100.0)
	ingressosGeral := float64(ingressos) * (geral / 100.0)
	ingressosArquibancada := float64(ingressos) * (arquibancada / 100.0)
	ingressosCadeiras := float64(ingressos) * (cadeiras / 100.0)

	// Calcula a renda total do jogo
	rendaTotal := (ingressosPopular * 1.0) + (ingressosGeral * 5.0) + (ingressosArquibancada * 10.0) + (ingressosCadeiras * 20.0)
	return rendaTotal
}

func main() {
	var numCasos int
	fmt.Scanln(&numCasos)

	for i := 1; i <= numCasos; i++ {
		var ingressos int
		var popular, geral, arquibancada, cadeiras float64
		fmt.Scanln(&ingressos, &popular, &geral, &arquibancada, &cadeiras)

		renda := calcularRendaTotal(ingressos, popular, geral, arquibancada, cadeiras)
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}