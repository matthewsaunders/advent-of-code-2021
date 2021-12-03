package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	const SIZE = 12
	var counts [SIZE][]int
	var token string
	var index int
	var gamma_bit int
	var epsilon_bit int

	gamma_rate := 0
	epsilon_rate := 0

	counts = [SIZE][]int{}

	for i := 0; i < SIZE; i++ {
		counts[i] = []int{ 0, 0 }
	}

	for s.Scan() {
		token = s.Text()
		fmt.Printf("Token: %v\n", token)

		for i, c := range token {
			fmt.Printf("  i: %v, c: %c\n", i, c)
			if index = 0; int(c) == 49 { index = 1 }
			counts[i][index]++
		}
	}

	for i, arr := range counts {
		fmt.Printf("i: %v, [%v, %v]\n", i, arr[0], arr[1])
		
		gamma_bit = 0
		epsilon_bit = 0

		if (arr[0] > arr[1]) {
			gamma_bit = 1
		} else {
			epsilon_bit = 1
		}

		gamma_rate = gamma_rate << 1 | gamma_bit
		epsilon_rate = epsilon_rate << 1 | epsilon_bit
	}

	fmt.Printf("gamma: %v\n", gamma_rate)
	fmt.Printf("epsilon: %v\n", epsilon_rate)
	fmt.Printf("product: %v\n", gamma_rate * epsilon_rate)
}
