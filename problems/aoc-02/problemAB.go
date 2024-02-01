package aoc02

import (
	"bufio"
	"fmt"
	"os"
)

func RunAB() {
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
	powerSum := 0

	for _, game := range gameList {
		valid := true

		var redMinumun = 0
		var greenMinumun = 0
		var blueMinumun = 0

		for _, hand := range game.list {
			if hand.red > redMinumun {
				redMinumun = hand.red
			}
			if hand.green > greenMinumun {
				greenMinumun = hand.green
			}
			if hand.blue > blueMinumun {
				blueMinumun = hand.blue
			}
			// fmt.Println("HAND", hand)
			if valid && (hand.red > maxRed || hand.blue > maxBlue || hand.green > maxGreen) {
				valid = false
			}
		}

		game.power = redMinumun * greenMinumun * blueMinumun
		powerSum += game.power
		// fmt.Println("VALID", valid)
		if valid {
			count += game.id
		}
	}

	fmt.Println("Sum of valid gameIds", count)
	fmt.Println("Sum of game powers", powerSum)
}
