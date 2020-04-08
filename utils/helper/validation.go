package helper

import "regexp"

// IsEmail check valid Email
func IsEmail(email string) bool {
	valid, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return valid
}

// IsPassword check valid password  8 - 32
func IsPassword(password string) bool {
	valid, _ := regexp.MatchString("^.{6,20}$", password)
	return valid
}
