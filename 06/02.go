package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Simulation struct {
	fish [9]int
	day int
}

func (s *Simulation) Init(inputs []int) {
	s.day = 0

	for i := 0; i < len(inputs); i++ {
		s.fish[inputs[i]]++
	}
}

func (s *Simulation) simulateDays(numDays int) {
	for s.day = 1; s.day <= numDays; s.day++ {
		s.simulateDay()
		s.printSchoolSize()
	}
}

func (s *Simulation) simulateDay() {
	oldCount, tmpCount := 0, 0

	for i := len(s.fish) - 1; i >= 0; i-- {
		tmpCount = s.fish[i]
		s.fish[i] = oldCount
		oldCount = tmpCount
	}

	s.fish[6] += oldCount // restart existing fish count
	s.fish[8] += oldCount // add new fish
}

func (s *Simulation) printSchool() {
	fmt.Printf("Fish: \n")
	for i := 0; i < len(s.fish); i++ {
		fmt.Printf(" [%v, %v]\n", i, s.fish[i])
	}
}

func (s *Simulation) printSchoolSize() {
	count := 0
	for i := 0; i < len(s.fish); i++ {
		count += s.fish[i]
	}
	fmt.Printf("Day %v: %v\n", s.day, count)
}

func main() {
	f, err := os.Open("06/input.txt")
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

	sim := new(Simulation)
	sim.Init(inputs)
	sim.simulateDays(256)
}