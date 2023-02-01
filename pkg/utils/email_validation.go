package utils

import (
	"strings"
)

func SimpleEmailValidation(email string) bool {
	parsedMail := strings.Split(email, "@")
	if len(parsedMail) == 1 {
		return false
	}

	domainMail := strings.Split(parsedMail[1], ".")
	return len(domainMail) != 1
}
