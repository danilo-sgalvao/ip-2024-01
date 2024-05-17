package main

import "fmt"

func intercalarOrdenado(v1, v2 []int) []int {
	vr := make([]int, len(v1)+len(v2))
	i, j, k := 0, 0, 0

	for i < len(v1) && j < len(v2) {
		if v1[i] < v2[j] {
			vr[k] = v1[i]
			i++
		} else {
			vr[k] = v2[j]
			j++
		}
		k++
	}

	for i < len(v1) {
		vr[k] = v1[i]
		i++
		k++
	}

	for j < len(v2) {
		vr[k] = v2[j]
		j++
		k++
	}

	return vr
}

func main() {
	var q1, q2 int
	fmt.Scan(&q1, &q2)

	v1 := make([]int, q1)
	v2 := make([]int, q2)

	for i := 0; i < q1; i++ {
		fmt.Scan(&v1[i])
	}

	for i := 0; i < q2; i++ {
		fmt.Scan(&v2[i])
	}

	vr := intercalarOrdenado(v1, v2)

	for _, num := range vr {
		fmt.Println(num)
	}
}
