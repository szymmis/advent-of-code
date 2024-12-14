package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input, 101, 103))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

type Vec struct {
	x int
	y int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.x + v2.x, v.y + v2.y}
}

type Grid [][]rune

func (g Grid) InBounds(x int, y int) bool {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return false
	}

	return true
}

func (g Grid) Get(x int, y int) rune {
	if !g.InBounds(x, y) {
		return 0
	}

	return g[y][x]
}

func (g Grid) String() string {
	var output string

	for y := range len(g) {
		for x := range len(g[y]) {
			if g[y][x] == 'X' {
				output += string(g[y][x])
			} else {
				output += "."
			}
		}
		output += "\n"
	}

	return output
}

func RobotsAsGrid(robots []*Robot) Grid {
	grid := Grid(make([][]rune, 103))
	for line := range 103 {
		grid[line] = make([]rune, 101)
	}

	for _, robot := range robots {
		grid[robot.pos.y][robot.pos.x] = 'X'
	}
	return grid
}

type Robot struct {
	pos      Vec
	velocity Vec
}

func ParseRobots(input []string) []*Robot {
	var robots []*Robot
	for _, line := range input {
		parts := strings.Split(line, " ")
		pos := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		vel := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])
		robots = append(robots, &Robot{pos: Vec{x, y}, velocity: Vec{vx, vy}})
	}
	return robots
}

func (r *Robot) Simulate(width int, height int) {
	r.pos = Vec{x: (r.pos.x + r.velocity.x + width) % width, y: (r.pos.y + r.velocity.y + height) % height}
}

func (r Robot) String() string {
	return fmt.Sprintf("{%v %v}", r.pos, r.velocity)
}

func PartOne(input []string, width int, height int) int {
	robots := ParseRobots(input)

	for range 100 {
		for _, robot := range robots {
			robot.Simulate(width, height)
		}
	}

	var q1, q2, q3, q4 int

	for _, robot := range robots {
		if robot.pos.x < width/2 {
			if robot.pos.y < height/2 {
				q1++
			} else if robot.pos.y > height/2 {
				q2++
			}
		} else if robot.pos.x > width/2 {
			if robot.pos.y < height/2 {
				q3++
			} else if robot.pos.y > height/2 {
				q4++
			}
		}
	}

	return q1 * q2 * q3 * q4
}

func PartTwo(input []string) int {
	robots := ParseRobots(input)

	for s := 1; s <= 10000; s++ {
		for _, robot := range robots {
			robot.Simulate(101, 103)
		}

		grid := RobotsAsGrid(robots)

		const GridSize = 10
		var sum int

		for y := 0; y < len(grid); y += GridSize {
			for x := 0; x < len(grid[y]); x += GridSize {
				var occupied int
				for i := range GridSize {
					for j := range GridSize {
						if grid.Get(x+j, y+i) == 'X' {
							occupied++
						}
					}
				}
				if occupied >= GridSize*GridSize/2 {
					sum++
				}
			}
		}

		if sum > 0 {
			fmt.Printf("%v", grid)
			return s
		}
	}

	return -1
}
