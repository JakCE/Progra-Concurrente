package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Parcial struct {
	sumX  float64
	sumY  float64
	sumXY float64
	sumX2 float64
}

func generarDatos(n int) ([]float64, []float64) {
	X := make([]float64, n)
	Y := make([]float64, n)

	for i := 0; i < n; i++ {
		X[i] = rand.Float64() * 100
		Y[i] = 2*X[i] + 5 + rand.Float64()*10
	}

	return X, Y
}

func calcularParcial(X, Y []float64, inicio, fin int, ch chan<- Parcial, wg *sync.WaitGroup) {
	defer wg.Done()

	var parcial Parcial
	for i := inicio; i < fin; i++ {
		parcial.sumX += X[i]
		parcial.sumY += Y[i]
		parcial.sumXY += X[i] * Y[i]
		parcial.sumX2 += X[i] * X[i]
	}

	ch <- parcial
}

func calcularRegresionLinealConcurrente(X, Y []float64, numGoroutines int) (float64, float64) {
	n := len(X)
	tamañoBloque := n / numGoroutines
	ch := make(chan Parcial, numGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		inicio := i * tamañoBloque
		fin := inicio + tamañoBloque
		if i == numGoroutines-1 {
			fin = n
		}

		wg.Add(1)
		go calcularParcial(X, Y, inicio, fin, ch, &wg) // Aquí SÍ usamos `go`
	}

	wg.Wait()
	close(ch)

	var total Parcial
	for parcial := range ch {
		total.sumX += parcial.sumX
		total.sumY += parcial.sumY
		total.sumXY += parcial.sumXY
		total.sumX2 += parcial.sumX2
	}

	a := (float64(n)*total.sumXY - total.sumX*total.sumY) / (float64(n)*total.sumX2 - total.sumX*total.sumX)
	b := (total.sumY - a*total.sumX) / float64(n)

	return a, b
}

func main() {
	n := 2000000
	X, Y := generarDatos(n)

	start := time.Now()
	numGoroutines := 8
	a, b := calcularRegresionLinealConcurrente(X, Y, numGoroutines)
	elapsed := time.Since(start)

	fmt.Printf("Resultado Concurrente:\n")
	fmt.Printf("Pendiente (a): %.6f\n", a)
	fmt.Printf("Intersección (b): %.6f\n", b)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}