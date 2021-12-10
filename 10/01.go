package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []rune

func (s *Stack) Push(elem rune) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (rune, bool) {
	if (s.IsEmpty()) { return '0', false }

	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func charsMap() map[rune]rune {
	myMap := make(map[rune]rune)

	myMap['('] = ')'
	myMap['['] = ']'
	myMap['{'] = '}'
	myMap['<'] = '>'

	return myMap
}

func pointsMap() map[rune]int {
	myMap := make(map[rune]int)

	myMap[')'] = 3
	myMap[']'] = 57
	myMap['}'] = 1197
	myMap['>'] = 25137

	return myMap
}

func processLineScore(line string) int {
	var stack Stack
	charMap := charsMap()
	score := 0

	fmt.Println(line)

	checkLine:
	for _, char := range line {
		if _, found := charMap[char]; found {
			stack.Push(char)
		} else {
			elem, ok := stack.Pop()
			if (ok && charMap[elem] != char) {
				fmt.Printf("-- Expected %c, but found %c instead.\n", charMap[elem], char)
				score = pointsMap()[char]
				break checkLine
			}
		}
	}

	return score
}

func main() {
	// f, err := os.Open("10/sample_input.txt")
	f, err := os.Open("10/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	points := 0

	for _, line := range lines {
		points += processLineScore(line)
	}

	fmt.Printf("points: %v\n", points)
}