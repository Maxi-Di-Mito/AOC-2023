package aoc02

import (
	"bufio"
	"fmt"
	"os"
)

type hand struct {
	red   int
	blue  int
	green int
}

type game struct {
	id   int
	list []hand
}

func RunA() {
	file, err := os.Open("problems/aoc-02/data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	gameList := []game{}

	for scanner.Scan() {
		line := scanner.Text()

		theGame := parseLine(line)

		gameList = append(gameList, theGame)
	}

	count := 0

	for _, game := range gameList {
		valid := true
		for _, hand := range game.list {
			fmt.Println("HAND", hand)
			if hand.red > maxRed || hand.blue > maxBlue || hand.green > maxGreen {
				valid = false
				break
			}
		}
		fmt.Println("VALID", valid)
		if valid {
			count += game.id
		}
	}

	fmt.Println(count)
}
