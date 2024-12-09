package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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
	if !g.InBounds(x, y) {
		return 0
	}

	return g[y][x]
}

func (g Grid) InBounds(x int, y int) bool {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return false
	}

	return true
}

func (g *Grid) Set(x int, y int, char rune) {
	(*g)[y][x] = char
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

func (g Grid) Simulate(guard *Guard) bool {
	// start := Guard{guard.x, guard.y, guard.direction}
	occurences := make(map[string]int)

	for guard.InBounds(g) {
		str := fmt.Sprintf("%d,%d,%d,%d", guard.x, guard.y, guard.direction.x, guard.direction.y)
		if guard.InBounds(g) {
			if occurences[str] += 1; occurences[str] > 1 {
				return true
			}
		}
		guard.Move(g)
		// if guard.x == start.x && guard.y == start.y && guard.direction.x == start.direction.x && guard.direction.y == start.direction.y {
		// return true
		// }
	}

	return false
}

func (g Grid) SimulateWithObstacle(guard Guard, obstacle Vec) bool {
	if g.Get(obstacle.x, obstacle.y) == '#' || !g.InBounds(obstacle.x, obstacle.y) {
		return false
	}

	g.Set(obstacle.x, obstacle.y, '#')
	loop := g.Simulate(&guard)
	g.Set(obstacle.x, obstacle.y, '.')
	return loop
}

func (g *Grid) Clear() {
	for y := range len(*g) {
		for x := range len((*g)[y]) {
			if (*g)[y][x] != '#' {
				(*g)[y][x] = '.'
			}
		}
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

func (v Vec) AsString() string {
	return fmt.Sprintf("%v,%v", v.x, v.y)
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

	grid.Simulate(&guard)

	return grid.CountVisitedFields()
}

func PartTwo(input []string) int {
	grid, guard := ParseGrid(input)
	x, y, dir := guard.x, guard.y, guard.direction

	var states []Guard

	for guard.InBounds(grid) {
		x, y, dir := guard.x, guard.y, guard.direction
		if !guard.Move(grid) {
			states = append(states, Guard{x, y, dir})
		}
	}

	var positions []string

	for _, state := range states {
		obstacle := Vec{state.x + state.direction.x, state.y + state.direction.y}
		if !slices.Contains(positions, obstacle.AsString()) {
			if grid.SimulateWithObstacle(Guard{x, y, dir}, obstacle) {
				positions = append(positions, obstacle.AsString())
			}
		}
	}

	return len(positions)
}
