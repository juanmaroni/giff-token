package main

import (
	"flag"
	"fmt"
)

const HELP_TEST = `
༼ つ ◕_◕ ༽つ Giff Token (giff-token)
`

func main() {
	flag.Usage = func() {
		fmt.Print(HELP_TEST)
	}

	length := flag.Uint("length", 24, "Token length")
	mode := flag.String("mode", "alphanumeric", "Mode to select base characters")
	characters := flag.String("characters", "", "Use custom set of characters")
	includeChars := flag.String("includeChars", "", "Custom charset to include")
	excludeChars := flag.String("excludeChars", "", "Custom charset to exclude")

	flag.Parse()

	fmt.Println("length:", *length)
    fmt.Println("mode:", *mode)
	fmt.Println("characters:", *characters)
    fmt.Println("includeChars:", *includeChars)
	fmt.Println("excludeChars:", *excludeChars)
    fmt.Println("tail:", flag.Args())
}
