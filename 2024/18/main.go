package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input, 71, 1024))
	fmt.Printf("Part Two: %v\n", PartTwo(input, 71))
}

func MakeGrid(size int) Grid {
	grid := make([][]byte, size)
	for i := range size {
		grid[i] = make([]byte, size)
		for j := range size {
			grid[i][j] = '.'
		}
	}
	return Grid(grid)
}

func (g Grid) FindShortestPath() int {
	start, end := Vec{0, 0}, Vec{len(g) - 1, len(g) - 1}

	costToVisit := make(map[Vec]int)
	previousPoint := make(map[Vec]Vec)

	var Visit func(previous Vec, current Vec)
	Visit = func(previous Vec, current Vec) {
		if g.GetVec(current) != '.' {
			return
		}

		if _, ok := costToVisit[current]; !ok {
			costToVisit[current] = math.MaxInt
		}

		if costToVisit[previous]+1 < costToVisit[current] {
			costToVisit[current] = costToVisit[previous] + 1
			previousPoint[current] = previous

			for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				Visit(current, current.Add(dir))
			}
		}

		if current == end {
			return
		}

	}

	previousPoint[(Vec{1, 0})] = start
	Visit(start, Vec{1, 0})
	previousPoint[(Vec{0, 1})] = start
	Visit(start, Vec{0, 1})

	return costToVisit[end]
}

func (g Grid) HasPath() bool {
	start, end := Vec{0, 0}, Vec{len(g) - 1, len(g) - 1}
	visited := Set[Vec]{}
	foundEnd := false

	var Visit func(point Vec)
	Visit = func(point Vec) {
		visited.Push(point)

		if point == end {
			foundEnd = true
			return
		}

		if g.GetVec(point) != '.' {
			return
		}

		for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			if !visited.Contains(point.Add(dir)) {
				Visit(point.Add(dir))
			}
		}

	}

	Visit(start)
	return foundEnd
}

func SimulateGrid(input []string, size int, step int) Grid {
	grid := MakeGrid(size)

	for _, step := range input[:step] {
		parts := strings.Split(step, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		grid.Set(x, y, '#')
	}

	return grid
}

func PartOne(input []string, size int, steps int) int {
	grid := MakeGrid(size)

	for _, step := range input[:steps] {
		parts := strings.Split(step, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		grid.Set(x, y, '#')
	}

	return grid.FindShortestPath()
}

func PartTwo(input []string, size int) string {
	a, b := 0, len(input)
	step := (a + b) / 2

	for step > 0 && step < len(input) {
		pathAtStep := SimulateGrid(input, size, step).HasPath()
		if pathAtStep && !SimulateGrid(input, size, step+1).HasPath() {
			return input[step]
		}

		if pathAtStep {
			a = (a + b) / 2
		} else {
			b = (a + b) / 2
		}
		step = (a + b) / 2
	}

	return "-1,-1"
}
