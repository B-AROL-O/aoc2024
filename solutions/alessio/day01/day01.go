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
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, l := range lines {
		_, err := fmt.Sscanf(l, "%d %d", &left[i], &right[i])
		check(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	sumDiff := 0
	for i := range lines {
		sumDiff += abs(right[i] - left[i])
	}

	_, err := fmt.Println(sumDiff)
	check(err)
}

func part2(lines []string) {
	leftCnt := make(map[int]int, len(lines))
	rightCnt := make(map[int]int, len(lines))

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
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
