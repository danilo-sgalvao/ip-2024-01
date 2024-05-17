package main

import (
    "fmt"
)

func main() {
    var n int
    for {
        fmt.Scan(&n)
        if n == 0 {
            break
        }

        vetor := make([]int, n)
        var maior, indiceMaior int
        for i := 0; i < n; i++ {
            fmt.Scan(&vetor[i])
            if vetor[i] > maior {
                maior = vetor[i]
                indiceMaior = i
            }
        }

        fmt.Printf("%d %d\n\n", indiceMaior, maior)
    }
}
