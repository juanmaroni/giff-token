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

type TokenOptions struct {
	length uint8
	mode Mode
	Characters []rune
}

// Token options with default values
func DefaultTokenOptions() TokenOptions {
	const defaultMode = Alphanumeric
	chars, err := GetCharacters(defaultMode, "")

	if err != nil {
		panic(err)
	}

	return TokenOptions {
		length: 24,
		mode: defaultMode,
		Characters: chars,
	}
}

func NewTokenOptions(length uint8, mode Mode, customChars string, includeChars string, excludeChars string) TokenOptions {
	if (mode != Custom) {
		customChars = ""
	}

	chars, err := GetCharacters(mode, customChars)

	if err != nil {
		panic(err)
	}

	return TokenOptions {
		length: length,
		mode: mode,
		Characters: chars,
	}
}

// Allowed chars based on the mode selected
func GetCharacters(mode Mode, customCharacters string) ([]rune, error) {
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
