package functions

import (
	"runtime"
)

func MergeSortMulti(arr, buf []int, a, c int) {
	n := c - a
	g := runtime.NumCPU()
	m := n / g

	done := make(chan bool)
	l := a
	r := a + m
	var i int
	for i = 1; i < g; i++ {
		go MergeSortSingle(arr, buf, l, r, done)
		l = r
		r += m
	}
	r = a + c
	go MergeSortSingle(arr, buf, l, r, done)

	for i = 0; i < g; i++ {
		<-done
	}

	var count int
	for i = m; i < n; i += i {
		count = 0
		l = a
		m = l + i
		r = m + i
		for r < c {
			count++
			go mergeMulti(arr, buf, l, m, r, done)
			l = r
			m = l + i
			r = m + i
		}
		if m < c {
			count++
			go mergeMulti(arr, buf, l, m, c, done)
		}
		for j := 0; j < count; j++ {
			<-done
		}
	}
}

func MergeSortSingle(arr, buf []int, a, c int, done chan<- bool) {
	i := a + 1
	for i < c {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
		i += 2
	}
	n := c - a
	var l, m, r int
	for i = 2; i < n; i += i {
		l = a
		m = l + i
		r = m + i
		for r < c {
			mergeSingle(arr, buf, l, m, r)
			l = r
			m = l + i
			r = m + i
		}
		if m < c {
			mergeSingle(arr, buf, l, m, c)
		}
	}
	done <- true
}

func mergeSingle(arr, buf []int, a, b, c int) {
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

func mergeMulti(arr, buf []int, a, b, c int, done chan<- bool) {
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
	done <- true
}
