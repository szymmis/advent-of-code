package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

type Maze struct {
	grid  Grid
	start Vec
	end   Vec
}

func ParseMaze(input []string) Maze {
	var start, end Vec
	grid := Grid(make([][]byte, len(input)))
	for y, line := range input {
		grid[y] = []byte(line)
		for x := range line {
			if grid.Get(x, y) == 'S' {
				start = Vec{x, y}
				grid.Set(x, y, '.')
			} else if grid.Get(x, y) == 'E' {
				end = Vec{x, y}
				grid.Set(x, y, '.')
			}
		}
	}
	return Maze{grid, start, end}
}

type Position struct {
	point Vec
	dir   Vec
}

func (m Maze) Traverse(startingDirection Vec) (Set[Vec], int) {
	var grid = m.grid
	var start = m.start
	var end = m.end

	var endPos Position
	var minLength = math.MaxInt

	var distances = make(map[Position]int)
	var previousVectors = make(map[Position][]Position)

	var Visit func(previous Vec, current Position)
	Visit = func(previous Vec, current Position) {
		if grid.GetVec(current.point) != '.' {
			return
		}

		previousVectors[current] = []Position{}
		directions := []Vec{current.dir, current.dir.Rotate(-math.Pi / 2), current.dir.Rotate(math.Pi / 2)}

		for _, dir := range directions {
			positionToCheck := Position{previous, dir}

			if _, ok := distances[positionToCheck]; !ok {
				continue
			}

			if _, ok := distances[current]; !ok {
				distances[current] = math.MaxInt
			}

			var cost int
			if current.dir.x*dir.x+current.dir.y*dir.y == 0 {
				cost = 1001
			} else {
				cost = 1
			}
			distance := distances[positionToCheck] + cost

			if distance > distances[current] {
				continue
			} else if distance < distances[current] {
				distances[current] = distance
				previousVectors[current] = []Position{positionToCheck}
			} else if distance == distances[current] {
				previousVectors[current] = append(previousVectors[current], positionToCheck)
			}
		}

		if current.point == end {
			if distances[current] < minLength {
				endPos = current
				minLength = distances[current]
			}
			return
		}

		for _, dir := range directions {
			newPos := Position{current.point.Add(dir), dir}
			if _, ok := distances[newPos]; !ok {
				distances[newPos] = math.MaxInt
			}
			if distances[current] < distances[newPos] {
				Visit(current.point, newPos)
			}
		}
	}

	for _, dir := range []Vec{startingDirection, startingDirection.Rotate(-math.Pi / 2), startingDirection.Rotate(math.Pi / 2)} {
		startingPosition := Position{start, dir}
		if dir.IsParallelTo(startingDirection) {
			distances[startingPosition] = 0
		} else {
			distances[startingPosition] = 1000
		}
		Visit(start, Position{start.Add(dir), dir})
	}

	var points Set[Vec]

	var Collect func(current Position)
	Collect = func(current Position) {
		points.Push(current.point)
		for _, pos := range previousVectors[current] {
			Collect(pos)
		}
	}

	Collect(endPos)
	return points, minLength
}

func PartOne(input []string) int {
	maze := ParseMaze(input)
	_, cost := maze.Traverse(Vec{1, 0})
	return cost
}

func PartTwo(input []string) int {
	maze := ParseMaze(input)
	points, _ := maze.Traverse(Vec{1, 0})
	return len(points)
}
