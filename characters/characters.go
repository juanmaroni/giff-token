package characters

func GetAlphabeticUppercase() []rune {
	var letters []rune

    for ch := 'A'; ch <= 'Z'; ch++ {
        letters = append(letters, ch)
    }

	return letters
}

func GetAlphabeticLowercase() []rune {
	var letters []rune

    for ch := 'a'; ch <= 'z'; ch++ {
        letters = append(letters, ch)
    }

	return letters
}

func GetNumeric() []rune {
	var nums []rune

    for ch := '0'; ch <= '9'; ch++ {
        nums = append(nums, ch)
    }

	return nums
}
