package main

import (
	"flag"
	"fmt"
	"giff-token/generator"
	"giff-token/token"
	"os"
)

const HELP_TEST = `
༼ つ ◕_◕ ༽つ Giff Token (giff-token)

Usage:
    giff-token
    giff-token -help
    giff-token [-length <length> -mode <mode> -includeChars <characters> -excludeChars <characters>]
    giff-token [-length <length> -characters <characters>]
    giff-token -config <file>

Options:
    -h | -help      Show this screen.
    -length         Minimum is 10, maximum is 4096. Default: 24.
    -mode           Select predefined base characters. Choose from: "alphanumeric", "alphanumeric_lowercase",
                    "alphanumeric_uppercase", "allchars", "allchars_lowercase", "allchars_uppercase",
                    "alphabetic", "alphabetic_lowercase", "alphabetic_uppercase" and "digits".
                    Default: "alphanumeric".
    -characters     Use a custom set of characters. "-mode", "-includeChars" and "-excludeChars" will be ignored.
    -includeChars   Include custom set of characters to a selected mode.
    -excludeChars   Exclude custom set of characters from a selected mode.
    -config         Use a config file to specify the previous parameters.

No options equals to:
    giff-token -length 24 -mode alphanumeric

If there is a valid config file called "config.giff" in the same directory as the binary, no options equals to:
    giff-token -config config.giff

Examples:
    giff-token -help
    giff-token -length 100
    giff-token -mode allchars
    giff-token -length 24 -characters asdfghjklqwerty
    giff-token -length 2000 -mode alphabetic -includeChars 123 -excludeChars abc
    giff-token -config /path/to/config.giff

`

const MIN_TOKEN_LENGTH = 10
const MAX_TOKEN_LENGTH = 4096
const CONFIG_FILENAME = "config.giff"

func main() {
	flag.Usage = func() {
		fmt.Print(HELP_TEST)
	}

	length := flag.Uint("length", 24, "Token length")
	modeStr := flag.String("mode", "alphanumeric", "Mode to select base characters")
	characters := flag.String("characters", "", "Use custom set of characters")
	includeChars := flag.String("includeChars", "", "Custom charset to include")
	excludeChars := flag.String("excludeChars", "", "Custom charset to exclude")
	configFile := flag.String("config", "", "Configuration file")

	flag.Parse()

	var config token.TokenConfig
	var mode token.Mode
	
	if flag.NFlag() == 0 {
		_, err := os.Stat(CONFIG_FILENAME);

		if err == nil {
			parser, err := token.ParseConfigFile(CONFIG_FILENAME)

			if err != nil {
				// TODO: Handle error
			} else {
				config = *parser
			}
		} else {
			config = token.NewTokenConfig(24, token.Alphanumeric, "", "", "")
		}
	} else {
		if *configFile != "" {
			_, err := os.Stat(*configFile);

			if err != nil {
				// TODO: parse file as config
			} else {
				// TODO: Handle error
			}
		} else {
			if *length < MIN_TOKEN_LENGTH {
				*length = MIN_TOKEN_LENGTH
			} else if *length > MAX_TOKEN_LENGTH {
				*length = MAX_TOKEN_LENGTH
			}
		
			mode, _ = token.GetModeFromString(*modeStr) // TODO: Handle error
		
			if *characters != "" {
				// Custom mode ignores includeChars and excludeChars
				mode = token.Custom
			}

			config = token.NewTokenConfig(uint16(*length), mode, *characters, *includeChars, *excludeChars)
		}
	}
	
	fmt.Println(generator.GenerateToken(config))
}
