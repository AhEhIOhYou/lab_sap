package lab2_test

import (
	"lab2"
	"reflect"
	"testing"
)

func TestCreateArrData(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want map[string]int
	}{
		{
			name: "valid dir and files",
			dir:  "/Users/ahehiohyou/Desktop/1 course/sap/lab2/",
			want: map[string]int{
				"test-1.txt": 5,
				"test-2.txt": 12,
				"test-3.txt": 0,
			},
		},
		{
			name: "invalid dir",
			dir:  "/Users/aheh",
			want: map[string]int{},
		},
		{
			name: "valid dir and no txt files",
			dir:  "/Users/ahehiohyou/Desktop/1 course/sap/lab1/",
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lab2.CreateArrData(tt.dir)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcFx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
