package main

import (
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]int)

func CountPossibilities(pattern string, towels []string) int {
	if value, ok := memo[pattern]; ok {
		return value
	}

	var count int
	for _, towel := range towels {
		if pattern == towel {
			count++
		} else {
			if !strings.HasPrefix(pattern, towel) {
				continue
			}

			count += CountPossibilities(pattern[len(towel):], towels)
		}
	}

	memo[pattern] = count
	return count
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) int {
	var count int
	towels := strings.Split(input[0], ", ")

	for _, pattern := range input[2:] {
		if CountPossibilities(pattern, towels) > 0 {
			count++
		}
	}

	return count
}

func PartTwo(input []string) int {
	var sum int
	towels := strings.Split(input[0], ", ")

	for _, pattern := range input[2:] {
		sum += CountPossibilities(pattern, towels)
	}

	return sum
}
