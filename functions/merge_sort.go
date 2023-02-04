package functions

import "runtime"

func MergeSortMulti(array, buffer []int, start, size int) {
	cpu := runtime.NumCPU()
	divisions := 0
	for i := cpu; i >= 2; i /= 2 {
		divisions++
	}
	done := make(chan bool)
	go mergeSortMulti(array, buffer, start, size, divisions, done)
	<-done
}

func mergeSortMulti(array, buffer []int, a, c, d int, done chan<- bool) {
	b := (a + c) / 2
	if d > 0 {
		d--
		wait := make(chan bool)
		go mergeSortMulti(array, buffer, a, b, d, wait)
		go mergeSortMulti(array, buffer, b, c, d, wait)
		<-wait
		<-wait
	} else {
		MergeSortSingle(array, buffer, a, b)
		MergeSortSingle(array, buffer, b, c)
	}
	merge(array, buffer, a, b, c)
	done <- true
}

func MergeSortSingle(array, buffer []int, a, c int) {
	if c-a == 2 {
		if array[a] > array[a+1] {
			array[a], array[a+1] = array[a+1], array[a]
		}
	} else {
		b := (a + c) / 2
		if b-a > 1 {
			MergeSortSingle(array, buffer, a, b)
		}
		if c-b > 1 {
			MergeSortSingle(array, buffer, b, c)
		}
		merge(array, buffer, a, b, c)
	}
}

func merge(array, buffer []int, a, b, c int) {
	copy(buffer[a:c], array[a:c])
	i, j := a, b
	for i < b && j < c {
		if buffer[i] < buffer[j] {
			array[a] = buffer[i]
			i++
		} else {
			array[a] = buffer[j]
			j++
		}
		a++
	}
	if i < b {
		copy(array[a:c], buffer[i:b])
	} else {
		copy(array[a:c], buffer[j:c])
	}
}
