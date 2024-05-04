package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Print("Digite a quantidade de times: ")
    fmt.Scan(&N)

    if N < 2 {
        fmt.Println("Campeonato invalido!")
        return
    }

    fmt.Println("Sequência de finais:")

    final := 1
    for i := 1; i <= N; i++ {
        for j := i + 1; j <= N; j++ {
            fmt.Printf("Final %d: Time%d X Time%d\n", final, i, j)
            final++
        }
    }
}
