package aoc02

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

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

	}
}
