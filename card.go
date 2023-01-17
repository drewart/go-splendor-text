package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var Cards []Card

type Card struct {
	Level int
	Token Token
	Gem   Gem
	//TokenNeeds []TokenCount
	GemCost Gems
	Points  int
}

func (c *Card) Display() string {

	gemCostStr := c.GemCost.String(false)
	return fmt.Sprintf("%-8s (%d) N: %-11s", c.Token.Name, c.Points, gemCostStr)
}

func LoadCards() {

	cardData, err := ioutil.ReadFile("cards.csv")
	if err != nil {
		log.Fatal(err)
	}

	cardDataStr := string(cardData)
	lines := strings.Split(cardDataStr, "\n")
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			log.Fatalf("csv wrong, line %d : %s", i, line)
		}

		levelStr := parts[0]
		level := int(levelStr[0]) - 48
		cardToken := GetToken(parts[1][0])
		tokensStr := parts[2]
		pointStr := parts[3]
		points := int(pointStr[0]) - 48
		gemNeed := GenGems(tokensStr)
		Cards = append(Cards, Card{Level: level, Token: cardToken, GemCost: gemNeed, Points: points})
	}
}
