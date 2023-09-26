package lab1

import (
	"math"
)

/*
3. Для заданного отрезка [a, b] построить массив значений функции f(x).
Реализовать алгоритм в виде отдельной функции.
*/

// f(x) = 2x * sin(x) + 13

func CalcFx(a, b float64) []float64 {
	if a == b {
		return nil
	}

	var res []float64

	if a < b {
		for i := a; i <= b; i += 1.0 {
			fx := 2*i*math.Sin(i) + 13
			res = append(res, math.Round(fx))
		}
	} else {
		for i := a; i >= b; i -= 1.0 {
			fx := 2*i*math.Sin(i) + 13
			res = append(res, math.Round(fx))
		}
	}

	return res
}
