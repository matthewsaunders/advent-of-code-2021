package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("01/input.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	readValuesCount := 0
	measurementCount := 0
	const WINDOW_SIZE = 4
	window := [WINDOW_SIZE]int{ 0, 0, 0, 0 }
	index := 0
	currentIndex := index

	var strValue string
	var intValue int

	var previousWindow int
	var currentWindow int

	for scanner.Scan() {
		strValue = scanner.Text()
		if intValue, err = strconv.Atoi(strValue); err != nil {
			panic("Error converting str value to int")
		}

		currentIndex = index
		window[index] = intValue
		index = (index + 1) % WINDOW_SIZE
		readValuesCount++

		// fmt.Printf("index: %d, window: [%d %d %d %d]\n", currentIndex, window[0], window[1], window[2], window[3])

		if (readValuesCount >= WINDOW_SIZE) {
			currentWindow = window[currentIndex] + window[(currentIndex + 3) % WINDOW_SIZE] + window[(currentIndex + 2) % WINDOW_SIZE]
			previousWindow = window[(currentIndex + 3) % WINDOW_SIZE] + window[(currentIndex + 2) % WINDOW_SIZE] + window[(currentIndex + 1) % WINDOW_SIZE]

			// fmt.Printf("previousWindow: %d, currentWindow: %d\n", previousWindow, currentWindow)

			if (currentWindow > previousWindow) { measurementCount++ }
		}
	}

	fmt.Printf("# of measurements: %d\n", measurementCount)
}
