package main

func getMessageCosts(messages []string) []float64 {
	// ?
	messagesLen := len(messages)
	messageCosts := make([]float64, messagesLen)

	for msgIdx := 0; msgIdx < messagesLen; msgIdx++ {
		messageCosts[msgIdx] = (float64)(len(messages[msgIdx])) * 0.01
	}

	return messageCosts
}
