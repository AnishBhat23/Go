package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	retString := [3]string{primary, secondary, tertiary}

	retCost := [3]int{len(primary), (len(primary) + len(secondary)), (len(primary) + len(secondary) + len(tertiary))}

	return retString, retCost
}
