package services

import "unicode"

func ValidatePassword(password string) []string {
	errors := make([]string, 0)

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, r := range password {
		if unicode.IsUpper(r) {
			hasUpper = true
		}

		if unicode.IsLower(r) {
			hasLower = true
		}

		if unicode.IsNumber(r) {
			hasNumber = true
		}

		if unicode.IsPunct(r) {
			hasSpecial = true
		}
	}

	if len(password) < 8 {
		errors = append(errors, "the password contains less than 8 characters")
	}

	if !hasUpper {
		errors = append(errors, "the password doesnt contain an uppercase character")
	}

	if !hasLower {
		errors = append(errors, "the password doesnt contain a lowercase character")
	}

	if !hasNumber {
		errors = append(errors, "the password doesnt contain a number")
	}

	if !hasSpecial {
		errors = append(errors, "the password doesnt contain a special character")
	}

	return errors
}
