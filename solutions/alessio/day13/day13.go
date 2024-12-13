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

type button struct {
	x, y int
}

func checkValidSolution(nA int, nB int, upperBound int) bool {
	if upperBound != -1 && (nA > upperBound || nB > upperBound) {
		return false
	}
	return nA > 0 && nB > 0
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func solveEquation(a button, b button, p button) (int, int) {
	nB := float64(p.y*a.x-p.x*a.y) / float64(b.y*a.x-a.y*b.x)
	if !isIntegral(nB) { // not int solution
		return -1, -1
	}

	nA := float64(p.x-b.x*int(nB)) / float64(a.x)
	if !isIntegral(nA) { // not int solution
		return -1, -1
	}

	return int(nA), int(nB)
}

func solve(lines []string, upperBound int) {
	totTokens := 0
	offset := 0
	for offset < len(lines) {
		data := lines[offset : offset+3]
		offset += 4 // skip also blank lines
		var a, b, prize button
		fmt.Sscanf(data[0], "Button A: X+%d, Y+%d", &a.x, &a.y)
		fmt.Sscanf(data[1], "Button B: X+%d, Y+%d", &b.x, &b.y)
		fmt.Sscanf(data[2], "Prize: X=%d, Y=%d", &prize.x, &prize.y)

		if upperBound == -1 {
			// part2
			prize.x += 10000000000000
			prize.y += 10000000000000
		}

		nA, nB := solveEquation(a, b, prize)
		if checkValidSolution(nA, nB, upperBound) {
			totTokens += 3*nA + nB
		}
	}

	fmt.Println(totTokens)
}

func part1(lines []string) {
	solve(lines, 100)

}

func part2(lines []string) {
	solve(lines, -1)
}

func main() {
	data, err := os.ReadFile("./input13_def.txt")
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
