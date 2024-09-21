# Giff Token
༼ つ ◕_◕ ༽つ Token/password/whatever

## Usage
```
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

```

## Status
Usable, but unfinished.
