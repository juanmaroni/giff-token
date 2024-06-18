package main

import (
	"errors"
	"flag"
	"fmt"
	"giff-token/generator"
	"giff-token/token"
	"os"
)

const HELP_TEXT = `
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
		fmt.Print(HELP_TEXT)
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
	var err error
	
	if flag.NFlag() == 0 {
		code, err := setConfigFromFile(CONFIG_FILENAME, &config)

		if err != nil {
			if code == 1 {
				config = token.NewTokenConfig(24, token.Alphanumeric, "", "", "")
			} else {
				fmt.Println(err)
				return
			}
		}
	} else {
		if *configFile != "" {
			code, err := setConfigFromFile(*configFile, &config)

			if err != nil {
				if code == 1 {
					config = token.NewTokenConfig(24, token.Alphanumeric, "", "", "")
				} else {
					fmt.Println(err)
					return
				}
			}
		} else {
			if *length < MIN_TOKEN_LENGTH {
				*length = MIN_TOKEN_LENGTH
			} else if *length > MAX_TOKEN_LENGTH {
				*length = MAX_TOKEN_LENGTH
			}
		
			mode, err = token.GetModeFromString(*modeStr)

			if err != nil {
				fmt.Println(err)
				return
			}
		
			if *characters != "" {
				// Custom mode ignores includeChars and excludeChars
				mode = token.Custom
			}

			config = token.NewTokenConfig(uint16(*length), mode, *characters, *includeChars, *excludeChars)
		}
	}
	
	fmt.Println(generator.GenerateToken(config))
}

func setConfigFromFile(filepath string, config *token.TokenConfig) (uint8, error) {
	_, err := os.Stat(filepath);

	if err != nil {
		return 1, errors.New("ERROR: Couldn't open config file.")
	} else {
		parser, err := token.ParseConfigFile(filepath)

		if err != nil {
			return 2, err
		} else {
			*config = *parser

			return 0, nil
		}
	}
}
