package main

import "fmt"

type Vec struct {
	x int
	y int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.x + v2.x, v.y + v2.y}
}

type Grid [][]byte

func (g Grid) InBounds(x int, y int) bool {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return false
	}

	return true
}

func (g Grid) Get(x int, y int) byte {
	if !g.InBounds(x, y) {
		return 0
	}

	return g[y][x]
}

func (g Grid) GetVec(v Vec) byte {
	return g.Get(v.x, v.y)
}

func (g *Grid) Set(x int, y int, b byte) {
	(*g)[y][x] = b
}

func (g Grid) String() string {
	var output string

	for y := range len(g) {
		for x := range len(g[y]) {
			output += string(g[y][x])
		}
		if y < len(g)-1 {
			output += "\n"
		}
	}

	return output
}

func (g Grid) Print() {
	fmt.Printf("%v\n", g)
}
