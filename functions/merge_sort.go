package functions

import "runtime"

func MergeSortMulti(arr, buf []int, start, size int) {
	cpu := runtime.NumCPU()
	divisions := 0
	for i := cpu; i >= 2; i /= 2 {
		divisions++
	}
	done := make(chan bool)
	go mergeSortMulti(arr, buf, start, size, divisions, done)
	<-done
}

func mergeSortMulti(arr, buf []int, a, c, d int, done chan<- bool) {
	b := (a + c) / 2
	if d > 0 {
		d--
		wait := make(chan bool)
		go mergeSortMulti(arr, buf, a, b, d, wait)
		go mergeSortMulti(arr, buf, b, c, d, wait)
		<-wait
		<-wait
	} else {
		MergeSortSingle(arr, buf, a, b)
		MergeSortSingle(arr, buf, b, c)
	}
	merge(arr, buf, a, b, c)
	done <- true
}

func MergeSortSingle(arr, buf []int, a, c int) {
	if c-a == 2 {
		if arr[a] > arr[a+1] {
			arr[a], arr[a+1] = arr[a+1], arr[a]
		}
	} else {
		b := (a + c) / 2
		if b-a > 1 {
			MergeSortSingle(arr, buf, a, b)
		}
		if c-b > 1 {
			MergeSortSingle(arr, buf, b, c)
		}
		merge(arr, buf, a, b, c)
	}
}

func merge(arr, buf []int, a, b, c int) {
	copy(buf[a:c], arr[a:c])
	i, j := a, b
	for i < b && j < c {
		if buf[i] < buf[j] {
			arr[a] = buf[i]
			i++
		} else {
			arr[a] = buf[j]
			j++
		}
		a++
	}
	if i < b {
		copy(arr[a:c], buf[i:b])
	} else {
		copy(arr[a:c], buf[j:c])
	}
}
