package main

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {

	var new_user User
	new_user.Name = name
	new_user.Membership.Type = membershipType

	if membershipType == "premium" {
		new_user.Membership.MessageCharLimit = 1000
	} else {
		new_user.Membership.MessageCharLimit = 100
	}

	return new_user
}
