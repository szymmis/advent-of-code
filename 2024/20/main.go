package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input, 100))
	fmt.Printf("Part Two: %v\n", PartTwo(input, 100))
}

func ParseInput(input []string) (Grid, Vec, Vec) {
	var start, end Vec
	grid := Grid(make([][]byte, len(input)))
	for y, line := range input {
		grid[y] = []byte(line)
	}

	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid.Get(x, y) == 'S' {
				start = Vec{x, y}
				grid.Set(x, y, '.')
			} else if grid.Get(x, y) == 'E' {
				end = Vec{x, y}
				grid.Set(x, y, '.')
			}
		}
	}

	return grid, start, end
}

func FindShortestPath(grid Grid, start Vec, end Vec) (map[Vec]int, []Vec) {
	var queue = []Vec{start}
	var visited = make(map[Vec]bool)
	var distance = map[Vec]int{
		start: 0,
	}
	var previous = make(map[Vec]Vec)

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		visited[point] = true

		if point == end {
			var path []Vec
			for {
				path = append(path, point)
				if p, ok := previous[point]; ok {
					point = p
				} else {
					break
				}
			}
			return distance, path
		}

		x, y := point.x, point.y
		for _, next := range []Vec{{x - 1, y}, {x, y - 1}, {x + 1, y}, {x, y + 1}} {
			if grid.GetVec(next) == '.' && !visited[next] {
				queue = append(queue, next)
				dist := distance[point] + 1
				if currDist, ok := distance[next]; !ok || dist < currDist {
					distance[next] = dist
					previous[next] = point
				}
			}
		}

	}

	panic("Cannot find a path")
}

type Cheat struct {
	start Vec
	end   Vec
}

func Abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func CheckCheat(distances map[Vec]int, point Vec, next Vec, length int) int {
	if cost, ok := distances[next]; ok {
		if cost > distances[point]+length {
			return cost - (distances[point] + length)
		}
	}

	return -1
}

func GetCheats(distances map[Vec]int, path []Vec, radius int, min int) int {
	var cheats = make(map[Cheat]bool)

	for _, point := range path {
		startX, startY := point.x, point.y
		for y := startY - radius; y <= startY+radius; y++ {
			for x := startX - radius; x <= startX+radius; x++ {
				dist := Abs(x-startX) + Abs(y-startY)
				if dist > 0 && dist <= radius {
					next := Vec{x, y}
					savedDist := CheckCheat(distances, point, next, dist)
					if savedDist >= min {
						cheats[Cheat{point, next}] = true
					}
				}
			}
		}
	}

	return len(cheats)
}

func PartOne(input []string, min int) int {
	grid, start, end := ParseInput(input)
	distances, path := FindShortestPath(grid, start, end)

	return GetCheats(distances, path, 2, min)
}

func PartTwo(input []string, min int) int {
	grid, start, end := ParseInput(input)
	distances, path := FindShortestPath(grid, start, end)

	return GetCheats(distances, path, 20, min)
}
