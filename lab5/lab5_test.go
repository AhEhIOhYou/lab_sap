package lab5_test

import (
	"lab5"
	"reflect"
	"testing"
)

func TestParallelMedian(t *testing.T) {
	tests := []struct {
		name string
		nums []float64
		n    int
		want []float64
	}{
		{"Empty slice", []float64{}, 1, []float64{}},
		{"One element", []float64{1.5}, 1, []float64{1.5}},
		{"Even length", []float64{1.2, 3.4, 5.6, 7.8}, 2, []float64{2.3, 6.7}},
		{"Odd length", []float64{1.2, 3.4, 5.6, 7.8, 9.0}, 3, []float64{1.2, 5.6, 9.0}},
		{"Unsorted slice", []float64{2.4, 6.8, 4.5, 8.9, 10.1}, 4, []float64{2.4, 4.5, 8.9, 10.1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lab5.ParallelMedian(tt.nums, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelMedian(%v, %d) = %v, want %v", tt.nums, tt.n, got, tt.want)
			}
		})
	}
}
