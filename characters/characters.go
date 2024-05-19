package characters

func GetAlphabeticUppercase() [26]rune {
	return [26]rune {
        'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
        'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
        'U', 'V', 'W', 'X', 'Y', 'Z',
    }
}

func GetAlphabeticLowercase() [26]rune {
	return [26]rune {
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
        'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
        'u', 'v', 'w', 'x', 'y', 'z',
    }
}

func GetNumeric() [10]rune {
	return [10]rune {
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
}
