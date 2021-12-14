package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

const RULES_INDEX = 2

type Simulation struct {
	polymer *list.List
	rules map[string]string
}

func initSimulation(lines []string) *Simulation {
	sim := new(Simulation)
	sim.polymer = list.New()
	sim.rules = make(map[string]string)

	polymer := lines[0]
	for _, r := range polymer {
		char := fmt.Sprintf("%c", r)
		sim.polymer.PushBack(char)
	}

	for i := RULES_INDEX; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		sim.rules[parts[0]] = parts[1]
	}

	return sim
}

func (s *Simulation) PrintPolymer() {
	fmt.Printf("Polymer: ")
	elem := s.polymer.Front()
	for ok := elem != nil; ok; ok = elem != nil {
		fmt.Printf("%v", elem.Value)
		elem = elem.Next()
	}
	fmt.Printf("\n")
}

func (s *Simulation) PrintRules() {
	fmt.Println("Rules:")
	for k, v := range s.rules {
		fmt.Printf(" %v => %v\n", k, v)
	}
}

func (s *Simulation) SimulateIterations(numIterations int) {
	for i := 0; i < numIterations; i++ {
		s.SimulateIteration()
	}
}

func (s *Simulation) SimulateIteration() {
	left := s.polymer.Front()
	for ok := left.Next() != nil; ok; ok = left.Next() != nil {
		right := left.Next()

		key := fmt.Sprintf("%v%v", left.Value, right.Value)
		newElem, _ := s.rules[key]
		s.polymer.InsertAfter(newElem, left)

		left = right
	}
}

func (s *Simulation) CountElements() map[string]int {
	count := make(map[string]int)

	elem := s.polymer.Front()
	for ok := elem != nil; ok; ok = elem != nil {
		// This "fun" hack is needed because the element value is of type
		// interface... yay
		key := fmt.Sprintf("%v", elem.Value)
		count[key]++

		elem = elem.Next()
	}

	return count
}

func getMax(mapp map[string]int) int {
	max := 0

	for _, v := range mapp {
		if (v > max) { max = v }
	}

	return max
}

func getMin(mapp map[string]int) int {
	min := int(^uint(0) >> 1) // max int value

	for _, v := range mapp {
		if (v < min) { min = v }
	}

	return min
}

func main() {
	// f, err := os.Open("14/sample_input.txt")
	f, err := os.Open("14/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	sim := initSimulation(lines)
	sim.PrintRules()
	sim.PrintPolymer()
	sim.SimulateIterations(10)

	counts := sim.CountElements()
	max := getMax(counts)
	min := getMin(counts)
	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Min: %v\n", min)
	fmt.Printf("Answer: %v\n", max - min)
}
