package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const IGNORE_TEXT = "fold along "

type Simulation struct {
	grid [][]bool
	instructions []*Instruction
}

type Instruction struct {
	x bool
	index int
}

func (s *Simulation) printInstructions() {
	fmt.Printf("Instructions:\n")
	for i, line := range s.instructions {
		fmt.Printf(" %v: %v\n", i, line)
	}
}

func (s *Simulation) printGrid() {
	fmt.Printf("Grid:\n")
	for j := 0; j < len(s.grid[1]); j++ {
		for i := 0; i < len(s.grid); i++ {
			char := "."
			if (s.grid[i][j]) { char = "#" }
			fmt.Printf("%v", char)
		}
		fmt.Printf("\n")
	}
}

func (s *Simulation) followInstructions() {
	// for _, inst := range s.instructions {
	// 	if (inst.x) {
	// 		fmt.Printf("--> Folding @ x=%v\n", inst.index)
	// 		s.xFold(inst.index)
	// 	} else {
	// 		fmt.Printf("--> Folding @ y=%v\n", inst.index)
	// 		s.yFold(inst.index)
	// 	}
	// }

	inst := s.instructions[0]
	if (inst.x) {
		fmt.Printf("--> Folding @ x=%v\n", inst.index)
		s.xFold(inst.index)
	} else {
		fmt.Printf("--> Folding @ y=%v\n", inst.index)
		s.yFold(inst.index)
	}
}

func (s *Simulation) xFold(foldIndex int) {
	newGrid := makeGrid(foldIndex, len(s.grid[0]))

	for i := 0; i < len(s.grid); i++ {
		for j := 0; j < len(s.grid[i]); j++ {
			if (i != foldIndex) {
				newGridI := i
				if (i > foldIndex) { newGridI = abs(i - (2 * foldIndex)) }
				newGrid[newGridI][j] = newGrid[newGridI][j] || s.grid[i][j]
			}
		}
	}

	s.grid = newGrid
}

func (s *Simulation) yFold(foldIndex int) {
	newGrid := makeGrid(len(s.grid), foldIndex)

	for i := 0; i < len(s.grid); i++ {
		for j := 0; j < len(s.grid[i]); j++ {
			if (j != foldIndex) {
				newGridJ := j
				if (j > foldIndex) { newGridJ = abs(j - (2 * foldIndex)) }
				newGrid[i][newGridJ] = newGrid[i][newGridJ] || s.grid[i][j]
			}
		}
	}

	s.grid = newGrid
}

func (s *Simulation) numDots() int {
	count := 0

	for i := 0; i < len(s.grid); i++ {
		for j := 0; j < len(s.grid[i]); j++ {
			if (s.grid[i][j]) { count++ }
		}
	}

	return count
}

func abs(x int) int {
	if ( x < 0 ) { x = -x }
	return x
}

func makeGrid(xSize, ySize int) [][]bool {
	grid := make([][]bool, xSize)
	for i := range grid {
		grid[i] = make([]bool, ySize)
		for j := range grid[i] {
			grid[i][j] = false
		}
	}
	return grid
}

func initSimulation(lines []string) *Simulation {
	sim := new(Simulation)

	// Init grid
	gridPoints := [][]int{}
	foldIndex := 0

	loadGridPoints:
	for i, line := range lines {
		if (line == "") {
			foldIndex = i + 1
			break loadGridPoints
		}

		parts := strings.Split(line, ",")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		gridPoints = append(gridPoints, []int{left, right})
	}

	xMax := 0
	yMax := 0
	for _, point := range gridPoints {
		if (point[0] > xMax) { xMax = point[0] }
		if (point[1] > yMax) { yMax = point[1] }
	}

	fmt.Printf("Max points:\n")
	fmt.Printf(" x: %v\n", xMax)
	fmt.Printf(" y: %v\n", yMax)

	sim.grid = makeGrid(xMax + 1, yMax + 1)

	for _, point := range gridPoints {
		x, y := point[0], point[1]
		sim.grid[x][y] = true
	}

	// Init instructions
	for i := foldIndex; i < len(lines); i++ {
		line := lines[i][len(IGNORE_TEXT):]
		parts := strings.Split(line, "=")
		index, _ := strconv.Atoi(parts[1])
		instruction := Instruction{
			x: parts[0] == "x",
			index: index,
		}
		sim.instructions = append(sim.instructions, &instruction)
	}

	return sim
}

func main() {
	// f, err := os.Open("13/sample_input.txt")
	f, err := os.Open("13/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	sim := initSimulation(lines)
	// sim.printGrid()
	sim.printInstructions()
	sim.followInstructions()

	fmt.Printf("count: %v\n", sim.numDots())
}