package functions

import (
	"fmt"
	"math/rand"
	"os"
)

func FillArrayRandomly(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
}

func MakeRandomArray(n int) []int {
	arr := make([]int, n)
	sorted := MakeSortedArray(n)
	index := 0
	for i := 0; i < n; i++ {
		index = rand.Intn(len(sorted))
		arr[i] = sorted[index]
		sorted = append(sorted[:index], sorted[index+1:]...)
	}
	return arr
}

func MakeSortedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func MakeReversedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	return arr
}

func SaveResults(results []float64, cpu, iterations, size int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(file, `
                                       GO
--------------------------------------------------------------------------------
| Number of logical CPUs used = %d
| Number of Iterations        = %d
| Size of the array           = %d
--------------------------------------------------------------------------------
| Times taken for each iteration (in seconds):
`, cpu, iterations, size)
	if err != nil {
		panic(err)
	}

	for i := 0; i < iterations; i++ {
		str := "  "
		if (i+1)/100 < 1 {
			str += " "
		}
		if (i+1)/10 < 1 {
			str += " "
		}
		_, err = fmt.Fprintf(file, "\n%d.%s%.6f", i+1, str, results[i])
		if err != nil {
			panic(err)
		}
	}

	min, max, sum := 0.0, 0.0, 0.0
	for i := 0; i < iterations; i++ {
		if min > results[i] {
			min = results[i]
		}
		if max < results[i] {
			max = results[i]
		}
		if i == 0 {
			min, max = results[i], results[i]
		}
		sum += results[i]
	}

	_, err = fmt.Fprintf(file, `

--------------------------------------------------------------------------------
| Minimum time: %.6f seconds
| Average time: %.6f seconds
| Maximum time: %.6f seconds
| Total time: %.6f seconds
--------------------------------------------------------------------------------`,
		min, sum/float64(iterations), max, sum)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
