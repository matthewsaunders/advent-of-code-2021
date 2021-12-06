package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	positions map[string]int
}

func (g *Grid) Init(lines []string) {
	g.positions = make(map[string]int)
	g.drawLines(lines)
}

func (g *Grid) drawLines(lines []string) {
	for i := 0; i < len(lines); i++ {
		fmt.Printf("Line: %v\n", lines[i])
		boundaries := strings.Split(lines[i], " -> ")
		start := strings.Split(boundaries[0], ",")
		stop := strings.Split(boundaries[1], ",")

		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(stop[0])
		y2, _ := strconv.Atoi(stop[1])

		// row
		if (x1 == x2) {
			lower, upper := y2, y1
			if (y1 < y2) {
				lower, upper = y1, y2
			}

			for i := lower; i <= upper; i++ {
				g.positions[g.key(x1, i)]++
			}
		}

		// col
		if (y1 == y2) {
			lower, upper := x2, x1
			if (x1 < x2) {
				lower, upper = x1, x2
			}

			for i := lower; i <= upper; i++ {
				g.positions[g.key(i, y1)]++
			}
		}
	}
}

func (g *Grid) key(x int, y int) string {
	return fmt.Sprintf("%v:%v", x, y)
}

func (g *Grid) numOverlappingPoints() int {
	numOverlaps := 0

	fmt.Println("Positions:")
	for k, v := range g.positions {
		fmt.Printf(" ->%v -- %v\n", k, v)
		if (v > 1) { numOverlaps++ }
	}

	return numOverlaps
}

func main() {
	f, err := os.Open("05/input.txt")
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

	grid := new(Grid)
	grid.Init(lines)
	fmt.Printf("# of overlapping points: %v\n", grid.numOverlappingPoints())
}