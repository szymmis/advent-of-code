package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
}

func CheckKey(key []int, lock []int) bool {
	for i := range lock {
		if key[i]+lock[i] > 5 {
			return false
		}
	}

	return true
}

func PartOne(input []string) int {
	var locks [][]int
	var keys [][]int

	for i := 0; i < len(input); i++ {
		line := input[i]

		if line == "#####" {
			i++
			lock := make([]int, 5)
			for y := 0; y < 5; y++ {
				line := input[i]
				for x := 0; x < 5; x++ {
					if line[x] == '#' {
						lock[x]++
					}
				}
				i++
			}
			locks = append(locks, lock)
		} else if line == "....." {
			i++
			key := make([]int, 5)
			for y := 0; y < 5; y++ {
				line := input[i]
				for x := 0; x < 5; x++ {
					if line[x] == '#' {
						key[x]++
					}
				}
				i++
			}
			keys = append(keys, key)
		}
	}

	var pairs int

	for _, key := range keys {
		for _, lock := range locks {
			if CheckKey(key, lock) {
				pairs++
			}
		}
	}

	return pairs
}
