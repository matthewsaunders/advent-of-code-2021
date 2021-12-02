package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("02/input.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	position := 0
	depth := 0

	var inputs []string
	var command string
	var value int

	for s.Scan() {
		inputs = strings.Fields(s.Text())

		command = inputs[0]
		if value, err = strconv.Atoi(inputs[1]); err != nil {
			panic("Error converting value to int")
		}

		switch command {
		case "forward":
			position += value
		case "up":
			depth -= value
			if depth < 0 { depth = 0 }
		case "down":
			depth += value
		default:
			panic("Error: Unknown sub command given")
		}
	}

	fmt.Printf("position: %d, depth: %d\n", position, depth)
	fmt.Printf("Multiplied together: %d\n", position * depth)
}