package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	for i := 0; i < len(messages); i++ {
		//sms_tags := []string{}
		sms_tags := tagger(messages[i])
		messages[i].tags = sms_tags
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
