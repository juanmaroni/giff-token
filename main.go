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

	mode, _ := token.GetModeFromString(*modeStr)

	config := token.NewTokenConfig(uint16(*length), mode, *characters, *includeChars, *excludeChars)
	fmt.Println(generator.GenerateToken(config))
}
