package main

func isValidPassword(password string) bool {
	// ?
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasUpperCase := false
	hasDigit := false
	for _, c := range password {
		if c >= 48 && c <= 57 {
			hasDigit = true
		} else if c >= 65 && c <= 90 {
			hasUpperCase = true
		}
	}
	return hasDigit && hasUpperCase
}
