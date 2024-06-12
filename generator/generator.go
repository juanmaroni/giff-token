package generator

import (
	"giff-token/token"
	"math/rand"
	"time"
)

func GenerateToken(config token.TokenConfig) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	
	charset := config.Characters
	charsetLen := len(charset)

	token := make([]rune, config.Length)

	for i := range token {
		token[i] = charset[rng.Intn(charsetLen)]
	}
	
	return string(token);
}
