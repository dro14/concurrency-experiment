package main

import (
	"concurrency_experiment/functions"
	"runtime"
	"sort"
	"time"
)

const (
	size       = 100000000
	iterations = 100
)

func main() {
	array := make([]int, size)
	buffer := make([]int, size)
	results := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		functions.FillArrayRandomly(array)
		start := time.Now()
		functions.MergeSortMulti(array, buffer, 0, size)
		results[i] = time.Since(start).Seconds()
	}
	functions.SaveResults(results, runtime.NumCPU(), iterations, size, "multi.txt")

	for i := 0; i < iterations; i++ {
		functions.FillArrayRandomly(array)
		start := time.Now()
		functions.MergeSortSingle(array, buffer, 0, size)
		results[i] = time.Since(start).Seconds()
	}
	functions.SaveResults(results, runtime.NumGoroutine(), iterations, size, "single.txt")

	for i := 0; i < iterations; i++ {
		functions.FillArrayRandomly(array)
		start := time.Now()
		sort.Ints(array)
		results[i] = time.Since(start).Seconds()
	}
	functions.SaveResults(results, runtime.NumGoroutine(), iterations, size, "builtin.txt")
}
