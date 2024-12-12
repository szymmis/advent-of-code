package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := strings.Split(`125 17`, "\n")
	expected := 55312

	output := PartOne(input)

	if output != expected {
		t.Errorf("Output is %v, expeced %v", output, expected)
	}
}
