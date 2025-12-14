package main

import (
	"errors"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	user_elem, ok := users[name]
	if ok {
		if user_elem.scheduledForDeletion {
			delete(users, name)
			return true, err
		} else {
			return false, nil
		}
	} else {
		var err error = errors.New("not found")
		return false, err
	}
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
