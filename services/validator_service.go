package services

import "unicode"

func ValidatePassword(password string) string {
	var errorMsg string
	if len(errorMsg) < 8 {
		errorMsg += "The password contains less than 8 characters!\n"
	}

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

		if unicode.IsSymbol(r) {
			hasSpecial = true
		}
	}

	if !hasUpper {
		errorMsg += "The password doesnt contain an uppercase character!\n"
	}

	if !hasLower {
		errorMsg += "The password doesnt contain a lowercase character!\n"
	}

	if !hasNumber {
		errorMsg += "The password doesnt contain a number!\n"
	}

	if !hasSpecial {
		errorMsg += "The password doesnt contain a special character!\n"
	}

	return errorMsg
}
