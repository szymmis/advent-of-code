package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

type Vec struct {
	x int
	y int
}

type Grid []string

func ParseGrid(input []string) (Grid, []Vec) {
	var starts []Vec
	grid := Grid(input)

	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid[y][x] == '0' {
				starts = append(starts, Vec{x, y})
			}
		}
	}

	return grid, starts
}

func (g Grid) Get(x int, y int) byte {
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

func FindReachablePoints(grid Grid, point Vec, path []Vec) []Vec {
	current := grid.Get(point.x, point.y)
	if current == '9' {
		return []Vec{point}
	}

	var reachablePoints []Vec
	for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if grid.Get(point.x+dir.x, point.y+dir.y)-current == 1 {
			reachablePoints = slices.Concat(reachablePoints, FindReachablePoints(grid, Vec{point.x + dir.x, point.y + dir.y}, slices.Concat(path, []Vec{point})))
		}
	}
	return reachablePoints
}

func FindPathsCount(grid Grid, point Vec, path []Vec) int {
	current := grid.Get(point.x, point.y)
	if current == '9' {
		return 1
	}

	var trailsCount int
	for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if grid.Get(point.x+dir.x, point.y+dir.y)-current == 1 {
			trailsCount += FindPathsCount(grid, Vec{point.x + dir.x, point.y + dir.y}, slices.Concat(path, []Vec{point}))
		}
	}
	return trailsCount
}

func PartOne(input []string) int {
	sum := 0
	grid, starts := ParseGrid(input)

	for _, start := range starts {
		var points []Vec
		for _, point := range FindReachablePoints(grid, start, []Vec{}) {
			if !slices.ContainsFunc(points, func(p Vec) bool {
				return point.x == p.x && point.y == p.y
			}) {
				points = append(points, point)
			}
		}
		sum += len(points)
	}

	return sum
}

func PartTwo(input []string) int {
	sum := 0
	grid, starts := ParseGrid(input)

	for _, start := range starts {
		sum += FindPathsCount(grid, start, []Vec{})
	}

	return sum
}
