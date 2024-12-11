package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIfSafe(levels []int) bool {
	isIncr := levels[1]-levels[0] > 0
	for i := range len(levels) - 1 {
		diff := levels[i+1] - levels[i]
		if isIncr && (diff < 1 || diff > 3) {
			return false
		} else if !isIncr && (diff > -1 || diff < -3) {
			return false
		}
	}
	return true
}

func part1(lines []string) {
	safes := 0

	for _, l := range lines {
		levelsStr := strings.Split(l, " ")
		levels := make([]int, len(levelsStr))
		for i, s := range levelsStr {
			levels[i], _ = strconv.Atoi(s)
		}

		if checkIfSafe(levels) {
			safes++
		}
	}

	_, err := fmt.Println(safes)
	check(err)
}

func part2(lines []string) {
	safes := 0

	for _, l := range lines {
		levelsStr := strings.Split(l, " ")
		levels := make([]int, len(levelsStr))
		for i, s := range levelsStr {
			levels[i], _ = strconv.Atoi(s)
		}

		if checkIfSafe(levels) {
			safes++
			continue
		}

		for i := range levels {
			if checkIfSafe(append(levels[:i], levels[i+1:]...)) {
				safes++
				break
			}
		}
	}

	_, err := fmt.Println(safes)
	check(err)
}

func main() {
	data, err := os.ReadFile("./input02.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
