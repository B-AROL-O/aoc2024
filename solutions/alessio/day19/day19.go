package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve(lines []string) {

}

func matchesPermutation(towels []string, match string, pos int, currComb string) bool {
	if pos == len(match) {
		return true
	}
	for _, t := range towels {
		if pos+len(t) <= len(match) && t == match[pos:pos+len(t)] {
			if matchesPermutation(towels, match, pos+len(t), currComb+t) {
				return true
			}
		}
	}
	return false
}

func countMatchingPermutations(towels []string, match string, pos int, currComb string, memo *map[string]int) int {
	if _, exists := (*memo)[currComb]; exists {
		return (*memo)[currComb]
	}

	if pos == len(match) {
		return 1
	}
	cnt := 0
	for _, t := range towels {
		if pos+len(t) <= len(match) && t == match[pos:pos+len(t)] {
			cnt += countMatchingPermutations(towels, match, pos+len(t), currComb+t, memo)
		}
	}
	(*memo)[currComb] = cnt
	return cnt
}

func part1(lines []string) {
	towels := strings.Split(lines[0], ", ")

	cnt := 0
	for _, match := range lines[2:] {
		if matchesPermutation(towels, match, 0, "") {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func part2(lines []string) {
	towels := strings.Split(lines[0], ", ")

	cnt := 0
	for _, match := range lines[2:] {
		memo := make(map[string]int)
		cnt += countMatchingPermutations(towels, match, 0, "", &memo)
	}
	fmt.Println(cnt)
}

func main() {
	data, err := os.ReadFile("./input19_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(lines)
	fmt.Printf("part2: %s\n", time.Since(start))
}
