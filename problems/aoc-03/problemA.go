package aoc03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	x     int
	y     int
	value string
}

func (n *node) getUppper(content []row) *node {
	upperY := n.y - 1
	if upperY < 0 {
		return nil
	}

	return &content[n.x].chars[upperY]
}

func (n *node) getDowner(content []row) *node {
	downerY := n.y + 1
	if downerY >= len(content) {
		return nil
	}

	return &content[n.x].chars[downerY]
}

func (n *node) getLefter(content []row) *node {
	lefterX := n.x - 1
	if lefterX < 0 {
		return nil
	}

	return &content[lefterX].chars[n.y]
}
func (n *node) getRighter(content []row) *node {
	righterX := n.x + 1
	if righterX >= len(content[n.y].chars) {
		return nil
	}

	return &content[righterX].chars[n.y]
}

var regSim = regexp.MustCompile(`[%*#&%@=\/\-\$+]`)
var regNum = regexp.MustCompile(`[0-9]`)

func checkSurroundings(n *node, content []row) bool {
	nodes := []*node{}

	nodes = append(nodes, n.getUppper(content))
	if n.getUppper(content) != nil {
		nodes = append(nodes, n.getUppper(content).getLefter(content))
		nodes = append(nodes, n.getUppper(content).getRighter(content))
	}
	nodes = append(nodes, n.getLefter(content))
	nodes = append(nodes, n.getRighter(content))
	nodes = append(nodes, n.getDowner(content))
	if n.getDowner(content) != nil {
		nodes = append(nodes, n.getDowner(content).getLefter(content))
		nodes = append(nodes, n.getDowner(content).getRighter(content))
	}

	for _, node := range nodes {
		if node != nil && regSim.MatchString(node.value) {
			return true
		}
	}

	return false
}

type row struct {
	chars []node
	y     int
}

func RunA() {
	file, err := os.Open("problems/aoc-03/data.txt")
	if err != nil {
		panic(err)
	}

	var lines []row

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rowCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		chars := strings.Split(line, "")
		var nodes []node

		for index, char := range chars {
			nodes = append(nodes, node{
				x:     index,
				y:     rowCount,
				value: char,
			})
		}

		lines = append(lines, row{
			chars: nodes,
			y:     rowCount,
		})

		rowCount++
	}

	sum := 0

	for lineIndex, line := range lines {

		currentNumbers := []string{}
		hasSimbol := false
		lineSum := 0

		for charIndex, char := range line.chars {
			fmt.Printf("UN CHAR %d:%d = %s\n", lineIndex, charIndex, char.value)
			if regNum.MatchString(char.value) {
				currentNumbers = append(currentNumbers, char.value)
				if !hasSimbol {
					hasSimbol = checkSurroundings(&char, lines)
					if hasSimbol {
						fmt.Println("HAS SIMBOL TRUE")
					}
				}

			} else if len(currentNumbers) > 0 && hasSimbol {
				stringNum := strings.Join(currentNumbers, "")
				num, err := strconv.Atoi(stringNum)
				if err != nil {
					lineSum += num
				}
				fmt.Println("CURRENT NUM:", num)
				currentNumbers = []string{}
				hasSimbol = false
			}
		}

		sum += lineSum

	}
	fmt.Println("TOTAL", sum)

}
