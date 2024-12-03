package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var first_list []int
	var second_list []int
	sum_diff := 0

	for _, l := range lines {
		vals := strings.Split(l, "   ")
		first, err := strconv.Atoi(vals[0])
		check(err)
		second, err := strconv.Atoi(vals[1])
		check(err)
		first_list = append(first_list, first)
		second_list = append(second_list, second)
	}

	sort.Ints(first_list)
	sort.Ints(second_list)

	for i := range len(lines) {
		sum_diff += abs(second_list[i] - first_list[i])
	}
	fmt.Println(sum_diff)
}

func part2(lines []string) {
	first_cnt := make(map[int]int)
	second_cnt := make(map[int]int)

	for _, l := range lines {
		vals := strings.Split(l, "   ")
		first, err := strconv.Atoi(vals[0])
		check(err)
		second, err := strconv.Atoi(vals[1])
		check(err)
		first_cnt[first] += 1
		second_cnt[second] += 1
	}

	sim_score := 0

	for k, v := range first_cnt {
		sim_score += k * v * second_cnt[k]
	}

	fmt.Println(sim_score)
}

func main() {
	data, err := os.ReadFile("./input01.txt")
	check(err)
	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	part1(lines)
	part2(lines)
}
