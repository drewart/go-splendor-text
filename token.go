package main

import (
	"fmt"
	"log"
)

type Gem int
type Gems [6]int

const (
	Emerald Gem = iota
	Sapphire
	Ruby
	Diamond
	Onyx
	Gold
)

var GemList = []Gem{Emerald, Sapphire, Ruby, Diamond, Onyx, Gold}

type Token struct {
	Code  byte
	Gem   Gem
	Color string
	Name  string
}

/*
// used
type TokenCount struct {
	Token Token
	Count int
}*/

func (t *Token) GetShort() string {
	return string(t.Code)
}

func (g Gem) String() string {
	return string(GemCode(g))
}

func (g *Gems) String(showZero bool) string {
	var out string
	delim := ""
	for i, c := range g {
		if c > 0 || showZero {
			gem := Gem(i)
			out = out + delim + fmt.Sprintf("%d%s", c, gem.String())
			delim = ":"
		}
	}
	return out
}

func (g *Gems) Add(gems Gems) Gems {
	var sum Gems
	for _, gem := range g {
		sum[gem] = g[gem] + gems[gem]
	}
	return sum
}

func (g *Gems) Sub(gems Gems) Gems {
	var diff Gems
	for _, gem := range g {
		diff[gem] = g[gem] - gems[gem]
	}
	return diff
}

// GenTokens
/*
func GenTokenCount(s string) []TokenCount {

	var tokens []TokenCount

	for i := 0; i < len(s); i++ {
		numByte := s[i]
		num := int(numByte) - 48

		if num < 0 || num > 9 {
			log.Fatal("number out of range", s)
		}
		i = i + 1
		c := s[i]

		token := GetToken(c)
		tokenCnt := TokenCount{Token: token, Count: num}

		tokens = append(tokens, tokenCnt)
	}
	return tokens
}
*/

func GenGems(s string) Gems {
	var gems Gems

	for i := 0; i < len(s); i++ {
		numByte := s[i]
		num := int(numByte) - 48

		if num < 0 || num > 9 {
			log.Fatal("number out of range", s)
		}
		i = i + 1
		c := s[i]

		token := GetToken(c)
		g := token.Gem
		gems[g] += num
	}
	return gems
}

func GemCode(g Gem) byte {
	var code byte
	switch g {
	case Emerald:
		code = 'e'
	case Gold:
		code = 'j'
	case Sapphire:
		code = 's'
	case Onyx:
		code = 'o'
	case Diamond:
		code = 'd'
	case Ruby:
		code = 'r'
	default:
		log.Fatalf("unknown Token Code %v", g)
	}
	return code
}

func GetGemToken(g Gem) Token {
	var token Token
	switch g {
	case Emerald:
		token.Code = 'e'
		token.Gem = Emerald
		token.Color = "green"
		token.Name = "Emerald"
	case Gold:
		token.Code = 'j'
		token.Gem = Gold
		token.Color = "gold"
		token.Name = "Joker"
	case Sapphire:
		token.Code = 's'
		token.Gem = Sapphire
		token.Color = "blue"
		token.Name = "Sapphire"
	case Onyx:
		token.Code = 'o'
		token.Gem = Onyx
		token.Color = "black"
		token.Name = "Onyx"
	case Diamond:
		token.Code = 'd'
		token.Gem = Diamond
		token.Color = "white"
		token.Name = "Diamond"
	case Ruby:
		token.Code = 'r'
		token.Gem = Ruby
		token.Color = "red"
		token.Name = "Ruby"
	default:
		log.Fatalf("unknown Token Code %v", g)
	}
	return token
}

func GetToken(b byte) Token {
	var token Token
	token.Code = b
	switch b {
	case 'e':
		token.Gem = Emerald
		token.Color = "green"
		token.Name = "Emerald"
	case 'j':
		token.Gem = Gold
		token.Color = "gold"
		token.Name = "Joker"
	case 's':
		token.Gem = Sapphire
		token.Color = "blue"
		token.Name = "Sapphire"
	case 'o':
		token.Gem = Onyx
		token.Color = "black"
		token.Name = "Onyx"
	case 'd':
		token.Gem = Diamond
		token.Color = "white"
		token.Name = "Diamond"
	case 'r':
		token.Gem = Ruby
		token.Color = "red"
		token.Name = "Ruby"
	default:
		log.Fatalf("unknown Token Code %v", b)
	}
	return token
}
