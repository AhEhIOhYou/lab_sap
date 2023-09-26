package lab5

import (
	"sort"
	"sync"
)

// median которая вычисляет медиану среза чисел
func median(nums []float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Float64s(nums)
	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	} else {
		return nums[n/2]
	}
}

// ParallelMedian вычисляет параллельно медиану каждого подсреза среза nums и возвращает их в срезе medians
func ParallelMedian(nums []float64, n int) []float64 {
	var wg sync.WaitGroup
	ch := make(chan float64, n)
	m := len(nums) / n
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * m
			end := start + m
			if i == n-1 {
				end = len(nums)
			}
			sub := nums[start:end]
			ch <- median(sub)
		}(i)
	}
	wg.Wait()
	close(ch)
	var medians []float64
	for m := range ch {
		medians = append(medians, m)
	}

	return medians
}
