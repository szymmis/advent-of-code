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

func (v Vec) AsString() string {
	return fmt.Sprintf("%v,%v", v.x, v.y)
}

func (v Vec) InBounds(grid []string) bool {
	return v.x >= 0 && v.y >= 0 && v.x < len(grid[0]) && v.y < len(grid)
}

type Antena struct {
	pos  Vec
	char rune
}

func ParseAntenas(input []string) map[rune][]Antena {
	antenas := make(map[rune][]Antena)

	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antenas[char] = append(antenas[char], Antena{Vec{x, y}, char})
			}
		}
	}

	return antenas
}

func PartOne(input []string) int {
	antenas := ParseAntenas(input)
	antinodes := make([]string, 0)

	for _, antenas := range antenas {
		for i := 0; i < len(antenas); i++ {
			for j := 0; j < len(antenas); j++ {
				if i != j {
					pos := antenas[j].pos.Add(Vec{antenas[j].pos.x - antenas[i].pos.x, antenas[j].pos.y - antenas[i].pos.y})
					if pos.InBounds(input) && !slices.Contains(antinodes, pos.AsString()) {
						antinodes = append(antinodes, pos.AsString())
					}
				}
			}
		}
	}

	return len(antinodes)
}

func PartTwo(input []string) int {
	antenas := ParseAntenas(input)
	antinodes := make([]string, 0)

	for _, antenas := range antenas {
		for i := 0; i < len(antenas); i++ {
			for j := 0; j < len(antenas); j++ {
				if i != j {
					vec := Vec{antenas[j].pos.x - antenas[i].pos.x, antenas[j].pos.y - antenas[i].pos.y}
					pos := antenas[i].pos
					for pos.InBounds(input) {
						if !slices.Contains(antinodes, pos.AsString()) {
							antinodes = append(antinodes, pos.AsString())
						}
						pos = pos.Add(vec)
					}
				}
			}
		}
	}

	return len(antinodes)
}
