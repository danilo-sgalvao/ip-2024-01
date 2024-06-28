package main

import "fmt"

func main() {
	var n int
	var tolerancia float64

	fmt.Scanln(&n)
	fmt.Scanln(&tolerancia)

	inicio := 1.0

	for {
		r := (float64(inicio) + float64(n)/float64(inicio)) / 2
		inicio = r

		erro := float64(n) - float64(inicio)*float64(inicio) 
		if erro < 0 {
			erro = -erro
		}
		fmt.Printf("r: %.9f, erro: %.9f\n", r, erro)
		if erro <= tolerancia {
			break
		}

	}
}
