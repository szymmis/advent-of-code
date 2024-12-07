package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func ConcatInts(a int, b int) int {
	str := strconv.Itoa(a) + strconv.Itoa(b)
	result, _ := strconv.Atoi(str)
	return result
}

func Calculate(elements []int) []int {
	if len(elements) == 1 {
		return elements
	}

	x := elements[0]
	y := elements[1]
	rest := elements[2:]

	return slices.Concat(
		Calculate(slices.Concat([]int{x + y}, rest)),
		Calculate(slices.Concat([]int{x * y}, rest)),
	)
}

func CalculateWithConcat(elements []int) []int {
	if len(elements) == 1 {
		return elements
	}

	x := elements[0]
	y := elements[1]
	rest := elements[2:]

	return slices.Concat(
		CalculateWithConcat(slices.Concat([]int{x + y}, rest)),
		CalculateWithConcat(slices.Concat([]int{x * y}, rest)),
		CalculateWithConcat(slices.Concat([]int{ConcatInts(x, y)}, rest)),
	)
}

type Equation struct {
	result   int
	elements []int
}

func ParseInput(input []string) []Equation {
	equations := make([]Equation, 0)

	for _, line := range input {
		split := strings.Split(line, ": ")
		result, _ := strconv.Atoi(split[0])
		values := strings.Split(strings.TrimSpace(split[1]), " ")
		elements := make([]int, 0)
		for _, val := range values {
			val, _ := strconv.Atoi(val)
			elements = append(elements, val)
		}
		equations = append(equations, Equation{result, elements})
	}

	return equations
}

func PartOne(input []string) int {
	count := 0
	equations := ParseInput(input)

	for _, equation := range equations {
		if slices.Contains(Calculate(equation.elements), equation.result) {
			count += equation.result
		}
	}

	return count
}

func PartTwo(input []string) int {
	count := 0
	equations := ParseInput(input)

	for _, equation := range equations {
		if slices.Contains(CalculateWithConcat(equation.elements), equation.result) {
			count += equation.result
		}
	}

	return count
}
