package main

import (
	"flag"
	"strings"
)

var (
	Debug = true
)

/*
 */

func main() {

	var players string
	var playerNames []string
	flag.StringVar(&players, "players", "", "players command delimited")
	flag.Parse()
	if players != "" {
		playerNames = strings.Split(players, ",")
	}
	splendor := Splendor{}
	for _, name := range playerNames {
		splendor.Players = append(splendor.Players, Player{Name: name})
	}
	splendor.Setup()
	splendor.Play()
}
