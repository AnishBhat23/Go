package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	switch plan {
	case planPro:
		retString := messages[:]
		return retString, nil
	case planFree:
		retString := messages[:2]
		return retString, nil
	default:
		var err error = errors.New("unsupported plan")
		return nil, err
	}

}
