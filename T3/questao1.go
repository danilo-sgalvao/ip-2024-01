package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inverter(original string) string {
	invertida := []rune(original) 
	for i, j := 0, len(invertida)-1; i < j; i, j = i+1, j-1 {
		invertida[i], invertida[j] = invertida[j], invertida[i]
	}
	return string(invertida) 
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linha := scanner.Text()
		palavras := strings.Fields(linha) 

		for i, j := 0, len(palavras)-1; i < j; i, j = i+1, j-1 {
			palavras[i], palavras[j] = palavras[j], palavras[i]
		}
		
		for i, palavra := range palavras {
			palavras[i] = inverter(palavra)
		}
		fmt.Println(strings.Join(palavras, " "))
		break
	}

}

