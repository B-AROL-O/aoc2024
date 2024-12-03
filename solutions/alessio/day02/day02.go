package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_if_safe(levels []int) bool {
	is_incr := levels[1]-levels[0] > 0
	for i := range len(levels) - 1 {
		diff := levels[i+1] - levels[i]
		if is_incr && (diff < 1 || diff > 3) {
			return false
		} else if !is_incr && (diff > -1 || diff < -3) {
			return false
		}
	}
	return true
}

func part1(lines []string) {
	safes := 0

	for _, l := range lines {
		levels_str := strings.Split(l, " ")

		var levels []int
		for i := range len(levels_str) {
			lvl, err := strconv.Atoi(levels_str[i])
			check(err)
			levels = append(levels, lvl)
		}

		if check_if_safe(levels) {
			safes += 1
		}
	}

	fmt.Println(safes)
}

func part2(lines []string) {
	safes := 0

	for _, l := range lines {
		levels_str := strings.Split(l, " ")

		var levels []int
		for i := range len(levels_str) {
			lvl, err := strconv.Atoi(levels_str[i])
			check(err)
			levels = append(levels, lvl)
		}

		is_safe := check_if_safe(levels)
		i := 0
		for !is_safe && i < len(levels) {
			levels_slice := slices.Concat(levels[:i], levels[i+1:]) // basically splice on i-th element
			is_safe = check_if_safe(levels_slice)
			i++
		}

		if is_safe {
			safes += 1
		}
	}

	fmt.Println(safes)
}

func main() {
	data, err := os.ReadFile("./input02.txt")
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	part1(lines)
	part2(lines)
}
