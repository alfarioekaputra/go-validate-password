package validate

import (
	"fmt"
	"unicode"
)

const defaultMinLength = 14

// Options configures password validation rules.
type Options struct {
	// MinLength is the minimum required password length.
	// Defaults to 14 if zero or not set.
	MinLength int
}

// ValidatePassword checks if the password meets strict security criteria:
// - Minimum 14 characters
// - At least one uppercase letter
// - At least one lowercase letter
// - At least one digit
// - At least one special character
func ValidatePassword(password string) (bool, string) {
	return ValidatePasswordWithOptions(password, Options{})
}

// ValidatePasswordWithOptions is like ValidatePassword but with configurable options.
func ValidatePasswordWithOptions(password string, opts Options) (bool, string) {
	minLength := opts.MinLength
	if minLength <= 0 {
		minLength = defaultMinLength
	}

	if len(password) < minLength {
		return false, fmt.Sprintf("Password must be at least %d characters long", minLength)
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return false, "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter"
	}
	if !hasNumber {
		return false, "Password must contain at least one number"
	}
	if !hasSpecial {
		return false, "Password must contain at least one special character"
	}

	return true, "Password is valid"
}
