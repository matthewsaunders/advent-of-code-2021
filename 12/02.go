package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Simulation struct {
	caves map[string]*Cave
	start *Cave
	end *Cave
}

func (s *Simulation) initSimulation(lines []string) {
	s.caves = make(map[string]*Cave)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		from := parts[0]
		to := parts[1]

		// Find or create caves
		fromCave, foundFrom := s.caves[from]
		toCave, foundTo := s.caves[to]

		if !foundFrom {
			fromCave = createCave(from)
			s.caves[from] = fromCave
		}
		if !foundTo {
			toCave = createCave(to)
			s.caves[to] = toCave
		}

		// Determine start and end caves
		if (from == "start") { s.start = fromCave }
		if (to == "start") { s.start = toCave }
		if (from == "end") { s.end = fromCave }
		if (to == "end") { s.end = toCave }

		// Connect caves
		fromCave.neighbors = append(fromCave.neighbors, toCave)
		toCave.neighbors = append(toCave.neighbors, fromCave)
	}
}

type Cave struct {
	name string
	small bool
	neighbors []*Cave
}

func createCave(name string) *Cave {
	return &Cave{
		name: name,
		small: name == strings.ToLower(name),
		neighbors: []*Cave{},
	}
}

func (s *Simulation) printCaveSystem() {
	fmt.Printf("Caves:\n")

	for _, cave := range s.caves {
		fmt.Printf("  %v ->", cave.name)
		for _, n := range cave.neighbors {
			fmt.Printf(" %v", n.name)
		}
		fmt.Printf("\n")
	}
}

func (s *Simulation) findAllPaths() []string {
	return s.findPaths(s.start, []string{}, false)
}

func (s *Simulation) findPaths(current *Cave, prefix []string, smallVisitedTwice bool) []string {
	paths := []string{}

	// Terminate if we find start again after starting
	if (current == s.start && contains(prefix, current.name)) {
		return paths
	}

	if (current.small && contains(prefix, current.name) && smallVisitedTwice) {
		return paths
	}

	visitedTwice := smallVisitedTwice || (current.small && contains(prefix, current.name))

	// Extend the current path
	path := append(prefix, current.name)

	// Terminate if we are at the end
	if (current == s.end) {
		return append(paths, pathToStr(path))
	}

	// If we have not terminated yet, keep it going
	for _, n := range current.neighbors {
		paths = append(paths, s.findPaths(n, path, visitedTwice)...)
	}

	return paths
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func pathToStr(nodes []string) string {
	str := ""

	for _, node := range nodes {
		str += node + " "
	}

	return str
}

func main() {
	// f, err := os.Open("12/sample_input_3.txt")
	f, err := os.Open("12/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	sim := new(Simulation)
	sim.initSimulation(lines)
	sim.printCaveSystem()
	paths := sim.findAllPaths()

	fmt.Println("Paths:")
	for i, path := range paths {
		fmt.Printf("%v: %v\n", i, path)
	}
}