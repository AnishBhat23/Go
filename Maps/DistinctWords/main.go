package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	// ?
	if len(messages) == 0 {
		return 0
	}
	distinctWords := make(map[string]struct{})
	for i := 0; i < len(messages); i++ {
		uniq_msg := strings.Fields(messages[i])
		for j := 0; j < len(uniq_msg); j++ {
			if _, ok := distinctWords[strings.ToLower(uniq_msg[j])]; !ok {
				distinctWords[strings.ToLower(uniq_msg[j])] = struct{}{}
			}
		}
	}
	return len(distinctWords)
}
