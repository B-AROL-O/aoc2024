package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(lines []string) {
	var left []int
	var right []int
	sumDiff := 0

	for _, l := range lines {
		var first, second int
		_, err := fmt.Sscanf(l, "%d %d", &first, &second)
		check(err)
		left = append(left, first)
		right = append(right, second)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range len(lines) {
		sumDiff += abs(right[i] - left[i])
	}
	_, err := fmt.Println(sumDiff)
	check(err)
}

func part2(lines []string) {
	leftCnt := make(map[int]int)
	rightCnt := make(map[int]int)

	for _, l := range lines {
		var first, second int
		_, err := fmt.Sscanf(l, "%d %d", &first, &second)
		check(err)
		leftCnt[first]++
		rightCnt[second]++
	}

	simScore := 0

	for k, v := range leftCnt {
		simScore += k * v * rightCnt[k]
	}

	_, err := fmt.Println(simScore)
	check(err)
}

func main() {
	data, err := os.ReadFile("./input01.txt")
	check(err)
	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	part1(lines)
	part2(lines)
}
