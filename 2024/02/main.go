package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func ParseLine(line string) []int {
	vals := strings.Split(line, " ")
	nums := make([]int, len(vals))
	for i, val := range vals {
		nums[i], _ = strconv.Atoi(val)
	}
	return nums
}

func GetSign(a int, b int) int {
	if diff := a - b; diff >= 0 {
		return 1
	} else {
		return -1
	}
}

func CheckSlice(s []int, excludeIndex int) bool {
	var prevSign bool

	for a, b := 0, 1; a < len(s) && b < len(s); a, b = a+1, b+1 {
		if a == excludeIndex {
			a++
			if a == b {
				b++
			}
		}
		if b == excludeIndex {
			b++
		}

		if a >= len(s) || b >= len(s) {
			break
		}

		if diff := int(math.Abs(float64(s[a] - s[b]))); diff < 1 || diff > 3 {
			return false
		}

		if sign := GetSign(s[a], s[b]); prevSign == 0 || prevSign == sign {
			prevSign = sign
		} else {
			return false
		}
	}

	return true
}

func PartOne(input string) int {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		nums := ParseLine(line)

		if CheckSlice(nums, -1) {
			count++
		}
	}

	return count
}

func PartTwo(input string) int {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		nums := ParseLine(line)

		for i := 0; i < len(nums); i++ {
			if CheckSlice(nums, i) {
				count++
				break
			}
		}
	}

	return count
}
