package xdata_test

import (
	"testing"
	"xdata"
)

func TestSequenceCalculator(t *testing.T) {
	xCalc := &xdata.XCalc{}

	sequence := []float64{1, 2, 3, 4, 5}

	average, err := xCalc.Average(sequence)
	if err != nil {
		t.Errorf("Average() returned an error: %v", err)
	}
	if average != 3 {
		t.Errorf("Average() = %v; want 3", average)
	}

	min, err := xCalc.Min(sequence)
	if err != nil {
		t.Errorf("Min() returned an error: %v", err)
	}
	if min != 1 {
		t.Errorf("Min() = %v; want 1", min)
	}

	max, err := xCalc.Max(sequence)
	if err != nil {
		t.Errorf("Max() returned an error: %v", err)
	}
	if max != 5 {
		t.Errorf("Max() = %v; want 5", max)
	}

	median, err := xCalc.Median(sequence)
	if err != nil {
		t.Errorf("Median() returned an error: %v", err)
	}
	if median != 3 {
		t.Errorf("Median() = %v; want 3", median)
	}

	geometricMean, err := xCalc.GeometricMean(sequence)
	if err != nil {
		t.Errorf("GeometricMean() returned an error: %v", err)
	}
	if geometricMean != 2.605171084697352 {
		t.Errorf("GeometricMean() = %v; want 2.605171084697352", geometricMean)
	}
}
