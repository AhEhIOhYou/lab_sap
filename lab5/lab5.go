package lab5

import (
	"math"
	"sort"
)

// FindMedian - параллельное вычисление медианного значения числовой последовательности.
func FindMedian(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}

	ch := make(chan int)

	go func() {
		sort.Ints(arr)
		ch <- arr[n/2]
	}()

	return <-ch
}

// CalculateNorm - параллельное вычисление нормы квадратной матрицы.
func CalculateNorm(matrix [][]float64) float64 {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	ch := make(chan float64)

	for i := 0; i < n; i++ {
		go func(row []float64) {
			norm := calculateRowNorm(row)
			ch <- norm
		}(matrix[i])
	}

	norms := make([]float64, n)
	for i := 0; i < n; i++ {
		norms[i] = <-ch
	}

	sum := 0.0
	for _, value := range norms {
		sum += value * value
	}

	return math.Sqrt(sum)
}

func calculateRowNorm(row []float64) float64 {
	norm := 0.0
	for _, value := range row {
		norm += value * value
	}
	return math.Sqrt(norm)
}

// FindMinMax - параллельное вычисление максимального и минимального элементов одномерного массива
func FindMinMax(arr []int) (int, int) {
	n := len(arr)
	if n == 0 {
		return 0, 0
	}
	if n == 1 {
		return arr[0], arr[0]
	}

	chMin := make(chan int)
	chMax := make(chan int)

	go func() {
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		chMin <- min
	}()

	go func() {
		max := arr[0]
		for _, v := range arr {
			if v > max {
				max = v
			}
		}
		chMax <- max
	}()

	return <-chMin, <-chMax
}
