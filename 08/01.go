package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("08/input.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	count := 0

	for i := 0; i < len(lines); i++ {
		// fmt.Printf("%v\n", lines[i])
		parts := strings.Split(lines[i], " | ")
		right := parts[1]
		inputs := strings.Fields(right)
		partialCount := 0

		for j := 0; j < len(inputs); j++ {
			switch len(inputs[j]) {
			case 2, 4, 3, 7:
				partialCount++
			}
		}

		count += partialCount

		fmt.Printf("Line %v: %v\n", i, partialCount)
	}

	fmt.Printf("Count: %v\n", count)
}