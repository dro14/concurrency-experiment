package functions

import (
	"sort"
	"testing"
)

type testPair struct {
	input  []int
	output []int
}

var test = []testPair{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{8, 1, 5, 0, 3, 7, 2, 9, 6, 4}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{7}, []int{7}},
	{[]int{}, []int{}},
}

func TestMergeSortMulti(t *testing.T) {
	for _, pair := range test {
		n := len(pair.input)
		array := make([]int, n)
		buffer := make([]int, n)
		copy(array, pair.input)
		MergeSortMulti(array, buffer)
		for i, v := range array {
			if v != pair.output[i] {
				t.Error(
					"[MergeSortMulti] ERROR:",
					"\nfor:     ", pair.input,
					"\nexpected:", pair.output,
					"\ngot:     ", array,
				)
			}
		}
	}
}

func TestMergeSortSingle(t *testing.T) {
	for _, pair := range test {
		n := len(pair.input)
		array := make([]int, n)
		buffer := make([]int, n)
		copy(array, pair.input)
		MergeSortSingle(array, buffer, 0, n)
		for i, v := range array {
			if v != pair.output[i] {
				t.Error(
					"[MergeSortSingle] ERROR:",
					"\nfor:     ", pair.input,
					"\nexpected:", pair.output,
					"\ngot:     ", array,
				)
			}
		}
	}
}

func TestSort(t *testing.T) {
	for _, pair := range test {
		array := make([]int, len(pair.input))
		copy(array, pair.input)
		sort.Ints(array)
		for i, v := range array {
			if v != pair.output[i] {
				t.Error(
					"[sort.Ints] ERROR:",
					"\nfor:     ", pair.input,
					"\nexpected:", pair.output,
					"\ngot:     ", array,
				)
			}
		}
	}
}
