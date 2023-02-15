package functions

import (
	"sort"
	"testing"
)

type testPair struct {
	input  []int
	output []int
}

func constructTests() []testPair {
	var testCases []testPair

	testCase := testPair{[]int{}, []int{}}
	testCases = append(testCases, testCase)

	//for i := 0; i < 100; i++ {
	//	testCase = testPair{[]int{i}, []int{i}}
	//	testCases = append(testCases, testCase)
	//}

	//for i := 2; i < 1025; i++ {
	//	testCase = testPair{MakeRandomArray(i), MakeSortedArray(i)}
	//	testCases = append(testCases, testCase)
	//}

	testCase = testPair{MakeSortedArray(1025), MakeSortedArray(1025)}
	testCases = append(testCases, testCase)
	testCase = testPair{MakeRandomArray(1025), MakeSortedArray(1025)}
	testCases = append(testCases, testCase)
	testCase = testPair{MakeReversedArray(1025), MakeSortedArray(1025)}
	testCases = append(testCases, testCase)

	return testCases
}

func TestMergeSortMulti(t *testing.T) {
	tests := constructTests()
	for _, test := range tests {
		n := len(test.input)
		arr := make([]int, n)
		buf := make([]int, n)
		copy(arr, test.input)
		MergeSortMulti(arr, buf, 0, n)
		for i, v := range arr {
			if v != test.output[i] {
				t.Error(
					"[MergeSortMulti] ERROR:",
					"\nfor:     ", test.input,
					"\nexpected:", test.output,
					"\ngot:     ", arr,
				)
			}
		}
	}
}

func TestMergeSortSingle(t *testing.T) {
	tests := constructTests()
	for _, test := range tests {
		n := len(test.input)
		arr := make([]int, n)
		buf := make([]int, n)
		copy(arr, test.input)
		done := make(chan bool)
		go MergeSortSingle(arr, buf, 0, n, done)
		<-done
		for i, v := range arr {
			if v != test.output[i] {
				t.Error(
					"[MergeSortSingle] ERROR:",
					"\nfor:     ", test.input,
					"\nexpected:", test.output,
					"\ngot:     ", arr,
				)
			}
		}
	}
}

func TestSort(t *testing.T) {
	tests := constructTests()
	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)
		sort.Ints(arr)
		for i, v := range arr {
			if v != test.output[i] {
				t.Error(
					"[sort.Ints] ERROR:",
					"\nfor:     ", test.input,
					"\nexpected:", test.output,
					"\ngot:     ", arr,
				)
			}
		}
	}
}
