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

func checkEquation(seq []int, tot int, currPos int, currTot int, concatOp bool) bool {
	if currPos == len(seq) {
		if tot == currTot {
			return true
		} else {
			return false
		}
	}
	if currPos == 1 {
		currTot = seq[0]
	}
	if currTot+seq[currPos] <= tot {
		if checkEquation(seq, tot, currPos+1, currTot+seq[currPos], concatOp) {
			return true
		}
	}
	if currTot*seq[currPos] <= tot {
		if checkEquation(seq, tot, currPos+1, currTot*seq[currPos], concatOp) {
			return true
		}
	}
	if concatOp {
		resStr := fmt.Sprintf("%d%d", currTot, seq[currPos])
		resNum, err := strconv.Atoi(resStr)
		check(err)
		if resNum <= tot && checkEquation(seq, tot, currPos+1, resNum, concatOp) {
			return true
		}
	}
	return false
}

func part1(lines []string) {
	sumRes := 0
	for _, l := range lines {
		tot, err := strconv.Atoi(strings.Split(l, ": ")[0])
		check(err)
		numsStr := strings.Split(strings.Split(l, ": ")[1], " ")
		nums := make([]int, 0, len(numsStr))
		for _, s := range numsStr {
			n, err := strconv.Atoi(s)
			check(err)
			nums = append(nums, n)
		}
		if checkEquation(nums, tot, 1, 0, false) {
			sumRes += tot
		}
	}
	fmt.Println(sumRes)
}

func part2(lines []string) {
	sumRes := 0
	for _, l := range lines {
		tot, err := strconv.Atoi(strings.Split(l, ": ")[0])
		check(err)
		numsStr := strings.Split(strings.Split(l, ": ")[1], " ")
		nums := make([]int, 0, len(numsStr))
		for _, s := range numsStr {
			n, err := strconv.Atoi(s)
			check(err)
			nums = append(nums, n)
		}
		if checkEquation(nums, tot, 1, 0, true) {
			sumRes += tot
		}
	}
	fmt.Println(sumRes)
}

func main() {
	data, err := os.ReadFile("./input07.txt")
	check(err)

	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
