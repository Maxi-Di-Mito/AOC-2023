package aoc01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type found struct {
	index  int
	word   string
	number int
}

func RunB() {

	file, err := os.Open("problems/aoc-01/data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numberWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		founds := []found{}

		for wordIndex, word := range numberWords {
			indexFirstWord := strings.Index(line, word)
			indexLastWord := strings.LastIndex(line, word)

			if indexFirstWord != -1 {
				value := found{word: word, index: indexFirstWord}
				if wordIndex < 11 {
					value.number = wordIndex
				} else {
					numberConv, _ := strconv.Atoi(word)
					value.number = numberConv
				}
				founds = append(founds, value)
			}

			if indexLastWord != -1 {
				value := found{word: word, index: indexLastWord}
				if wordIndex < 11 {
					value.number = wordIndex
				} else {
					numberConv, _ := strconv.Atoi(word)
					value.number = numberConv
				}
				founds = append(founds, value)
			}
		}

		first := found{index: -1}
		last := found{index: -1}

		for _, value := range founds {
			if first.index == -1 || value.index < first.index {
				first = value
			}
			if last.index == -1 || value.index > last.index {
				last = value
			}
		}

		actualNumber := first.number*10 + last.number

		count += actualNumber

	}

	fmt.Println(count)

}
