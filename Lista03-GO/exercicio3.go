package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }

    for i := n - 1; i >= 0; i-- {
        fmt.Printf("%d ", vetor[i])
    }
    fmt.Println()
}
