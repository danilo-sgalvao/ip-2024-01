package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Println("Digite o número de elementos da sequência:")
    fmt.Scanln(&n)

    sequence := make([]int, n)
    fmt.Println("Digite a sequência de números inteiros:")
    for i := 0; i < n; i++ {
        fmt.Scan(&sequence[i])
    }

    maxCount := 1
    currentCount := 1

    for i := 1; i < n; i++ {
        if sequence[i] == sequence[i-1] {
            currentCount++
        } else {
            if currentCount > maxCount {
                maxCount = currentCount
            }
            currentCount = 1
        }
    }

    if currentCount > maxCount {
        maxCount = currentCount
    }

    fmt.Println("A maior subsequência de elementos iguais possui", maxCount, "elementos.")
}
