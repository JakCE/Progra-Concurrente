package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarDatos(n int) ([]float64, []float64) {
	X := make([]float64, n)
	Y := make([]float64, n)

	for i := 0; i < n; i++ {
		X[i] = rand.Float64() * 100
		Y[i] = 2*X[i] + 5 + rand.Float64()*10 // Ruido añadido
	}

	return X, Y
}

func calcularRegresionLineal(X, Y []float64) (float64, float64) {
	var sumX, sumY, sumXY, sumX2 float64
	n := len(X)

	for i := 0; i < n; i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumX2 += X[i] * X[i]
	}

	a := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	b := (sumY - a*sumX) / float64(n)

	return a, b
}

func main() {
	n := 2000000
	X, Y := generarDatos(n)

	start := time.Now()
	a, b := calcularRegresionLineal(X, Y)
	elapsed := time.Since(start)

	fmt.Printf("Resultado Secuencial:\n")
	fmt.Printf("Pendiente (a): %.6f\n", a)
	fmt.Printf("Intersección (b): %.6f\n", b)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}