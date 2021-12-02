package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("01/01_input.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	previous := -1
	current := -1

	var strValue string
	var intValue int

	for scanner.Scan() {
		strValue = scanner.Text()
		if intValue, err = strconv.Atoi(strValue); err != nil {
			panic("Error converting str value to int")
		}

		previous = current
		current = intValue

		if (previous != -1 && previous < current) {
			count++
		}

		// fmt.Printf("[%d, %d, %d]\n", previous, current, count)
	}

	fmt.Printf("# of measurements: %d\n", count)
}
