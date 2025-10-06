package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(word string) bool {
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-1-i] {
			return false
		}
	}
	return true
}

func allWordsPalindrome(sentence string) bool {
	words := strings.Fields(sentence)
	for _, word := range words {
		var builder strings.Builder
		for _, r := range word {
			if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
				builder.WriteString(strings.ToLower(string(r)))
			}
		}
		cleanWord := builder.String()
		if cleanWord == "" {
			continue
		}
		if !isPalindrome(cleanWord) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter a sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	result := allWordsPalindrome(sentence)
	fmt.Println(result)
}