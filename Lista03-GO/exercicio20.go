package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

func main() {
	var N int
	fmt.Scanln(&N)

	var points []Point

	// Leitura dos pontos
	for i := 0; i < N; i++ {
		var x, y, z float64
		fmt.Scanln(&x, &y, &z)
		points = append(points, Point{x, y, z})
	}

	// Cálculo dos vetores e busca do maior valor absoluto
	for i := 0; i < N-1; i++ {
		vector := Point{
			points[i+1].x - points[i].x,
			points[i+1].y - points[i].y,
			points[i+1].z - points[i].z,
		}
		mod := math.Sqrt(vector.x*vector.x + vector.y*vector.y + vector.z*vector.z)
		fmt.Printf("%.2f\n", mod)
	}
}
