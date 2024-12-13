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
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

type Vec struct {
	x float64
	y float64
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.x + v2.x, v.y + v2.y}
}

type Machine struct {
	movA Vec
	movB Vec
	goal Vec
}

func ParseMachines(input []string) []Machine {
	var machines []Machine

	for i := range len(input) {
		line := input[i]

		var movA Vec
		var movB Vec
		var goal Vec

		if strings.HasPrefix(line, "Button A: ") {
			output := strings.Split(strings.TrimPrefix(line, "Button A: "), ", ")
			x, _ := strconv.Atoi(strings.TrimPrefix(output[0], "X+"))
			y, _ := strconv.Atoi(strings.TrimPrefix(output[1], "Y+"))
			movA = Vec{x: float64(x), y: float64(y)}

			line = input[i+1]

			if strings.HasPrefix(line, "Button B: ") {
				output := strings.Split(strings.TrimPrefix(line, "Button B: "), ", ")
				x, _ := strconv.Atoi(strings.TrimPrefix(output[0], "X+"))
				y, _ := strconv.Atoi(strings.TrimPrefix(output[1], "Y+"))
				movB = Vec{x: float64(x), y: float64(y)}
			}

			line = input[i+2]

			if strings.HasPrefix(line, "Prize: ") {
				output := strings.Split(strings.TrimPrefix(line, "Prize: "), ", ")
				x, _ := strconv.Atoi(strings.TrimPrefix(output[0], "X="))
				y, _ := strconv.Atoi(strings.TrimPrefix(output[1], "Y="))
				goal = Vec{x: float64(x), y: float64(y)}
			}

			machines = append(machines, Machine{movA, movB, goal})

			i += 2
		}
	}

	return machines
}

func Calculate(movA Vec, movB Vec, goal Vec) (float64, float64) {
	a, b, c, d := -movA.y/movB.y, -movA.x/movB.x, goal.y/movB.y, goal.x/movB.x

	Px := ((d - c) / (a - b))
	Py := (a*(d-c)/(a-b) + c)

	Px = math.Floor(math.Round(Px*100)) / 100
	Py = math.Floor(math.Round(Py*100)) / 100

	if Px != math.Trunc(Px) {
		Px = -1
	}
	if Py != math.Trunc(Py) {
		Py = -1
	}

	return Px, Py
}

func PartOne(input []string) int {
	var cost int
	machines := ParseMachines(input)

	for _, machine := range machines {
		A, B := Calculate(machine.movA, machine.movB, machine.goal)
		if A >= 0 && A <= 100 && B >= 0 && B <= 100 {
			cost += int(A)*3 + int(B)
		}
	}

	return cost
}

func PartTwo(input []string) int {
	var cost int
	machines := ParseMachines(input)

	for _, machine := range machines {
		A, B := Calculate(machine.movA, machine.movB, machine.goal.Add(Vec{10000000000000, 10000000000000}))
		if A >= 0 && B >= 0 {
			cost += int(A)*3 + int(B)
		}
	}

	return cost
}
