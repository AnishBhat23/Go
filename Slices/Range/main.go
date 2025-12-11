package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for i, msgVar := range msg {
		for _, badVar := range badWords {
			if msgVar == badVar {
				return i
			}
		}
	}
	return -1
}
