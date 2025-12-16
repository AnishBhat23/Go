package main

import (
	"strings"
)

func removeProfanity(message *string) {
	// ?
	/*
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
	 */
	replacements := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}
	for old, new := range replacements {
		*message = strings.ReplaceAll(*message, old, new)
	}
}
