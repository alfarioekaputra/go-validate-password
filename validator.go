package validate

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const (
	defaultMinLength = 14
	defaultMaxLength = 128
	minAllowedLength = 8
)

// Options configures password validation rules.
type Options struct {
	// MinLength is the minimum required password length in Unicode characters.
	// Defaults to 14 if zero or not set. Cannot be set below 8 (clamped).
	MinLength int

	// MaxLength is the maximum allowed password length in Unicode characters.
	// Defaults to 128 if zero or not set.
	MaxLength int
}

// ValidatePassword checks if the password meets strict security criteria:
// - Minimum 14 characters (Unicode runes)
// - Maximum 128 characters
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
	if minLength < minAllowedLength {
		minLength = minAllowedLength
	}

	maxLength := opts.MaxLength
	if maxLength <= 0 {
		maxLength = defaultMaxLength
	}

	length := utf8.RuneCountInString(password)

	if length < minLength {
		return false, fmt.Sprintf("Password must be at least %d characters long", minLength)
	}
	if length > maxLength {
		return false, fmt.Sprintf("Password must be at most %d characters long", maxLength)
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
		case unicode.IsDigit(char):
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
