package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	myMap[')'] = 1
	myMap[']'] = 2
	myMap['}'] = 3
	myMap['>'] = 4

	return myMap
}

func isLineBroken(line string) bool {
	var stack Stack
	charMap := charsMap()
	broken := false

	fmt.Println(line)

	checkLine:
	for _, char := range line {
		if _, found := charMap[char]; found {
			stack.Push(char)
		} else {
			elem, ok := stack.Pop()
			if (ok && charMap[elem] != char) {
				fmt.Printf("-- Expected %c, but found %c instead.\n", charMap[elem], char)
				broken = true
				break checkLine
			}
		}
	}

	return broken
}

func completeLine(line string) string {
	var stack Stack
	charMap := charsMap()

	fmt.Println(line)

	for _, char := range line {
		if _, found := charMap[char]; found {
			stack.Push(char)
		} else {
			stack.Pop()
		}
	}

	missingStr := ""

	for ok := true; ok; ok = !stack.IsEmpty() {
		token, _ := stack.Pop()
		missingStr = fmt.Sprintf("%v%c", missingStr, charMap[token])
	}

	return missingStr
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

	incompleteLines := []string{}

	for _, line := range lines {
		if (!isLineBroken(line)) {
			incompleteLines = append(incompleteLines, line)
		}
	}

	pointsMap := pointsMap()
	scores := []int{}

	for _, line := range incompleteLines {
		missing := completeLine(line)
		fmt.Printf("-- %v\n", missing)

		points := 0

		for _, char := range missing {
			points = (points * 5) + pointsMap[char]
		}

		fmt.Printf("-- %v\n", points)
		scores = append(scores, points)
	}

	sort.Ints(scores)
	middle := (len(scores)) / 2
	fmt.Printf("%v/%v\n", middle, len(scores))
	fmt.Printf("points: %v\n", scores[middle])
}