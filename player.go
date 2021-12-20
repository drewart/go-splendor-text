package main

type Player struct {
	Name         string
	TokenGems    Gems
	Cards        []*Card
	ReserveCards []*Card
	Nobles       []Noble
}

func (p *Player) GetCardGems() Gems {
	var gems Gems
	for _, c := range p.Cards {
		g := c.Token.Gem
		gems[g] += 1
	}
	return gems
}

func (p *Player) GetTokenGems() Gems {
	return p.TokenGems
}

func (p *Player) CanBuy(c *Card) bool {
	var gems Gems
	canBuy := true
	tokenGems := p.GetTokenGems()
	cardGems := p.GetCardGems()
	for _, g := range GemList {
		gems[g] = tokenGems[g] + cardGems[g]
		// if card gem cost > gem count
		if c.GemCost[g] > gems[g] {
			canBuy = false
		}
	}
	return canBuy
}

func (p *Player) TokenCost(c *Card) Gems {
	var gems Gems
	cardGems := p.GetCardGems()
	for _, g := range GemList {
		// gem or token cost remander
		gems[g] = c.GemCost[g] - cardGems[g]
	}
	return gems
}

func (p *Player) BuyCard(c *Card, b *Board) {
	if p.CanBuy(c) {
		// assume CanBuy == true
		tokenCost := p.TokenCost(c)
		for _, g := range GemList {
			count := tokenCost[g]
			if count > 0 {
				p.ChangeTokens(g, 0-count)
				b.ReplaceToken(GetGemToken(g), count)
				p.Cards = append(p.Cards, c)
				b.FindReplaceCard(c)
			}
		}
	}
}

func (p *Player) ChangeTokens(gem Gem, count int) {
	p.TokenGems[gem] += count

}

func (p *Player) TotalTokens() int {
	var count int
	for i := 0; i < len(p.TokenGems); i++ {
		count += p.TokenGems[i]
	}
	return count
}

func (p *Player) TakeTokens(gems Gems, b *Board) {
	for _, gem := range GemList {
		count := gems[gem]
		if count > 0 {
			for i := 0; i < count; i++ {
				token := b.TakeToken(gem)
				// handle empty token stack
				if token != nil {
					p.ChangeTokens(gem, 1)
				}
			}
		}
	}
}
