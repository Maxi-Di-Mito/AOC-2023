package aoc01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunA() {
	file, err := os.Open("problems/aoc-01/data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	reg := regexp.MustCompile(`[0-9]`)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstDigit := reg.FindString(line)
		lastDigit := reg.FindString(reverse(line))

		stringNumber := strings.Join([]string{firstDigit, lastDigit}, "")

		num, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}

		count += num
	}

	fmt.Println("RESULT", count)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
