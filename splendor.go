package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Splendor struct {
	GameBoard       *Board
	Players         []Player
	RoundNum        int
	PlayerTurnIndex int
}

func (s Splendor) Usage() {
	commands := `
e - Emerald
s - Sapphire
r - Ruby
d - Diamon
o - Onyx
j - Joker (Gold)

tokens 1e1r - 
buy <level> <position> - buy card
quit
`
	fmt.Println(commands)
}

func (s *Splendor) Play() {
	// record play
	// Setup
	// Rounds
	// Player turn
	// youngest goes first
	var command string
	var currentPlayer *Player
	turnOver := false
	reader := bufio.NewReader(os.Stdin)
	for {
		currentPlayer = &s.Players[s.PlayerTurnIndex]
		fmt.Printf("r:%d: %s> ", s.RoundNum, currentPlayer.Name)
		//fmt.Scanln()
		command, _ = reader.ReadString('\n')
		args := strings.Fields(command)
		switch args[0] {
		case "quit":
			// TODO save ?
			os.Exit(0)
		case "tokens":
			if len(args) < 2 {
				fmt.Println("no tokens")
				break
			}
			tokensStr := args[1]
			gems := GenGems(tokensStr)
			currentPlayer.TakeTokens(gems, s.GameBoard)
			turnOver = true
		case "buy":
			if len(args) < 3 {
				fmt.Println("no buy level position")
				break
			}
			levelStr := args[1]
			posStr := args[2]
			lvl, _ := strconv.Atoi(levelStr)
			pos, _ := strconv.Atoi(posStr)
			card := s.GameBoard.GetCard(lvl, pos)
			if !currentPlayer.CanBuy(card) {
				fmt.Println("can not buy card")
			} else {
				currentPlayer.BuyCard(card, s.GameBoard)
			}
		case "help":
			s.Usage()
		case "done":
			turnOver = true
		default:
			fmt.Println("Unknown command", command)
		}

		if turnOver {
			s.PlayerTurnIndex = s.PlayerTurnIndex + 1
			if s.PlayerTurnIndex > len(s.Players)-1 {
				s.RoundNum += 1
				s.PlayerTurnIndex = 0
			}
			turnOver = false
			s.GameBoard.RenderBoard()
			s.DisplayPlayerHands()
		}

	}

}

func (s *Splendor) PlayerCount() int {
	return len(s.Players)
}

func (s *Splendor) DisplayPlayerHands() {
	for _, p := range s.Players {
		cardGems := p.GetCardGems()
		tokenGems := p.GetTokenGems()
		fmt.Printf("N:%-20s C:(%s) T:(%s)\n", p.Name, cardGems.String(true), tokenGems.String(true))
	}
}

func (s *Splendor) PlayerSetup() {
	// pick players
	var playerCountStr string
	fmt.Print("How many Players? > ")
	fmt.Scanln(&playerCountStr)
	playerCnt, err := strconv.Atoi(playerCountStr)
	if err != nil {
		log.Fatalln(err)
	}
	if playerCnt > 4 {
		fmt.Println("Player count too big", playerCnt)
		return
	}
	var playerName string
	for c := 0; c < playerCnt; c++ {
		fmt.Printf("Player Number #%d Name? > ", c+1)
		fmt.Scanln(&playerName)
		player := Player{Name: playerName}
		s.Players = append(s.Players, player)
	}
}

func (s *Splendor) Setup() {
	s.RoundNum = 1
	fmt.Println("setting up ")
	fmt.Println("load nobles")
	err := LoadNobleData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Noble Count", len(Nobles))
	if Debug {
		for i, n := range Nobles {
			fmt.Println(i, n)
		}
	}
	fmt.Println("loading cards")
	LoadCards()

	fmt.Println("Cards", len(Cards))

	if s.PlayerCount() == 0 {
		s.PlayerSetup()
	}

	// 7 = 4, 5 = 3, 4 = 2

	var tokenGems Gems
	var tokenCount = 7
	if len(s.Players) == 2 {
		tokenCount = 4
	} else if len(s.Players) == 3 {
		tokenCount = 5
	}

	for _, gem := range GemList {
		if gem == Gold {
			tokenGems[gem] = 5
		} else {
			tokenGems[gem] = tokenCount
		}
	}

	/*tokens := GenTokenCount("7s7e7d7o5j")
	if len(s.Players) == 2 {
		tokens = GenTokenCount("4s4e4d4o5j")
	} else if len(s.Players) == 3 {
		tokens = GenTokenCount("5s5e5d5o5j")
	}*/
	// pick
	// pick Player Nobles
	//

	//shuffle decks
	deckCards := make([]Card, len(Cards))
	perm := rand.Perm(len(Cards))
	for i, v := range perm {
		deckCards[v] = Cards[i]
	}

	//shuffle nobels
	nobleTiles := make([]Noble, len(Nobles))
	perm = rand.Perm(len(Nobles))
	for i, v := range perm {
		nobleTiles[v] = Nobles[i]
	}

	s.GameBoard = NewBoard(nobleTiles[0:s.PlayerCount()+1], deckCards, tokenGems)
	if Debug {
		s.GameBoard.DebugNobles()
		s.GameBoard.DebugDecks()
	}
	s.GameBoard.Deal()
	s.GameBoard.RenderBoard()
	s.DisplayPlayerHands()
}
