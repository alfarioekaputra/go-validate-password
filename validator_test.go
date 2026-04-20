package validate

import "testing"

// --- Hardening: Unicode rune count ---

func TestValidatePassword_MultibyteCharCountedAsOneCharacter(t *testing.T) {
	// "éàüöñ" = 5 chars, each 2 bytes = 10 bytes total
	// padding to 13 runes total — should fail min length of 14
	password := "éàüöñAb1!xyzq" // 13 runes, but >14 bytes
	ok, _ := ValidatePassword(password)
	if ok {
		t.Error("expected invalid: 13 Unicode runes should not satisfy minLength=14")
	}
}

// --- Hardening: Max length ---

func TestValidatePassword_RejectsPasswordExceedingMaxLength(t *testing.T) {
	long := make([]byte, 129)
	for i := range long {
		long[i] = 'A'
	}
	long[0] = 'a'
	long[1] = '1'
	long[2] = '!'
	ok, _ := ValidatePassword(string(long))
	if ok {
		t.Error("expected invalid: password exceeds default max length of 128")
	}
}

func TestValidatePasswordWithOptions_CustomMaxLength(t *testing.T) {
	opts := Options{MinLength: 8, MaxLength: 16}
	ok, _ := ValidatePasswordWithOptions("ValidP@ss1ValidP@ss1", opts) // 20 chars > 16
	if ok {
		t.Error("expected invalid: password exceeds custom MaxLength=16")
	}
}

// --- Hardening: IsDigit not IsNumber ---

func TestValidatePassword_FractionCharDoesNotCountAsDigit(t *testing.T) {
	// ½ is a "number" (unicode.IsNumber) but not a "digit" (unicode.IsDigit)
	// password has upper, lower, special, length ok — but digit is only ½
	password := "ValidPassword½!AA" // 17 chars, has upper/lower/special, digit=only ½
	ok, _ := ValidatePassword(password)
	if ok {
		t.Error("expected invalid: fraction ½ should not count as a digit")
	}
}

// --- Hardening: MinLength floor ---

func TestValidatePasswordWithOptions_MinLengthFloorEnforced(t *testing.T) {
	opts := Options{MinLength: 1} // dangerously low
	// password of 4 chars that meets all other criteria
	ok, _ := ValidatePasswordWithOptions("aA1!", opts)
	if ok {
		t.Errorf("expected invalid: MinLength=1 should be clamped to minimum safe floor")
	}
}

func TestValidatePassword_DefaultMinLength(t *testing.T) {
	ok, msg := ValidatePassword("Short1!")
	if ok {
		t.Errorf("expected invalid, got: %s", msg)
	}
}

func TestValidatePasswordWithOptions_CustomMinLength(t *testing.T) {
	opts := Options{MinLength: 8}
	ok, msg := ValidatePasswordWithOptions("Valid1!x", opts)
	if !ok {
		t.Errorf("expected valid with minLength=8, got: %s", msg)
	}
}

func TestValidatePasswordWithOptions_BelowCustomMinLength(t *testing.T) {
	opts := Options{MinLength: 8}
	ok, msg := ValidatePasswordWithOptions("V1!x", opts)
	if ok {
		t.Errorf("expected invalid, got: %s", msg)
	}
	_ = msg
}

func TestValidatePasswordWithOptions_DefaultsTo14WhenZero(t *testing.T) {
	opts := Options{}
	ok, _ := ValidatePasswordWithOptions("Short1!x", opts)
	if ok {
		t.Error("expected invalid when MinLength=0 defaults to 14")
	}
}
