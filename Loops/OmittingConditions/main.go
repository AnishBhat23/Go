package main

func maxMessages(thresh int) int {
	// ?
	basicCost := 100
	for msg, cost := 0, 0; ; msg++ {
		cost += (basicCost + msg)
		if cost > thresh {
			return msg
		}
	}

}
