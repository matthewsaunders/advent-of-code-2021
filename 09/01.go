package main

import (
	"bufio"
	"fmt"
	"os"
)

const ASCII_ZERO = 48

type Point struct {
	height int
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

	grid := [][]int{}

	for _, line := range lines {
		values := []int{}
		for _, x := range line {
			values = append(values, int(x) - ASCII_ZERO)
		}
		grid = append(grid, values)
		fmt.Println(values)
	}

	lowPoints := []*Point{}

	fmt.Printf("i -- %v -> %v\n", 0, len(grid))
	fmt.Printf("j -- %v -> %v\n", 0, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			height := grid[i][j]
			isLowpoint := true
			// fmt.Printf("[%v, %v]: %v\n", i, j, point)

			if (i != 0 && grid[i-1][j] <= height) {
				isLowpoint = false
			}

			if (i != len(grid) - 1 && grid[i+1][j] <= height) {
				isLowpoint = false
			}

			if (j != 0 && grid[i][j-1] <= height) {
				isLowpoint = false
			}

			if (j != len(grid[i]) - 1 && grid[i][j+1] <= height) {
				isLowpoint = false
			}

			if (isLowpoint) {
				fmt.Printf("==> low point [%v, %v] = %v\n", i, j, height)
				lowPoints = append(lowPoints, &Point{
					height: height,
				})
			}
		}
	}

	risk := 0
	fmt.Printf("# low points: %v\n", len(lowPoints))
	for _, point := range lowPoints {
		risk += point.height + 1
	}

	fmt.Printf("risk: %v\n", risk)
}