package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)
	soma := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
		soma += vetor[i]
	}

	fmt.Println(soma)
}
