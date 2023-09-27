package lab5_test

import (
	"lab5"
	"testing"
)

func TestFindMedian(t *testing.T) {
	arr := []int{-1, 0, 22, 34, 50}
	median := lab5.FindMedian(arr)
	if median != 22 {
		t.Errorf("Expected minimum value to be 22 but got %d", median)
	}
}

func TestCalculateNorm(t *testing.T) {
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	norm := lab5.CalculateNorm(matrix)
	expectedNorm := 16.881943016134134
	if norm != expectedNorm {
		t.Errorf("Expected to be %f, but got %f", expectedNorm, norm)
	}
}

func TestFindMinMax(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	min, max := lab5.FindMinMax(arr)
	if min != 1 {
		t.Errorf("Expected minimum value to be 1 but got %d", min)
	}
	if max != 5 {
		t.Errorf("Expected maximum value to be 5 but got %d", max)
	}
}
