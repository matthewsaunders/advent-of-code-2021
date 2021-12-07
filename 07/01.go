package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(val int) int {
	if (val < 0) { val = -val }
	return val
}

func main() {
	f, err := os.Open("07/input.txt")
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

	strInputs := strings.Split(lines[0], ",")
	inputs := []int{}
	for i := 0; i < len(strInputs); i++ {
		val, _ := strconv.Atoi(strInputs[i])
		inputs = append(inputs, val)
	}

	minIn := inputs[0]
	maxIn := inputs[0]

	for i := 0; i < len(inputs); i++ {
		if inputs[i] < minIn { minIn = inputs[i] }
		if inputs[i] > maxIn { maxIn = inputs[i] }
	}

	fmt.Printf("min: %v\n", minIn)
	fmt.Printf("max: %v\n", maxIn)

	position := minIn
	cost := int(^uint(0) >> 1) // max int

	for i := minIn; i < maxIn; i++ {
		tmpCost := 0

		for j := 0; j < len(inputs); j++ {
			tmpCost += abs(inputs[j] - i)
		}

		if (tmpCost < cost) {
			fmt.Printf("--> new low")
			cost = tmpCost
			position = i
		}

		fmt.Printf("--position: %v\n", i)
		fmt.Printf("--cost: %v\n", tmpCost)
	}

	fmt.Printf("position: %v\n", position)
	fmt.Printf("cost: %v\n", cost)
}