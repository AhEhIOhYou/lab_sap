package lab1_test

import (
	"lab1"
	"reflect"
	"testing"
)

func TestCalcFx(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want []float64
	}{
		{
			name: "valid data 1",
			a:    3,
			b:    12,
			want: []float64{14, 7, 3, 10, 22, 29, 20, 2, -9, 0},
		},
		{
			name: "valid data 2",
			a:    3,
			b:    1,
			want: []float64{14, 17, 15},
		},
		{
			name: "not valid data",
			a:    1,
			b:    1,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lab1.CalcFx(tt.a, tt.b)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcFx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
