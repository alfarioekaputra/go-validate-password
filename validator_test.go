package validate

import "testing"

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
