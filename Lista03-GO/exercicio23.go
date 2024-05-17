package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func calcularFrequenciaVogais(texto string) [5]int {
	vogais := "aeiou"
	frequencias := [5]int{}

	for _, char := range texto {
		lowerChar := strings.ToLower(string(char))
		if strings.Contains(vogais, lowerChar) {
			switch lowerChar {
			case "a":
				frequencias[0]++
			case "e":
				frequencias[1]++
			case "i":
				frequencias[2]++
			case "o":
				frequencias[3]++
			case "u":
				frequencias[4]++
			}
		}
	}

	return frequencias
}

func calcularDistancia(fa, fb [5]int) float64 {
	var soma float64
	for i := 0; i < 5; i++ {
		soma += math.Pow(float64(fa[i]-fb[i]), 2)
	}
	return math.Sqrt(soma)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	textos := strings.Split(texto, ";")
	if len(textos) != 2 {
		fmt.Println("FORMATO INVALIDO!")
		return
	}

	a := textos[0]
	b := textos[1]

	fa := calcularFrequenciaVogais(a)
	fb := calcularFrequenciaVogais(b)

	fmt.Printf("(%d,%d,%d,%d,%d)\n", fa[0], fa[1], fa[2], fa[3], fa[4])
	fmt.Printf("(%d,%d,%d,%d,%d)\n", fb[0], fb[1], fb[2], fb[3], fb[4])

	distancia := calcularDistancia(fa, fb)
	fmt.Printf("%.2f\n", distancia)
}
