package generator

import (
	"math"
	"math/rand"
)

const (
	defaultAlphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
	defaultTokenLength = 32
)

type TokenGenerator struct {
	alphabet string
}

func NewTokenGenerator(alphabet string) *TokenGenerator {
	tokenAlphabet := defaultAlphabet
	if alphabet != "" {
		tokenAlphabet = alphabet
	}
	return &TokenGenerator{tokenAlphabet}
}

func (g *TokenGenerator) Generate(length int) string {
	if length <= 0 {
		length = defaultTokenLength
	}
	token := ""
	alphabetLength := len(g.alphabet)
	for i := 0; i < length; i++ {
		token += string(g.alphabet[getRandomNumberFromRange(0, alphabetLength)])
	}
	return token
}

func getRandomNumberFromRange(min int, max int) int {
	rangeSize := max - min

	// not so random, just for sure
	if rangeSize < 0 {
		return min
	}

	log := math.Log2(float64(rangeSize))
	// length in bytes
	bytes := int(log/8) + 1
	// length in bits
	bits := int(log) + 1
	// set all lower bits to 1
	filter := (1 << uint(bits)) - 1

	for {
		randomBytes := make([]byte, bytes)
		rand.Read(randomBytes)
		rnd := int(randomBytes[0])
		// discard irrelevant bits
		rnd = rnd & filter
		if rnd < rangeSize {
			return min + rnd
		}
	}
}
