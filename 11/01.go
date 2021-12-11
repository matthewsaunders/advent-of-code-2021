package main

import (
	"bufio"
	"fmt"
	"os"
)

const ASCII_ZERO = 48

type Point struct {
	x int
	y int
	value int
}

func (p *Point) key() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

type Stack []*Point

func (s *Stack) Push(elem *Point) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (*Point, bool) {
	if (s.IsEmpty()) { return nil, false }

	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type Simulation struct {
	grid [][]*Point
	flashes int
}

func (s *Simulation) printGrid() {
	for _, row := range s.grid {
		for _, cell := range row {
			fmt.Printf("%v", cell.value)
		}
		fmt.Printf("\n")
	}
}

func (s *Simulation) SimulateDays(numDays int) {
	for i := 0; i < numDays ; i++ {
		s.SimulateDay()
		fmt.Printf("Day %v: %v\n", i, s.flashes)
	}
}

func (s *Simulation) SimulateDay() {
	points := Stack{}
	for _, row := range s.grid {
		for _, point := range row {
			points.Push(point)
		}
	}

	visited := make(map[string]bool)
	for ok := true; ok; ok = !points.IsEmpty() {
		point, _ := points.Pop()
		if _, found := visited[point.key()]; !found {
			point.value++
			if (point.value > 9) {
				visited[point.key()] = true
				point.value = 0
				s.flashes++

				x := point.x
				y := point.y

				if (x > 0 && y > 0) {
					points = append(points, s.grid[x - 1][y - 1])
				}

				if (y > 0) {
					points = append(points, s.grid[x][y - 1])
				}

				if (x < len(s.grid) - 1 && y > 0) {
					points = append(points, s.grid[x + 1][y - 1])
				}

				if (x > 0) {
					points = append(points, s.grid[x - 1][y])
				}

				if (x < len(s.grid) - 1) {
					points = append(points, s.grid[x + 1][y])
				}

				if (x > 0 && y < len(s.grid[0]) - 1) {
					points = append(points, s.grid[x - 1][y + 1])
				}

				if (y < len(s.grid[0]) - 1) {
					points = append(points, s.grid[x][y + 1])
				}

				if (x < len(s.grid) - 1 && y < len(s.grid[0]) - 1) {
					points = append(points, s.grid[x + 1][y + 1])
				}
			}
		}
	}
	
}

func main() {
	// f, err := os.Open("11/sample_input.txt")
	f, err := os.Open("11/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	sim := Simulation{
		grid: [][]*Point{},
		flashes: 0,
	}

	for i, line := range lines {
		row := []*Point{}
		for j, char := range line {
			row = append(row, &Point{
				x: i,
				y: j,
				value: int(char) - ASCII_ZERO,
			})
		}
		sim.grid = append(sim.grid, row)
	}

	sim.printGrid()
	sim.SimulateDays(100)
	sim.printGrid()

	fmt.Printf("flashes: %v\n", sim.flashes)
}