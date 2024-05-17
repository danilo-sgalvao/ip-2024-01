package main

import "fmt"

func main() {
    for {
        var n int
        fmt.Scan(&n)
        if n == 0 {
            break
        }

        numbers := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scan(&numbers[i])
        }

        max := findMax(numbers)
        frequency := make([]int, max+1)

        for _, num := range numbers {
            for i := 0; i <= max; i++ {
                if num <= i {
                    frequency[i]++
                }
            }
        }

        for i := 0; i <= max; i++ {
            fmt.Printf("(%d) %d\n", i, frequency[i])
        }
    }
}

func findMax(numbers []int) int {
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}
