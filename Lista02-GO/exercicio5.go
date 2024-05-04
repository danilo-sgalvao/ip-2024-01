package main

import "fmt"

func main() {
    var n int
    fmt.Println("Digite o número de elementos da sequência:")
    fmt.Scanln(&n)

    sequence := make([]int, n)
    fmt.Println("Digite a sequência de números inteiros:")
    for i := 0; i < n; i++ {
        fmt.Scan(&sequence[i])
    }

    maxSegment := 1
    currentSegment := 1

    for i := 1; i < n; i++ {
        if sequence[i] > sequence[i-1] {
            currentSegment++
        } else {
            if currentSegment > maxSegment {
                maxSegment = currentSegment
            }
            currentSegment = 1
        }
    }

    if currentSegment > maxSegment {
        maxSegment = currentSegment
    }

    fmt.Println("O comprimento do segmento crescente máximo é:", maxSegment-1)
}
