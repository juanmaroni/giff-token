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

func TestAddChar(t *testing.T) {
	expect := "1237"
	chars := NewCharset("123")
	chars.Add("7")
	result := extractAndSortCharacters(chars.ExtractCharset())
	
	if result != expect {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expect)
	}
}

func TestRemoveChar(t *testing.T) {
	expect := "123"
	chars := NewCharset("1239")
	chars.Remove("9")
	result := extractAndSortCharacters(chars.ExtractCharset())

	if result != expect {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expect)
	}
}

func TestMergeCharset(t *testing.T) {
	expect := "123456789"
	charset1 := NewCharset("123")
	charset2 := NewCharset("456")
	charset3 := NewCharset("789")
	charset1.MergeCharset(charset2, charset3)
	result := extractAndSortCharacters(charset1.ExtractCharset())

	if result != expect {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expect)
	}
}

func extractAndSortCharacters(array []rune) string {
	sort.Slice(array, func(i, j int) bool {
        return array[i] < array[j]
    })

	return string(array)
}
