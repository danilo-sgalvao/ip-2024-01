package main

import (
    "fmt"
    "math/big"
)

func calcularFatorial(n int64) *big.Int {
    if n == 0 {
        return big.NewInt(1)
    }

    fatorial := big.NewInt(n)
    for i := n - 1; i > 0; i-- {
        fatorial.Mul(fatorial, big.NewInt(i))
    }
    return fatorial
}

func main() {
    var n int64
    fmt.Print("Digite um número inteiro: ")
    fmt.Scan(&n)

    fatorial := calcularFatorial(n)
    fmt.Printf("%d! = %v\n", n, fatorial)
}
