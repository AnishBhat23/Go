package main

func createMatrix(rows, cols int) [][]int {
	// ?
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = i * j
		}
	}

	return grid
}
