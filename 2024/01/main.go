package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func ParseAndSortBothArrays(input string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, "   ")
		first, _ := strconv.Atoi(nums[0])
		left = append(left, first)
		second, _ := strconv.Atoi(nums[1])
		right = append(right, second)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func PartOne(input string) int {
	sum := 0
	left, right := ParseAndSortBothArrays(input)

	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		sum += distance
	}

	return sum
}

func PartTwo(input string) int {
	sum := 0
	left, right := ParseAndSortBothArrays(input)
	valuesCount := make(map[int]int)

	for _, val := range right {
		valuesCount[val]++
	}

	for _, number := range left {
		sum += number * valuesCount[number]
	}

	return sum
}
