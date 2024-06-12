package main

import (
	"flag"
	"fmt"
	"giff-token/generator"
	"giff-token/token"
)

const HELP_TEST = `
༼ つ ◕_◕ ༽つ Giff Token (giff-token)

`

const MIN_TOKEN_LENGTH = 10
const MAX_TOKEN_LENGTH = 4096

func main() {
	flag.Usage = func() {
		fmt.Print(HELP_TEST)
	}

	length := flag.Uint("length", 24, "Token length")
	modeStr := flag.String("mode", "alphanumeric", "Mode to select base characters")
	characters := flag.String("characters", "", "Use custom set of characters")
	includeChars := flag.String("includeChars", "", "Custom charset to include")
	excludeChars := flag.String("excludeChars", "", "Custom charset to exclude")

	flag.Parse()

	if (*length < MIN_TOKEN_LENGTH) {
		*length = MIN_TOKEN_LENGTH
	} else if (*length > MAX_TOKEN_LENGTH) {
		*length = MAX_TOKEN_LENGTH
	}

	mode, _ := token.GetModeFromString(*modeStr) // TODO: Handle error

	if (*characters != "") {
		// Custom mode ignores includeChars and excludeChars
		mode = token.Custom
	}

	config := token.NewTokenConfig(uint16(*length), mode, *characters, *includeChars, *excludeChars)
	fmt.Println(generator.GenerateToken(config))
}
