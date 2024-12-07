package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Grid [][]rune

func ParseGrid(input []string) (Grid, Guard) {
	grid := Grid(make([][]rune, 0))
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	var guard Guard
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid.Get(x, y) == '^' {
				guard = Guard{x, y, Vec{0, -1}}
				grid.Visit(x, y)
			}
		}
	}

	return grid, guard
}

func (g Grid) Get(x int, y int) rune {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return 0
	}

	return g[y][x]
}

func (g Grid) Visit(x int, y int) {
	g[y][x] = 'X'
}

func (g Grid) Print() {
	for _, line := range g {
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
		fmt.Println()
	}
}

func (g Grid) GetVisitedFields() []Vec {
	fields := make([]Vec, 0)
	for y, line := range g {
		for x, char := range line {
			if char == 'X' {
				fields = append(fields, Vec{x, y})
			}
		}
	}
	return fields
}

func (g Grid) CountVisitedFields() int {
	count := 0
	for _, line := range g {
		for _, char := range line {
			if char == 'X' {
				count += 1
			}
		}
	}
	return count
}

type Vec struct {
	x int
	y int
}

type Guard struct {
	x         int
	y         int
	direction Vec
}

func (g *Guard) Move(grid Grid) bool {
	x := g.x + g.direction.x
	y := g.y + g.direction.y

	if grid.Get(x, y) == '#' {
		g.Rotate()

		return true
	} else {
		g.x += g.direction.x
		g.y += g.direction.y

		if g.InBounds(grid) {
			grid.Visit(x, y)
		}

		return false
	}
}

func (g *Guard) Rotate() {
	angle := math.Pi / 2
	x := int(float64(g.direction.x)*(math.Cos(angle)) - float64(g.direction.y)*(math.Sin(angle)))
	y := int(float64(g.direction.x)*(math.Sin(angle)) + float64(g.direction.y)*(math.Cos(angle)))
	g.direction = Vec{x, y}
}

func (g Guard) InBounds(grid Grid) bool {
	return g.x >= 0 && g.y >= 0 && g.x < len(grid[0]) && g.y < len(grid)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) int {
	grid, guard := ParseGrid(input)

	for guard.InBounds(grid) {
		guard.Move(grid)
	}

	return grid.CountVisitedFields()
}

func PartTwo(input []string) int {
	return -1
}
