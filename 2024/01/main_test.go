package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31

	output := PartTwo(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
