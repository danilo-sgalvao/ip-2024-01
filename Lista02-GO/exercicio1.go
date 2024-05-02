package main

import (
	"fmt"
)

func main() {
	// Definindo constantes
	const (
		numProvas      = 8
		numListas      = 5
		presencaMinima = 75 // Em porcentagem
		horasMinimas   = 128
	)

	// Loop para ler os dados dos alunos
	for {
		// Ler os dados do aluno
		var matricula int
		var notasProvas [numProvas]float64
		var notasListas [numListas]float64
		var notaTrabalho, presenca float64
		_, err := fmt.Scan(&matricula)
		if err != nil || matricula == -1 {
			break // Finaliza o loop quando todos os dados forem lidos ou quando for inserido -1
		}
		for i := 0; i < numProvas; i++ {
			fmt.Scan(&notasProvas[i])
		}
		for i := 0; i < numListas; i++ {
			fmt.Scan(&notasListas[i])
		}
		fmt.Scan(&notaTrabalho, &presenca)

		// Calcular a média das notas das provas
		mediaProvas := calculaMedia(notasProvas[:])

		// Calcular a média das notas das listas
		mediaListas := calculaMedia(notasListas[:])

		// Calcular a nota final
		notaFinal := 0.7*mediaProvas + 0.15*mediaListas + 0.15*notaTrabalho

		// Verificar a situação do aluno
		situacao := ""
		if presenca >= presencaMinima && notaFinal >= 6 {
			situacao = "APROVADO"
		} else if presenca < presencaMinima && notaFinal < 6 {
			situacao = "REPROVADO POR NOTA E POR FREQUENCIA"
		} else if presenca < presencaMinima {
			situacao = "REPROVADO POR FREQUENCIA"
		} else if notaFinal < 6 {
			situacao = "REPROVADO POR NOTA"
		}

		// Imprimir resultado
		fmt.Printf("Matricula: %d, Nota Final: %.2f, Situação Final: %s\n", matricula, notaFinal, situacao)
	}
}

// Função para calcular a média de um conjunto de notas
func calculaMedia(notas []float64) float64 {
	total := 0.0
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}
