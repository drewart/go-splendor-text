package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Noble struct {
	Name       string
	TokensGems Gems
	Points     int
}

func (n *Noble) Display() string {
	need := n.TokensGems.String(false)

	return fmt.Sprintf("%-12s (%d) %s", n.Name, n.Points, need)
}

var (
	Nobles []Noble
)

func LoadNobleData() error {

	nobleData, err := ioutil.ReadFile("noble.csv")
	if err != nil {
		return err
	}
	nobleDataStr := string(nobleData)
	lines := strings.Split(nobleDataStr, "\n")
	for i, line := range lines {
		//fmt.Println(i, line)
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			return fmt.Errorf("line error line %d missing comma: %s", i, line)
		}
		name := parts[0]
		tokenStr := parts[1]
		pointStr := parts[2]
		//fmt.Println(name, tokenStr, pointStr)
		points, err := strconv.Atoi(pointStr)
		if err != nil {
			return err
		}
		gemsNeeded := GenGems(tokenStr)
		Nobles = append(Nobles, Noble{Name: name, TokensGems: gemsNeeded, Points: points})
	}
	return nil
}

/*
TODO






var CardOne []DevCard{
		,{TokenNeed=GenTokens("")
	}

*/
