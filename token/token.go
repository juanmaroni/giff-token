package token

import (
	"errors"
	"giff-token/characters"
)

type Mode string

const (
	Alphanumeric Mode        = "alphanumeric"
	AlphanumericLower Mode   = "alphanumeric_lowercase"
	AlphanumericUpper Mode   = "alphanumeric_uppercase"
	Allchars Mode            = "allchars"
	AllcharsLower Mode       = "allchars_lowercase"
	AllcharsUpper Mode       = "allchars_uppercase"
	Alphabetic Mode          = "alphabetic"
	AlphabeticLower Mode     = "alphabetic_lowercase"
	AlphabeticUpper Mode     = "alphabetic_uppercase"
	Custom Mode              = "custom"
	Digits Mode              = "digits"
)

// map of string to Mode for quick lookup
var modeMap = map[string]Mode {
	"alphanumeric":            Alphanumeric,
	"alphanumeric_lowercase":  AlphanumericLower,
	"alphanumeric_uppercase":  AlphanumericUpper,
	"allchars":                Allchars,
	"allchars_lowercase":      AllcharsLower,
	"allchars_uppercase":      AllcharsUpper,
	"alphabetic":              Alphabetic,
	"alphabetic_lowercase":    AlphabeticLower,
	"alphabetic_uppercase":    AlphabeticUpper,
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
	chars, err := GetCharacters(defaultMode, "", "", "")

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

	chars, err := GetCharacters(mode, customChars, includeChars, excludeChars)

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
func GetCharacters(mode Mode, customCharacters string, includeChars string, excludeChars string) ([]rune, error) {
	switch mode {
	case Alphanumeric:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.ALPHABETIC_LOWERCASE),
			characters.NewCharset(characters.DIGITS),
		)
		charset.Add(includeChars)
		charset.Remove(excludeChars)

		return charset.ExtractCharset(), nil
	case AlphanumericLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.DIGITS))
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case AlphanumericUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.DIGITS))
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case Allchars:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.ALPHABETIC_LOWERCASE),
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case AllcharsLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case AllcharsUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(
			characters.NewCharset(characters.DIGITS),
			characters.NewCharset(characters.SYMBOLS),
		)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil	
	case Alphabetic:
		charset := characters.NewCharset(characters.ALPHABETIC_UPPERCASE)
		charset.MergeCharsets(characters.NewCharset(characters.ALPHABETIC_LOWERCASE))
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case AlphabeticLower:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case AlphabeticUpper:
		charset := characters.NewCharset(characters.ALPHABETIC_LOWERCASE)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	case Custom:
		charset := characters.NewCharset(customCharacters)
		return charset.ExtractCharset(), nil
	case Digits:
		charset := characters.NewCharset(characters.DIGITS)
		charset.Add(includeChars)
		charset.Remove(excludeChars)
		
		return charset.ExtractCharset(), nil
	default:
		return nil, errors.New("Something went wrong with the mode selected!")
	}
}
