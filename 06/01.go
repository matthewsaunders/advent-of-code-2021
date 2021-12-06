package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Simulation struct {
	fish []*Fish
}

func (s *Simulation) Init(inputs []int) {
	s.fish = []*Fish{}

	for i := 0; i < len(inputs); i++ {
		newFish := Fish{
			timer: inputs[i],
		}
		s.fish = append(s.fish, &newFish)
	}
}

func (s *Simulation) simulateDays(numDays int) {
	for i := 1; i <= numDays; i++ {
		s.simulateDay()
		s.printSchoolSize()
	}
}

func (s *Simulation) simulateDay() {
	newFish := []*Fish{}

	for i := 0; i < len(s.fish); i++ {
		s.fish[i].timer--

		if (s.fish[i].timer < 0) {
			s.fish[i].timer = 6
			newFish = append(newFish, &Fish{
				timer: 8,
			})
		}
	}

	s.fish = append(s.fish, newFish...)
}

func (s *Simulation) printSchool() {
	fmt.Printf("Fish:\n")
	for i := 0; i < len(s.fish); i++ {
		fmt.Printf(" [%v, %v]\n", i, s.fish[i].timer)
	}
}

func (s *Simulation) printSchoolSize() {
	fmt.Printf("# of fish: %v\n", len(s.fish))
}

type Fish struct {
	timer int
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
	sim.simulateDays(80)
}