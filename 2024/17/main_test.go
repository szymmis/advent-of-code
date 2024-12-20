package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`, "\n")
	expected := "4,6,3,5,6,3,5,2,1,0"

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := strings.Split(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`, "\n")
	expected := 117440

	output := PartTwo(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
