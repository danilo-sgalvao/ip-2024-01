package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Digite um número inteiro positivo: ")
    _, err := fmt.Scan(&n)

    if err != nil || n <= 1 {
        fmt.Println("Numero invalido!")
        return
    }

    isPrime := true
    for i := 2; i <= n/2; i++ {
        if n%i == 0 {
            isPrime = false
            break
        }
    }

    if isPrime {
        fmt.Println("PRIMO")
    } else {
        fmt.Println("NAO PRIMO")
    }
}
