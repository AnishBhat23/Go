package main

func sum(nums ...int) int {
	// ?

	retSum := 0
	for idx := 0; idx < len(nums); idx++ {
		retSum += nums[idx]
	}

	return retSum
}
