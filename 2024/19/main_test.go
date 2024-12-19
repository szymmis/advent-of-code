package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`, "\n")
	expected := 6

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := strings.Split(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`, "\n")
	expected := 16 

	output := PartTwo(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
