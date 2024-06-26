package characters

import "sync"

const ALPHABETIC_UPPERCASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABETIC_LOWERCASE string = "abcdefghijklmnopqrstuvwxyz"
const DIGITS string = "0123456789"
const SYMBOLS string = "!\"#$%&'*+,-/:;<=>?@[\\]^_{}"

type Charset map[rune]struct{}

func NewCharset(characters string) Charset {
	charset := make(Charset)

    for _, c := range characters {
        charset[c] = struct{}{}
    }

    return charset
}

func (charset Charset) Add(characters string) {
	for _, c := range characters {
        charset[c] = struct{}{}
    }
}

func (charset Charset) Remove(characters string) {
	for _, c := range characters {
        delete(charset, c)
    }
}

func (charset Charset) ExtractCharset() []rune {
	chars := make([]rune, len(charset))
	i := 0
	
	for k := range charset {
		chars[i] = k
		i++
	}

	return chars
}

func (charset Charset) MergeCharsets(moreCharsets... Charset) {
	var wg sync.WaitGroup

	for _, anotherCharset := range moreCharsets {
		for k, v := range anotherCharset {
			wg.Add(1)

			go func(key rune, value struct{}) {
				defer wg.Done()
				charset[key] = value
			}(k, v)
		}
	}

	wg.Wait()
}
