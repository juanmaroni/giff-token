package generator

import (
	"giff-token/token"
	"math/rand"
	"sync"
	"time"
)

func GenerateToken(config token.TokenConfig) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	
	charset := config.Characters
	charsetLen := len(charset)

	token := make([]rune, config.Length)

	var wg sync.WaitGroup

	for i := range token {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			token[idx] = charset[rng.Intn(charsetLen)]
		}(i)
	}

	wg.Wait()
	
	return string(token);
}
