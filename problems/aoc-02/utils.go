package aoc02

import (
	"regexp"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	blue  int
	green int
}

type game struct {
	id    int
	list  []hand
	power int
}

func parseLine(line string) game {
	redReg := regexp.MustCompile(` (?P<value>\d\d?) red`)
	blueReg := regexp.MustCompile(` (?P<value>\d\d?) blue`)
	greenReg := regexp.MustCompile(` (?P<value>\d\d?) green`)

	split := strings.Split(line, ":")
	idText := strings.Replace(split[0], "Game ", "", 1)
	id, _ := strconv.Atoi(idText)
	handsText := split[1]
	hands := strings.Split(handsText, ";")

	theGame := game{id: id, list: []hand{}}

	for _, handText := range hands {
		theHand := hand{red: 0, blue: 0, green: 0}
		mat := redReg.FindStringSubmatch(handText)
		if len(mat) > 0 {
			n, _ := strconv.Atoi(mat[1])
			theHand.red = n
		}
		mat = blueReg.FindStringSubmatch(handText)
		if len(mat) > 0 {
			n, _ := strconv.Atoi(mat[1])
			theHand.blue = n
		}
		mat = greenReg.FindStringSubmatch(handText)
		if len(mat) > 0 {
			n, _ := strconv.Atoi(mat[1])
			theHand.green = n
		}
		theGame.list = append(theGame.list, theHand)
	}
	return theGame
}
