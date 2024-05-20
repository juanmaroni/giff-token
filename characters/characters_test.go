package characters

import "testing"

const ALPHABETIC_UPPERCASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABETIC_LOWERCASE string = "abcdefghijklmnopqrstuvwxyz"
const DIGITS string = "0123456789"
const SYMBOLS string = "!@#$%^&*_-+=[]{}/\\?<>,;:\"'"

func TestGetAlphabeticUppercase(t *testing.T) {
    chars := GetAlphabeticUppercase()
	result := string(chars[:])

	if result != ALPHABETIC_UPPERCASE {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, ALPHABETIC_UPPERCASE)
	}
}

func TestGetAlphabeticLowercase(t *testing.T) {
    chars := GetAlphabeticLowercase()
	result := string(chars[:])

	if result != ALPHABETIC_LOWERCASE {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, ALPHABETIC_LOWERCASE)
	}
}

func TestGetDigits(t *testing.T) {
    chars := GetDigits()
	result := string(chars[:])

	if result != DIGITS {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, DIGITS)
	}
}

func TestGetSymbols(t *testing.T) {
    chars := GetSymbols()
	result := string(chars[:])

	if result != SYMBOLS {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, SYMBOLS)
	}
}
