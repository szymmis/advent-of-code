package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	page   int
	before int
}
type Page []int
type Deps map[int][]int

func ParseRule(line string) Rule {
	vals := strings.Split(line, "|")
	page, _ := strconv.Atoi(vals[0])
	before, _ := strconv.Atoi(vals[1])
	return Rule{page, before}
}

func ParsePage(line string) Page {
	vals := strings.Split(line, ",")
	page := make([]int, 0)
	for _, val := range vals {
		n, _ := strconv.Atoi(val)
		page = append(page, n)
	}
	return page
}

func ParseDeps(rules []Rule) Deps {
	depsMap := make(map[int][]int)
	for _, rule := range rules {
		depsMap[rule.page] = append(depsMap[rule.page], rule.before)
	}
	return depsMap
}

func ParseInput(input []string) ([]Page, Deps) {
	rules := make([]Rule, 0)
	pages := make([]Page, 0)

	for _, line := range input {
		if strings.Contains(line, "|") {
			rules = append(rules, ParseRule(line))
		} else if strings.Contains(line, ",") {
			pages = append(pages, ParsePage(line))
		}
	}

	return pages, ParseDeps(rules)
}

func (page Page) Check(depsMap Deps) bool {
	for i := range len(page) {
		deps := depsMap[page[i]]
		for j := 0; j < i; j++ {
			if slices.Contains(deps, page[j]) {
				return false
			}
		}
	}

	return true
}

func (page Page) Fix(depsMap Deps) {
	slices.SortFunc(page, func(a int, b int) int {
		if slices.Contains(depsMap[a], b) {
			return -1
		} else {
			return 0
		}
	})
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

func PartOne(input []string) int {
	sum := 0
	pages, deps := ParseInput(input)

	for _, page := range pages {
		if page.Check(deps) {
			sum += page[int(len(page)/2)]
		}
	}

	return sum
}

func PartTwo(input []string) int {
	sum := 0
	pages, deps := ParseInput(input)

	for _, page := range pages {
		if !page.Check(deps) {
			page.Fix(deps)
			sum += page[int(len(page)/2)]
		}
	}

	return sum
}
