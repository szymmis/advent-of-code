package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`, "\n")
	expected := 44

	output := PartOne(input, 1)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := strings.Split(`###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`, "\n")
	expected := 285

	output := PartTwo(input, 50)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
