package main

import (
	"fmt"
)

func printPrimes(max int) {
	// ?
	for idx1 := 2; idx1 < max+1; idx1++ {
		if idx1 == 2 {
			fmt.Printf("%v\n", idx1)
		} else if idx1%2 == 0 {
			continue
		}
		for idx2 := 3; idx2*idx2 < idx1; idx2 = idx2 + 2 {
			if idx1%idx2 == 0 {
				break
			}
		}
		fmt.Printf("%v\n", idx1)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
