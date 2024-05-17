package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var cpf string
		fmt.Scan(&cpf)

		if isValidCPF(cpf) {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}

func isValidCPF(cpf string) bool {
	// Remove pontos e traços do CPF
	cpf = removeNonDigits(cpf)

	// Verifica se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	digit1 := 11 - remainder
	if digit1 >= 10 {
		digit1 = 0
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	digit2 := 11 - remainder
	if digit2 >= 10 {
		digit2 = 0
	}

	// Verifica se os dígitos verificadores são iguais aos do CPF
	return digit1 == int(cpf[9]-'0') && digit2 == int(cpf[10]-'0')
}

func removeNonDigits(s string) string {
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result += string(char)
		}
	}
	return result
}
