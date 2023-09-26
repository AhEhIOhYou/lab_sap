package xdata

import (
	"errors"
	"math"
	"sort"
)

/*
3. Создать пакет xdata предоставляющий интерфейс для вычислений на последовательностях (вычисление средних значений, минимальных и максимальных значений, медианных значений, средне геометрических значений и т. п.)
*/

type SequenceCalculator interface {
	Average(sequence []float64) (float64, error)
	Min(sequence []float64) (float64, error)
	Max(sequence []float64) (float64, error)
	Median(sequence []float64) (float64, error)
	GeometricMean(sequence []float64) (float64, error)
}

type XCalc struct{}

// Average вычисляет среднее значение последовательности
func (x *XCalc) Average(sequence []float64) (float64, error) {
	if len(sequence) == 0 {
		return 0, errors.New("sequence is empty")
	}

	sum := 0.0
	for _, value := range sequence {
		sum += value
	}

	return sum / float64(len(sequence)), nil
}

// Min находит минимальное значение в последовательности
func (x *XCalc) Min(sequence []float64) (float64, error) {
	if len(sequence) == 0 {
		return 0, errors.New("sequence is empty")
	}

	min := math.Inf(1)
	for _, value := range sequence {
		if value < min {
			min = value
		}
	}

	return min, nil
}

// Max находит максимальное значение в последовательности
func (x *XCalc) Max(sequence []float64) (float64, error) {
	if len(sequence) == 0 {
		return 0, errors.New("sequence is empty")
	}

	max := math.Inf(-1)
	for _, value := range sequence {
		if value > max {
			max = value
		}
	}

	return max, nil
}

// Median вычисляет медианное значение последовательности
func (x *XCalc) Median(sequence []float64) (float64, error) {
	if len(sequence) == 0 {
		return 0, errors.New("sequence is empty")
	}

	sortedSequence := make([]float64, len(sequence))
	copy(sortedSequence, sequence)
	sort.Float64s(sortedSequence)

	middleIndex := len(sortedSequence) / 2
	if len(sortedSequence)%2 == 1 {
		return sortedSequence[middleIndex], nil
	}

	return (sortedSequence[middleIndex-1] + sortedSequence[middleIndex]) / 2, nil
}

// GeometricMean вычисляет среднее геометрическое значение последовательности
func (x *XCalc) GeometricMean(sequence []float64) (float64, error) {
	if len(sequence) == 0 {
		return 0, errors.New("sequence is empty")
	}

	product := 1.0
	for _, value := range sequence {
		product *= value
	}

	return math.Pow(product, 1/float64(len(sequence))), nil
}
