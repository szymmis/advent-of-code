package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid []string

func (g Grid) Get(x int, y int) (byte, error) {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return 0, fmt.Errorf("indexing out of bounds for %d, %d", x, y)
	}

	return g[y][x], nil
}

func (g Grid) ContainsWord(word string, start Vec, direction Vec) bool {
	Search := func(index func(i int, word string) int) bool {
		for i := range len(word) {
			char, _ := g.Get(start.x+direction.x*i, start.y+direction.y*i)

			if word[index(i, word)] != char {
				break
			}

			if i == len(word)-1 {
				return true
			}
		}

		return false
	}

	return Search(func(i int, word string) int { return i }) || Search(func(i int, word string) int { return len(word) - 1 - i })
}

type Vec struct {
	x int
	y int
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) int {
	count := 0
	grid := Grid(input)
	for _, direction := range []Vec{{1, 0}, {0, 1}, {1, 1}, {1, -1}} {
		for y := range len(grid) {
			for x := range len(grid[y]) {
				if grid.ContainsWord("XMAS", Vec{x, y}, direction) {
					count++
				}
			}
		}
	}
	return count
}

func PartTwo(input []string) int {
	count := 0
	grid := Grid(input)
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid.ContainsWord("MAS", Vec{x - 1, y - 1}, Vec{1, 1}) && grid.ContainsWord("MAS", Vec{x + 1, y - 1}, Vec{-1, 1}) {
				count++
			}
		}
	}
	return count
}
