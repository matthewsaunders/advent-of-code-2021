package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type System struct {
	Games []*Game
	Moves []string
	CurrentMove int
	WinningGame *Game
}

func (s *System) Init(inputs []string) {
	s.CurrentMove = 0
	s.Moves = strings.Split(inputs[0], ",")

	// Built in assumption that the board is 5 rows
	for i := 2; i < len(inputs); i += 6 {
		newGame := new(Game)
		newGame.Init(inputs[i:i+5])
		s.Games = append(s.Games, newGame)
	}
}

func (s *System) PlayGames() {
	out:
	for s.CurrentMove = 0; s.CurrentMove < len(s.Moves); s.CurrentMove++ {
		for j := 0; j < len(s.Games); j++ {
			game := s.Games[j]
			game.MarkPosition(s.Moves[s.CurrentMove])

			if (!game.won && game.Won()) {
				game.won = true
				s.WinningGame = game
			}
		}

		if (s.AllGamesWon()) { break out }
	}
}

func (s *System) AllGamesWon() bool {
	allWon := true

	checkAllWon:
	for i := 0; i < len(s.Games); i++ {
		allWon = s.Games[i].won
		if (!allWon) { break checkAllWon }
	}

	return allWon
}

func (s *System) PrintFinalScore() {
	finalMove, err := strconv.Atoi(s.Moves[s.CurrentMove])
	if err != nil {
		panic("Error converting final move into a number")
	}

	fmt.Println("-- Winning game --")
	fmt.Printf("Final number: %v\n", finalMove)
	fmt.Printf("Game score: %v\n", s.WinningGame.Score())
	fmt.Printf("Final score: %v\n", finalMove * s.WinningGame.Score())
}

type Game struct {
	positions [][]*Position
	won bool
}

func (g *Game) Init(rows []string) {
	g.positions = [][]*Position{}

	for i := 0; i < len(rows); i++ {
		cols := strings.Fields(rows[i])
		row := []*Position{}

		for j := 0; j < len(cols); j++ {
			position := Position{
				x: j,
				y: i,
				marked: false,
				value: cols[j],
			}
			row = append(row, &position)
		}

		g.positions = append(g.positions, row)
	}
}

func (g *Game) PrintBoard() {
	fmt.Println("Printing board...")
	for i := 0; i < len(g.positions); i++ {
		fmt.Printf(" ")
		for j := 0; j < len(g.positions[i]); j++ {
			fmt.Printf("%v ", g.positions[i][j])
		}
		fmt.Printf("\n")
	}
}

func (g *Game) MarkPosition(value string) {
	out:
	for i := 0; i < len(g.positions); i++ {
		for j := 0; j < len(g.positions[i]); j++ {
			if (g.positions[i][j].value == value) {
				g.positions[i][j].marked = true
				break out
			}
		}
	}
}

func (g *Game) Won() bool {
	isWon := false

	allRowChecks:
	for i := 0; i < len(g.positions); i++ {
		if (isWon) { break allRowChecks }

		allMarked := true
		rowCheck:
		for j := 0; j < len(g.positions[i]); j++ {
			allMarked = g.positions[i][j].marked
			if (!allMarked) { break rowCheck }
		}
		
		if (allMarked) { isWon = true }
	}

	numCols := len(g.positions[0])

	allColChecks:
	for j := 0; j < numCols; j++ {
		if (isWon) { break allColChecks }

		allMarked := true
		colCheck:
		for i := 0; i < len(g.positions); i++ {
			allMarked = g.positions[i][j].marked
			if (!allMarked) { break colCheck }
		}

		if (allMarked) { isWon = true }
	}

	return isWon
}

func (g *Game) Score() int {
	score := 0

	for i := 0; i < len(g.positions); i++ {
		for j := 0; j < len(g.positions[i]); j++ {
			if (!g.positions[i][j].marked) {
				numVal, err := strconv.Atoi(g.positions[i][j].value)
				if err != nil {
					panic("Error parsing position value in score")
				}
				score += numVal
			}
		}
	}

	return score
}

type Position struct {
	x int
	y int
	marked bool
	value string
}

func main() {
	f, err := os.Open("04/input.txt");
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

	system := new(System)
	system.Init(lines)
	system.PlayGames()
	system.PrintFinalScore()
}
