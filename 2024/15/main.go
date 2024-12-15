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

func ParseInput(input []string) (Warehouse, []Move) {
	var grid Grid
	var moves []Move
	var robot Vec

	for i := 0; i < len(input); i++ {
		for i < len(input) && len(strings.TrimSpace(input[i])) > 0 {
			line := input[i]
			if strings.Contains(line, "#") {
				grid = append(grid, []byte(line))
			} else if strings.Contains(line, "<") {
				moves = slices.Concat(moves, []Move(line))
			}
			i++
		}
	}

Outer:
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid.Get(x, y) == '@' {
				robot = Vec{x, y}
				break Outer
			}
		}
	}

	return Warehouse{grid, robot}, moves
}

type Move byte

func (i Move) AsVec() Vec {
	switch i {
	case '<':
		return Vec{-1, 0}
	case '>':
		return Vec{1, 0}
	case '^':
		return Vec{0, -1}
	case 'v':
		return Vec{0, 1}
	}

	panic(fmt.Sprintf("Invalid move %q", i))
}

type VisitedVectors []Vec

func (vecs *VisitedVectors) Contains(v Vec) bool {
	return slices.ContainsFunc(*vecs, func(v2 Vec) bool {
		return v.x == v2.x && v.y == v2.y
	})
}

func (vecs *VisitedVectors) Push(v Vec) {
	if !vecs.Contains(v) {
		*vecs = append(*vecs, v)
	}
}

type Warehouse struct {
	grid  Grid
	robot Vec
}

func (w Warehouse) CanPush(pos Vec, dir Vec) bool {
	var visited VisitedVectors

	var Check func(p Vec) bool
	Check = func(p Vec) bool {
		object := w.grid.GetVec(p)

		if visited.Contains(p) {
			return true
		} else {
			visited.Push(p)
		}

		switch object {
		case '.':
			return true
		case '#':
			return false
		case '[':
			return Check(p.Add(dir)) && Check(p.Add(Vec{1, 0}))
		case ']':
			return Check(p.Add(dir)) && Check(p.Add(Vec{-1, 0}))
		default:
			return Check(p.Add((dir)))
		}
	}

	return Check(pos)
}

func (w *Warehouse) Push(pos Vec, dir Vec) bool {
	var visited VisitedVectors

	var PushMany func(points []Vec) bool
	PushMany = func(points []Vec) bool {
		for _, p := range points {
			if !w.CanPush(p, dir) {
				return false
			}
		}

		for _, p := range points {
			object := w.grid.GetVec(p)
			if object == '.' || object == '#' {
				return true
			}

			if visited.Contains(p) {
				continue
			} else {
				visited.Push(p)
			}

			var toVisit []Vec

			if object == '@' || object == 'O' {
				toVisit = []Vec{p.Add(dir)}
			} else if object == '[' {
				toVisit = []Vec{p.Add(Vec{1, 0}), p.Add(dir)}
			} else if object == ']' {
				toVisit = []Vec{p.Add(Vec{-1, 0}), p.Add(dir)}
			}

			PushMany(toVisit)
			w.grid.Set(p.x+dir.x, p.y+dir.y, object)
			w.grid.Set(p.x, p.y, '.')
		}

		return true
	}

	return PushMany([]Vec{pos})
}

func (w *Warehouse) MoveRobot(direction Vec) {
	if w.Push(w.robot, direction) {
		w.robot.x += direction.x
		w.robot.y += direction.y
	}
}

func (w Warehouse) Score() int {
	gps := 0
	for y := range len(w.grid) {
		for x := range len(w.grid[y]) {
			if w.grid.Get(x, y) == 'O' || w.grid.Get(x, y) == '[' {
				gps += 100*y + x
			}
		}
	}
	return gps
}

func (w *Warehouse) Expand() {
	expandedGrid := make([][]byte, len(w.grid))
	for y := range len(w.grid) {
		for x := range len(w.grid[y]) {
			var tiles []byte

			switch w.grid[y][x] {
			case '#':
				tiles = []byte{'#', '#'}
			case 'O':
				tiles = []byte{'[', ']'}
			case '.':
				tiles = []byte{'.', '.'}
			case '@':
				tiles = []byte{'@', '.'}
				w.robot.x = x * 2
			}

			expandedGrid[y] = slices.Concat(expandedGrid[y], tiles)
		}
	}
	w.grid = expandedGrid
}

func PartOne(input []string) int {
	warehouse, moves := ParseInput(input)

	for _, move := range moves {
		warehouse.MoveRobot(move.AsVec())
	}

	return warehouse.Score()
}

func PartTwo(input []string) int {
	warehouse, moves := ParseInput(input)
	warehouse.Expand()

	for _, move := range moves[:] {
		warehouse.MoveRobot(move.AsVec())
	}

	return warehouse.Score()
}
