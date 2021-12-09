package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const ASCII_ZERO = 48

type Point struct {
	height int
	x int
	y int
}

func (p *Point) key() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

type Grid struct {
	grid [][]int
}

func (g *Grid) createPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
		height: g.grid[x][y],
	}
}

func calculateBasinCount(visited map[string]bool, g *Grid, p *Point, height int) int {
	if (p.height == 9 || p.height <= height) { return 0 }
	if _, found := visited[p.key()]; found { return 0 }

	fmt.Printf(" --[%v, %v]: %v\n", p.x, p.y, p.height)
	visited[p.key()] = true
	count := 1 // include this point in count

	if (p.x != 0) {
		count += calculateBasinCount(visited, g, g.createPoint(p.x - 1, p.y), p.height)
	}

	if (p.x != len(g.grid) - 1) {
		count += calculateBasinCount(visited, g, g.createPoint(p.x + 1, p.y), p.height)
	}

	if (p.y != 0) {
		count += calculateBasinCount(visited, g, g.createPoint(p.x, p.y - 1), p.height)
	}

	if (p.y != len(g.grid[p.x]) - 1) {
		count += calculateBasinCount(visited, g, g.createPoint(p.x, p.y + 1), p.height)
	}

	return count
}

func main() {
	// f, err := os.Open("09/sample_input.txt")
	f, err := os.Open("09/input.txt")
	if err != nil { panic("Error opening input file") }
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	grid := Grid{
		grid: [][]int{},
	}

	for _, line := range lines {
		values := []int{}
		for _, x := range line {
			values = append(values, int(x) - ASCII_ZERO)
		}
		grid.grid = append(grid.grid, values)
		fmt.Println(values)
	}

	lowPoints := []*Point{}

	fmt.Printf("i -- %v -> %v\n", 0, len(grid.grid))
	fmt.Printf("j -- %v -> %v\n", 0, len(grid.grid[0]))
	for i := 0; i < len(grid.grid); i++ {
		for j := 0; j < len(grid.grid[i]); j++ {
			height := grid.grid[i][j]
			isLowpoint := true
			// fmt.Printf("[%v, %v]: %v\n", i, j, point)

			if (i != 0 && grid.grid[i-1][j] <= height) {
				isLowpoint = false
			}

			if (i != len(grid.grid) - 1 && grid.grid[i+1][j] <= height) {
				isLowpoint = false
			}

			if (j != 0 && grid.grid[i][j-1] <= height) {
				isLowpoint = false
			}

			if (j != len(grid.grid[i]) - 1 && grid.grid[i][j+1] <= height) {
				isLowpoint = false
			}

			if (isLowpoint) {
				fmt.Printf("==> low point [%v, %v] = %v\n", i, j, height)
				lowPoints = append(lowPoints, &Point{
					x: i,
					y: j,
					height: height,
				})
			}
		}
	}

	basins := []int{}

	for i := 0; i < len(lowPoints); i++ {
		lowPoint := lowPoints[i]
		visitedPoints := make(map[string]bool)
		basins = append(basins, calculateBasinCount(visitedPoints, &grid, lowPoint, -1))
		fmt.Printf("basin %v: %v -- [%v, %v]\n", i, basins[i], lowPoint.x, lowPoint.y)
	}

	sort.Ints(basins)
	end := len(basins)
	fmt.Printf("1. %v\n", basins[end - 1])
	fmt.Printf("2. %v\n", basins[end - 2])
	fmt.Printf("3. %v\n", basins[end - 3])
	fmt.Printf("answer: %v\n", basins[end - 1] * basins[end - 2] * basins[end - 3])
}