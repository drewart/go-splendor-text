package main

import (
	"fmt"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	DeckEmpty   = Error("deck empty")
	CardMissing = Error("card missing")
)

type Board struct {
	Nobles       []Noble
	DevCardDecks [][]Card
	DeckIndex    [3]int
	ShowCards    [][]*Card
	Tokens       Gems
}

func NewBoard(nobles []Noble, cards []Card, gems Gems) *Board {

	var board Board

	board.Nobles = nobles
	board.Tokens = gems
	board.DevCardDecks = make([][]Card, 3)
	board.ShowCards = make([][]*Card, 3)
	for l := 0; l < 3; l++ {
		var deckCards []Card
		for _, card := range cards {
			if card.Level == l+1 {
				deckCards = append(deckCards, card)
			}
		}

		board.DevCardDecks[l] = make([]Card, 0)
		board.DevCardDecks[l] = append(board.DevCardDecks[l], deckCards...)
		board.DeckIndex[l] = len(deckCards) - 1
	}
	for i := 0; i < 3; i++ {
		board.ShowCards[i] = make([]*Card, 4)
	}
	return &board
}

func (b *Board) Setup() {

}

func (b *Board) Deal() {
	// take top deck
	for i := 0; i < 4; i++ {
		for level, index := range b.DeckIndex {
			b.ShowCards[level][i] = &b.DevCardDecks[level][index-i]
		}
	}
	for y := 0; y < len(b.DeckIndex); y++ {
		b.DeckIndex[y] = b.DeckIndex[y] - 4
	}
}

func (b *Board) DisplayCards() {
	for level := len(b.ShowCards); level > 0; level-- {
		fmt.Printf("Deck %d: (%02d) | ", level, b.DeckIndex[level-1])
		for y := 0; y < len(b.ShowCards[level-1]); y++ {
			fmt.Print(b.ShowCards[level-1][y].Display(), " | ")
		}
		fmt.Println()
	}
}

func (b *Board) RenderBoard() {
	fmt.Print("\nNobles:\n")
	for i := 0; i < len(b.Nobles); i++ {
		fmt.Printf("\t%s\n", b.Nobles[i].Display())
	}
	fmt.Printf("\n\n")

	fmt.Println("Development Cards")
	b.DisplayCards()

	fmt.Print("\nTokens: ")
	for i := 0; i < len(b.Tokens); i++ {
		fmt.Printf("%s:%d ", Gem(i).String(), b.Tokens[i])
	}
	fmt.Println()
}

//
func (b *Board) GetTokenCount(g Gem) int {
	return b.Tokens[g]
}

func (b *Board) TakeToken(g Gem) *Token {
	var token *Token
	count := b.GetTokenCount(g)
	if count == 0 {
		return nil
	} else {
		b.Tokens[g] = count - 1
		gemToken := GetGemToken(g)
		token = &gemToken
	}
	return token
}

func (b *Board) TakeTokens(gems []Gem) []Token {
	var tokens []Token
	for _, g := range gems {
		t := b.TakeToken(g)
		if t != nil {
			tokens = append(tokens, *t)
		}
	}
	return tokens
}

func (b *Board) ReplaceToken(t Token, count int) {
	b.Tokens[t.Gem] += count
}

func (b *Board) Reserve(level, pos int) (*Card, *Token, error) {
	if pos == 0 {
		if b.DeckIndex[level-1] < 0 {
			return nil, nil, DeckEmpty
		} else {
			index := b.DeckIndex[level-1]
			b.DeckIndex[level-1] = b.DeckIndex[level-1] - 1
			gold := b.TakeToken(Gold)
			return &b.DevCardDecks[level-1][index], gold, nil
		}
	}
	if b.ShowCards[level][pos-1] == nil {
		return nil, nil, CardMissing
	} else {
		gold := b.TakeToken(Gold)
		return b.ShowCards[level][pos], gold, nil
	}
}

func (b *Board) replaceCard(level, pos int) {
	if b.DeckIndex[level-1] >= 0 {
		index := b.DeckIndex[level-1]
		b.ShowCards[level-1][pos-1] = &b.DevCardDecks[level-1][index]
		b.DeckIndex[level-1] = b.DeckIndex[level-1] - 1
	} else if b.DeckIndex[level-1] < 0 {
		b.ShowCards[level-1][pos-1] = nil
	}
}

func (b *Board) FindReplaceCard(card *Card) {
	levelIndex := card.Level - 1
	for i := 0; i < len(b.ShowCards[levelIndex]); i++ {
		if card.Display() == b.ShowCards[levelIndex][i].Display() {
			b.replaceCard(card.Level, i+1)
		}
	}
}

func (b *Board) CanBuyCard(tokenHave map[Gem]int) bool {
	return false
}

func (b *Board) GetCard(level, pos int) *Card {
	var c *Card
	if b.ShowCards[level][pos-1] == nil {
		return nil
	} else {
		c = b.ShowCards[level-1][pos-1]
	}
	return c
}

// Pick card from table
func (b *Board) PickCard(level, pos int) (*Card, error) {
	if b.ShowCards[level][pos-1] == nil {
		return nil, CardMissing
	} else {
		c := b.ShowCards[level-1][pos-1]
		b.replaceCard(level, pos)

		return c, nil
	}
}

func (b *Board) DebugNobles() {
	fmt.Println("Nobles Picked")
	for i := 0; i < len(b.Nobles); i++ {
		fmt.Println(b.Nobles[i])
	}

}

func (b *Board) DebugDecks() {
	fmt.Println("Decks")
	for i := 0; i < len(b.DevCardDecks); i++ {
		for y := 0; y < len(b.DevCardDecks[i]); y++ {
			fmt.Println(i, y, b.DevCardDecks[i][y].Display())
		}
	}
}
