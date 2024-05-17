package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	var N int
	fmt.Scan(&N)

	points := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&points[i].x, &points[i].y, &points[i].z)
	}

	for i := 1; i < N; i++ {
		dist := distance(points[i-1], points[i])
		fmt.Printf("%.2f\n", dist)
	}
}
