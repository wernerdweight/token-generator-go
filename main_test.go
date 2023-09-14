package generator

import "testing"

func TestNewTokenGenerator(t *testing.T) {
	g := NewTokenGenerator("")
	if g == nil {
		t.Error("NewTokenGenerator() should return not nil")
	}
}

func TestTokenGenerator_Generate(t *testing.T) {
	g := NewTokenGenerator("")
	token := g.Generate(32)
	if len(token) != 32 {
		t.Error("TokenGenerator.Generate() should return token with length 32")
	}
}

func TestTokenGenerator_GenerateWithCustomLength(t *testing.T) {
	g := NewTokenGenerator("")
	token := g.Generate(64)
	if len(token) != 64 {
		t.Error("TokenGenerator.Generate() should return token with length 64")
	}
}

func TestTokenGenerator_GenerateWithCustomAlphabet(t *testing.T) {
	g := NewTokenGenerator("abc")
	token := g.Generate(32)
	for _, c := range token {
		if c != 'a' && c != 'b' && c != 'c' {
			t.Error("TokenGenerator.Generate() should return token with alphabet 'abc'")
		}
	}
}

func TestTokenGenerator_GenerateWithCustomAlphabetAndLength(t *testing.T) {
	g := NewTokenGenerator("abc")
	token := g.Generate(64)
	for _, c := range token {
		if c != 'a' && c != 'b' && c != 'c' {
			t.Error("TokenGenerator.Generate() should return token with alphabet 'abc'")
		}
	}
}
