package main

func concurrentFib(n int) []int {
	// ?
	fibs := make(chan int)
	res := []int{}
	go fibonacci(n, fibs)
	for i := range fibs {
		res = append(res, i)
	}
	return res
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
