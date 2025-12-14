package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	if len(names) == 0 {
		return nil
	}
	nameCounts := make(map[rune]map[string]int)
	for i := 0; i < len(names); i++ {
		runes := []rune(names[i])
		if _, ok := nameCounts[runes[0]]; !ok {
			nameCounts[runes[0]] = make(map[string]int)
		}
		nameCounts[runes[0]][names[i]]++
	}
	return nameCounts
}
