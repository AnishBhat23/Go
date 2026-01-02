package main

func getLast[T any](s []T) T {
	// ?
	var last T
	if len(s) == 0 {
		return last
	} else {
		return s[len(s)-1]
	}
}
