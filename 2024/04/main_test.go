package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, "\n")
	expected := 18

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := strings.Split(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, "\n")
	expected := 9

	output := PartTwo(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
