package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		var err1 error = errors.New("invalid sizes")
		return nil, err1
	}
	userMap := make(map[string]user)

	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
