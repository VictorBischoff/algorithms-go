package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 3) - 5*math.Pow(x, 2) + 3*x + 7
}

func fPrime(x float64) float64 {
	return 3*math.Pow(x, 2) - 10*x + 3
}

func newtonRaphson(x0 float64, tolerance float64) float64 {
	x := x0
	deltaX := f(x) / fPrime(x)
	for math.Abs(deltaX) >= tolerance {
		x -= deltaX
		deltaX = f(x) / fPrime(x)
	}
	return x
}

func main() {
	root := newtonRaphson(5, 0.0001)
	fmt.Printf("Rod af funktionen: %f\n", root)
}