package main

import "fmt"

func main() {
    // Ler a população de ambos os países
    var popA, popB int
    fmt.Scan(&popA)
    fmt.Scan(&popB)

    // Definir as taxas de crescimento anual
    taxaA := 0.03
    taxaB := 0.015

    // Inicializar o contador de anos
    anos := 0

    // Loop para calcular quantos anos são necessários até que a população de A ultrapasse ou iguale a população de B
    for popA < popB {
        popA += int(float64(popA) * taxaA)
        popB += int(float64(popB) * taxaB)
        anos++
    }

    // Imprimir o número de anos necessários
    fmt.Printf("ANOS = %d\n", anos)
}
