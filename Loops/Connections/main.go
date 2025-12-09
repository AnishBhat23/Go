package main

func countConnections(groupSize int) int {
	// ?
	numConnections := 0
	for i := 0; i < groupSize; i++ {
		for j := i + 1; j < groupSize; j++ {
			numConnections++
		}
	}
	return numConnections
}
