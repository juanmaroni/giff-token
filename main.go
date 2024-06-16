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
