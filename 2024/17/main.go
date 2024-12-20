package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	a      int
	b      int
	c      int
	output []string
}

func (m *Machine) GetValue(operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return m.a
	case 5:
		return m.b
	case 6:
		return m.c
	}

	panic(fmt.Sprintf("invalid operand %v", operand))
}

func (m *Machine) Exec(program []string) {
	for i := 0; i < len(program); {
		opcode, _ := strconv.Atoi(program[i])
		operand, _ := strconv.Atoi(program[i+1])

		switch opcode {
		case 0:
			(*m).a = m.a / int(math.Pow(2.0, float64(m.GetValue(operand))))
		case 1:
			(*m).b = m.b ^ operand
		case 2:
			(*m).b = m.GetValue(operand) % 8
		case 3:
			if m.a != 0 {
				i = operand
				continue
			}
		case 4:
			(*m).b = m.b ^ m.c
		case 5:
			(*m).output = append((*m).output, strconv.Itoa(m.GetValue(operand)%8))
		case 6:
			(*m).b = m.a / int(math.Pow(2.0, float64(m.GetValue(operand))))
		case 7:
			(*m).c = m.a / int(math.Pow(2.0, float64(m.GetValue(operand))))
		}

		i += 2
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) string {
	a, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	program := strings.Split(strings.Split(input[4], ": ")[1], ",")

	machine := Machine{a, 0, 0, []string{}}
	machine.Exec(program)

	return strings.Join(machine.output[:], ",")
}

func PartTwo(input []string) int {
	program := strings.Split(strings.Split(input[4], ": ")[1], ",")

	for a := 1; a < math.MaxInt; {
		machine := Machine{a, 0, 0, []string{}}
		machine.Exec(program)

		if slices.Equal(program[max(0, len(program)-len(machine.output)):], machine.output) {
			fmt.Printf("a = %15v | %v\n", a, machine.output)
			if len(program) == len(machine.output) {
				return a
			}
			a *= 8
		} else {
			a++
		}
	}

	return -1
}
