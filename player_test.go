package main

import (
	"testing"
)

func Test_TakeTokens(t *testing.T) {
	boardGems := GenGems("4e4s4o")
	p := Player{Name: "Test"}
	b := NewBoard(Nobles, Cards, boardGems)
	b.Setup()
	gems := GenGems("1e1s1o")
	beforeBoardGems := b.Tokens
	beforePlayerGems := p.GetTokenGems()

	p.TakeTokens(gems, b)

	afterBoardGemsStr := b.Tokens.String(true)
	if afterBoardGemsStr == beforeBoardGems.String(true) {
		t.Error("board tokens match")
	}
	newPlayerGems := p.GetTokenGems()

	if newPlayerGems.String(true) == beforePlayerGems.String(true) {
		t.Error("player tokens stayed the same")
	}
	wanted := "1e1s0r0d1o0j"
	got := newPlayerGems.String(true)
	if got != wanted {
		t.Errorf("player gems to not match wanted: %s != got: %s", wanted, got)
	}

	wanted = "3e3s0r0d3o0j"
	got = afterBoardGemsStr
	if got != wanted {
		t.Errorf("player gems to not match wanted: %s != got: %s", wanted, got)
	}

}


func Test_CanBuy(t *testing.T) {
	tests := []struct {
		name string
		token string
		cards string
		cost string
		want bool

	}{
		{name: "a1", token: "4s",cards:"",cost:"4s",want:true}
	}

	
	
}
