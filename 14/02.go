package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const RULES_INDEX = 2

type Simulation struct {
	polymer map[string]int
	rules map[string]string
}

func initSimulation(lines []string) *Simulation {
	sim := new(Simulation)
	sim.polymer = make(map[string]int)
	sim.rules = make(map[string]string)

	polymer := lines[0]
	for i := 0; i < len(polymer) - 1; i++ {
		key := fmt.Sprintf("%c%c", polymer[i], polymer[i+1])
		sim.polymer[key] = 1
	}

	for i := RULES_INDEX; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		sim.rules[parts[0]] = parts[1]
	}

	return sim
}

func (s *Simulation) PrintRules() {
	fmt.Println("Rules:")
	for k, v := range s.rules {
		fmt.Printf(" %v => %v\n", k, v)
	}
}

func (s *Simulation) PrintPolymer() {
	fmt.Println("Polymer:")
	for k, v := range s.polymer {
		fmt.Printf(" %v => %v\n", k, v)
	}
}

func (s *Simulation) SimulateIterations(numIterations int) {
	count := 1
	for i := 0; i < numIterations; i++ {
		fmt.Printf("--iteration %v\n", count)
		count++
		s.SimulateIteration()
	}
}

func (s *Simulation) SimulateIteration() {
	newPolymer := make(map[string]int)
	for k, v := range s.polymer {
		newPolymer[k] = v
	}

	for k, v := range s.polymer {
		newElem, _ := s.rules[k]
		left := fmt.Sprintf("%c%v", k[0], newElem)
		right := fmt.Sprintf("%v%c", newElem, k[1])
		newPolymer[k] -= v
		newPolymer[left] += v
		newPolymer[right] += v
	}

	s.polymer = newPolymer
}

func (s *Simulation) PolymerLength() {
	count := 1

	for _, v := range s.polymer {
		count += v
	}

	fmt.Printf("Polymer length: %v\n", count)
}

func (s *Simulation) CountElements() map[string]int {
	counts := make(map[string]int)

	for k, v := range s.polymer {
		left := fmt.Sprintf("%c", k[0])
		right := fmt.Sprintf("%c", k[1])

		// Each letter is counted twice because it is in 2 pair. Int division
		// rounds up the literal edge cases where it should only be counted once
		counts[left] += v / 2
		counts[right] += v / 2
	}

	return counts
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
	sim.SimulateIterations(40)

	counts := sim.CountElements()
	max := getMax(counts)
	min := getMin(counts)
	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Min: %v\n", min)
	fmt.Printf("Answer: %v\n", max - min)
}
