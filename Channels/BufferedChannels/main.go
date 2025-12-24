package main

func addEmailsToQueue(emails []string) chan string {
	// ?
	emailLen := len(emails)
	emailCh := make(chan string, emailLen)

	for email := 0; email < emailLen; email++ {
		emailCh <- emails[email]
	}

	return emailCh
}
