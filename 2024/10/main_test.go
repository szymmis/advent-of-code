package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")
	expected := 36

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")
	expected := 81

	output := PartTwo(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
