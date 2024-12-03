package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) int {
	sum := 0

	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
		}
	}

	return sum
}

func PartTwo(input []string) int {
	sum := 0
	canMultiplicate := true

	re, _ := regexp.Compile(`(mul|do|don't)\(((\d{1,3}),(\d{1,3}))?\)`)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			command := match[1]

			switch command {
			case "mul":
				if canMultiplicate {
					a, _ := strconv.Atoi(match[3])
					b, _ := strconv.Atoi(match[4])
					sum += a * b
				}
			case "do":
				canMultiplicate = true
			case "don't":
				canMultiplicate = false
			}
		}
	}

	return sum
}
