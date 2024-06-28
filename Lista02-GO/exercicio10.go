package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var matricula int
	var horas float64
	var valor float64

	for {

		fmt.Scanf("%d %f %f", &matricula, &horas, &valor)
		bufio.NewReader(os.Stdin).ReadByte()
		
		if matricula == 0 {
			break
		}
		salario := horas * valor

		fmt.Printf("%d %.2f\n", matricula, salario)

	}

}