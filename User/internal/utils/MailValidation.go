package utils

import (
	"net/mail"
)

func MailValidation(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}
