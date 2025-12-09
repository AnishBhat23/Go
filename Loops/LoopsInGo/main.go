package main

func bulkSend(numMessages int) float64 {
	// ?
	var totalCost float64 = 0.0
	basicCost := 1.0
	basicFee := 0.01
	for msg := 0; msg < numMessages; msg++ {
		totalCost = totalCost + (basicCost + (float64(msg) * basicFee))
	}
	
	return totalCost
}
