package characters

import (
	"sort"
	"testing"
)

func TestGetAlphabeticUppercase(t *testing.T) {
    chars := NewCharset(ALPHABETIC_UPPERCASE)
	result := extractAndSortCharacters(chars.ExtractCharset())

	if result != ALPHABETIC_UPPERCASE {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, ALPHABETIC_UPPERCASE)
	}
}

func TestGetAlphabeticLowercase(t *testing.T) {
    chars := NewCharset(ALPHABETIC_LOWERCASE)
	result := extractAndSortCharacters(chars.ExtractCharset())

	if result !=ALPHABETIC_LOWERCASE {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, ALPHABETIC_LOWERCASE)
	}
}

func TestGetDigits(t *testing.T) {
    chars := NewCharset(DIGITS)
	result := extractAndSortCharacters(chars.ExtractCharset())

	if result != DIGITS {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, DIGITS)
	}
}

func TestGetSymbols(t *testing.T) {
    chars := NewCharset(SYMBOLS)
	result := extractAndSortCharacters(chars.ExtractCharset())

	if result != SYMBOLS {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, SYMBOLS)
	}
}

func extractAndSortCharacters(array []rune) string {
	sort.Slice(array, func(i, j int) bool {
        return array[i] < array[j]
    })

	return string(array)
}
