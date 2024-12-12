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

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.x + v2.x, v.y + v2.y}
}

func (v Vec) Equals(v2 Vec) bool {
	return v.x == v2.x && v.y == v2.y
}

func (v Vec) AsString() string {
	return fmt.Sprintf("%d,%d", v.x, v.y)
}

type Grid []string

func (g Grid) InBounds(x int, y int) bool {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return false
	}

	return true
}

func (g Grid) Get(x int, y int) byte {
	if !g.InBounds(x, y) {
		return 0
	}

	return g[y][x]
}

func (g Grid) Getv(v Vec) byte {
	return g.Get(v.x, v.y)
}

type Patch struct {
	name   byte
	points []Vec
}

func ExtractPatches(grid Grid) []Patch {
	var visited []string

	var Visit func(grid Grid, point Vec) []Vec
	Visit = func(grid Grid, point Vec) []Vec {
		if !grid.InBounds(point.x, point.y) || slices.Contains(visited, point.AsString()) {
			return []Vec{}
		}

		patch := []Vec{point}

		visited = append(visited, point.AsString())
		for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			if grid.Getv(point.Add(dir)) == grid.Getv(point) {
				patch = slices.Concat(patch, Visit(grid, point.Add(dir)))
			}
		}

		return patch
	}

	var patches []Patch
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if !slices.Contains(visited, Vec{x, y}.AsString()) {
				patches = append(patches, Patch{name: grid.Getv(Vec{x, y}), points: Visit(grid, Vec{x, y})})
			}
		}
	}
	return patches
}

func (p Patch) Perimeter(grid Grid) int {
	perimeter := 0
	for _, point := range p.points {
		for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			if grid.Getv(point.Add(dir)) != grid.Getv(point) {
				perimeter += 1
			}
		}
	}
	return perimeter
}

func (p Patch) BoundaryPoints(grid Grid) []Vec {
	var points []Vec
	for _, point := range p.points {
		for _, dir := range []Vec{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			if grid.Getv(point.Add(dir)) != grid.Getv(point) {
				if !slices.ContainsFunc(points, func(p Vec) bool {
					return point.Equals(p)
				}) {
					points = append(points, point)
				}
			}
		}
	}
	return points
}

func (p Patch) ComplexPerimeter(grid Grid) int {
	var visited []string

	MatchSide := func(normal Vec, dir Vec, point Vec) []Vec {
		var side []Vec
		p := Vec{point.x, point.y}
		str := p.AsString() + "," + normal.AsString()
		if !slices.Contains(visited, str) {
			for grid.Getv(p) == grid.Getv(point) && grid.Getv(p.Add(normal)) != grid.Getv(point) {
				str := p.AsString() + "," + normal.AsString()
				if !slices.Contains(visited, str) {
					visited = append(visited, str)
				} else {
					return []Vec{}
				}
				side = append(side, p)
				p = p.Add(dir)
			}
		}
		return side
	}

	var count int

	for _, point := range p.BoundaryPoints(grid) {
		if len(MatchSide(Vec{0, -1}, Vec{1, 0}, point)) > 0 {
			count++
		}
		if len(MatchSide(Vec{0, 1}, Vec{1, 0}, point)) > 0 {
			count++
		}
		if len(MatchSide(Vec{-1, 0}, Vec{0, 1}, point)) > 0 {
			count++
		}
		if len(MatchSide(Vec{1, 0}, Vec{0, 1}, point)) > 0 {
			count++
		}
	}

	return count
}

func (p Patch) Area() int {
	return len(p.points)
}

func PartOne(input []string) int {
	grid := Grid(input)
	patches := ExtractPatches(grid)

	sum := 0
	for _, patch := range patches {
		sum += patch.Perimeter(grid) * patch.Area()
	}
	return sum
}

func PartTwo(input []string) int {
	grid := Grid(input)
	patches := ExtractPatches(grid)

	sum := 0
	for _, patch := range patches {
		sum += patch.ComplexPerimeter(grid) * patch.Area()
	}
	return sum
}
