package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func SplitNumber(val int) (int, int) {
	str := strconv.Itoa(val)
	a, _ := strconv.Atoi(str[0 : len(str)/2])
	b, _ := strconv.Atoi(str[len(str)/2:])

	return a, b
}

type Stones map[int]int

func (s Stones) Count() int {
	count := 0
	for _, v := range s {
		count += v
	}
	return count
}

func ParseStones(input string) Stones {
	stones := make(map[int]int)
	for _, val := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(val)
		stones[num] = 1
	}
	return stones
}

func Blink(stones map[int]int) Stones {
	output := make(map[int]int)

	for k, v := range stones {
		if k == 0 {
			output[1] += v
		} else if len(strconv.Itoa(k))%2 == 0 {
			a, b := SplitNumber(k)
			output[a] += v
			output[b] += v
		} else {
			output[k*2024] += v
		}
	}

	return output
}

func PartOne(input []string) int {
	stones := ParseStones(input[0])

	for range 25 {
		stones = Blink(stones)
	}

	return stones.Count()
}

func PartTwo(input []string) int {
	stones := ParseStones(input[0])

	for range 75 {
		stones = Blink(stones)
	}

	return stones.Count()
}
