package utils

import "regexp"

func NormalizeContactNumber(contact, country string) string {
	if country == "BD" && contact[0] == '0' && len(contact) == 11 {
		return "+88" + contact
	}
	return contact
}

func ContactNumberValidation(contact, country string) (string, bool) {
	var bdMobileRegex = regexp.MustCompile(`^(\+8801|01[3-9][0-9]{8})$`)
	if !bdMobileRegex.MatchString(contact) {
		return "", false
	}
	contact = NormalizeContactNumber(contact, country)
	return contact, true
}
