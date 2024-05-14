package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func generateRandomPoints() []Point {
	rand.Seed(time.Now().UnixNano())
	points := make([]Point, 3)
	for i := 0; i < 3; i++ {
		points[i] = Point{X: rand.Float64() * 10, Y: rand.Float64() * 10}
	}
	return points
}

func transformCoordinates(points []Point) []Point {
	transformedPoints := make([]Point, 3)
	for i, point := range points {
		transformedPoints[i] = Point{X: 2*point.X + 10, Y: -3*point.Y - 5}
	}
	return transformedPoints
}

func main() {
	points := generateRandomPoints()
	fmt.Println("Случайные точки до преобразования:")
	for _, point := range points {
		fmt.Printf("(%.2f, %.2f)\n", point.X, point.Y)
	}

	transformedPoints := transformCoordinates(points)
	fmt.Println("\nТочки после преобразования:")
	for _, point := range transformedPoints {
		fmt.Printf("(%.2f, %.2f)\n", point.X, point.Y)
	}
}
