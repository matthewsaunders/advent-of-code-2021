package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// This solution feels inefficient and dirty. Im going to take a shower now.
func main() {
	var f *os.File
	var err error
	var s *bufio.Scanner
	
	if f, err = os.Open("03/input.txt"); err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}

	defer f.Close()

	s = bufio.NewScanner(f)

	originalTokens := []string{}
	newTokens := []string{}
	val := 0
	var values [2]int
	var selectedVal int

	for s.Scan() {
		originalTokens = append(originalTokens, s.Text())
	}

	index := 0
	tokens := originalTokens

	// oxygen
	for ok:= true; ok; ok = len(tokens) != 1 {
		fmt.Printf("--> index: %v\n", index)
		values = [2]int{ 0, 0 }

		for i := 0; i < len(tokens); i++ {
			fmt.Printf("Token: %v, Char: %c\n", tokens[i], tokens[i][index])
			val = 0
			if int(tokens[i][index]) == 49 { val = 1 }
			values[val]++
		}

		selectedVal = 1
		if (values[1] < values[0]) { selectedVal = 0 }

		fmt.Printf("index: %v, [%v, %v]\n", index, values[0], values[1])
		fmt.Printf("  --> selected %v\n", selectedVal)

		newTokens = []string{}

		for i := 0; i < len(tokens); i++ {
			val = 0
			if int(tokens[i][index]) == 49 { val = 1 }
			if val == selectedVal { newTokens = append(newTokens, tokens[i]) }
		}

		fmt.Printf("Token count: %v\n", len(newTokens))
		tokens = newTokens
		index++
	}

	oxygenRating := tokens[0]

	index = 0
	tokens = originalTokens

	// co2
	for ok:= true; ok; ok = len(tokens) != 1 {
		fmt.Printf("--> index: %v\n", index)
		values = [2]int{ 0, 0 }

		for i := 0; i < len(tokens); i++ {
			fmt.Printf("Token: %v, Char: %c\n", tokens[i], tokens[i][index])
			val = 0
			if int(tokens[i][index]) == 49 { val = 1 }
			values[val]++
		}

		selectedVal = 0
		if (values[1] < values[0]) { selectedVal = 1 }

		fmt.Printf("index: %v, [%v, %v]\n", index, values[0], values[1])
		fmt.Printf("  --> selected %v\n", selectedVal)

		newTokens = []string{}

		for i := 0; i < len(tokens); i++ {
			val = 0
			if int(tokens[i][index]) == 49 { val = 1 }
			if val == selectedVal { newTokens = append(newTokens, tokens[i]) }
		}

		fmt.Printf("Token count: %v\n", len(newTokens))
		tokens = newTokens
		index++
	}

	co2Rating := tokens[0]

	fmt.Printf("oxygen rating: %v\n", oxygenRating)
	fmt.Printf("co2 rating: %v\n", co2Rating)

	oxygenNum, err := strconv.ParseInt(oxygenRating, 2, 64)
	if err != nil {
		fmt.Println(err)
		panic("Error converting oxygenRating")
	}

	co2Num, err := strconv.ParseInt(co2Rating, 2, 64)
	if err != nil {
		fmt.Println(err)
		panic("Error converting co2Rating")
	}

	fmt.Printf("product: %v\n", oxygenNum * co2Num)
}
