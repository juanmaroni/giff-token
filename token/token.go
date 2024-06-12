package token

import (
	"errors"
	"giff-token/characters"
)

type Mode string

const (
	Alphanumeric Mode        = "alphanumeric"
	AlphanumericLower Mode   = "alphanumeric lowercase"
	AlphanumericUpper Mode   = "alphanumeric uppercase"
	Allchars Mode            = "allchars"
	AllcharsLower Mode       = "allchars lowercase"
	AllcharsUpper Mode       = "allchars uppercase"
	Alphabetic Mode          = "alphabetic"
	AlphabeticLower Mode     = "alphabetic lowercase"
	AlphabeticUpper Mode     = "alphabetic uppercase"
	Custom Mode              = "custom"
	Digits Mode              = "digits"
)

// map of string to Mode for quick lookup
var modeMap = map[string]Mode{
	"alphanumeric":            Alphanumeric,
	"alphanumeric lowercase":  AlphanumericLower,
	"alphanumeric uppercase":  AlphanumericUpper,
	"allchars":                Allchars,
	"allchars lowercase":      AllcharsLower,
	"allchars uppercase":      AllcharsUpper,
	"alphabetic":              Alphabetic,
	"alphabetic lowercase":    AlphabeticLower,
	"alphabetic uppercase":    AlphabeticUpper,
	"custom":                  Custom,
	"digits":                  Digits,
}

// Function to get Mode from string
func GetModeFromString(s string) (Mode, error) {
	mode, exists := modeMap[s]

	if !exists {
		return "", errors.New("Invalid mode")
	}

	return mode, nil
}

type TokenConfig struct {
	Length uint16
	mode Mode
	customChars string
	includeChars string
	excludeChars string
	Characters []rune
}

// Token options with default values // Delete or change, use NewTokenConfig
func DefaultTokenConfig() TokenConfig {
	const defaultMode = Alphanumeric
	chars, err := GetCharacters(defaultMode, "")

	if err != nil {
		panic(err)
	}

	return TokenConfig {
		Length: 24,
		mode: defaultMode,
		Characters: chars,
	}
}

func NewTokenConfig(length uint16, mode Mode, customChars string, includeChars string, excludeChars string) TokenConfig {
	if (mode != Custom) {
		customChars = ""
	}

	chars, err := GetCharacters(mode, customChars)

	if err != nil {
		panic(err)
	}

	return TokenConfig {
		Length: length,
		mode: mode,
		customChars: customChars,
		includeChars: includeChars,
		excludeChars: excludeChars,
		Characters: chars,
	}
}

// Allowed chars based on the mode selected
func GetCharacters(mode Mode, customCharacters string) ([]rune, error) { // Change parameters to TokenConfig? Add/include and remove/exclude
	switch mode {
	case Alphanumeric:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.ALPHABETIC_LOWERCASE),
			characters.NewCharset(characters.DIGITS),
		)
		return charset.ExtractCharset(), nil
	case AlphanumericLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.DIGITS))
		return charset.ExtractCharset(), nil
	case AlphanumericUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.DIGITS))
		return charset.ExtractCharset(), nil
	case Allchars:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.ALPHABETIC_LOWERCASE),
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		return charset.ExtractCharset(), nil
	case AllcharsLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		return charset.ExtractCharset(), nil
	case AllcharsUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		return charset.ExtractCharset(), nil	
	case Alphabetic:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.ALPHABETIC_LOWERCASE))
		return charset.ExtractCharset(), nil
	case AlphabeticLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		return charset.ExtractCharset(), nil
	case AlphabeticUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		return charset.ExtractCharset(), nil
	case Custom:
		charset := characters.NewCharset(customCharacters)
		return charset.ExtractCharset(), nil
	case Digits:
		charset := characters.NewCharset(characters.DIGITS)
		return charset.ExtractCharset(), nil
	default:
		return nil, errors.New("Something went wrong with the mode selected!")
	}
}
