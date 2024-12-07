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

type PrecKey struct {
	before int
	after  int
}

func part1(lines []string) {
	// see https://go.dev/blog/maps#key-types for more details about this implementation vs using a map of maps
	prec := make(map[PrecKey]bool)
	updatesLine := 0
	for i, pair := range lines {
		if pair == "" {
			updatesLine = i + 1
			break
		}
		var a, b int
		fmt.Sscanf(pair, "%d|%d", &a, &b)
		prec[PrecKey{a, b}] = true
	}

	sumMiddle := 0
	for i := updatesLine; i < len(lines); i++ {
		isValid := true
		strs := strings.Split(lines[i], ",")
		nums := make([]int, 0, len(strs))
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		for j := range len(nums) {
			for k := j + 1; k < len(nums); k++ {
				if (!prec[PrecKey{nums[j], nums[k]}]) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			sumMiddle += nums[len(nums)/2]
		}
	}

	fmt.Println(sumMiddle)
}

func part2(lines []string) {
	prec := make(map[PrecKey]bool)
	updatesLine := 0
	for i, pair := range lines {
		if pair == "" {
			updatesLine = i + 1
			break
		}
		var a, b int
		fmt.Sscanf(pair, "%d|%d", &a, &b)
		prec[PrecKey{a, b}] = true
	}

	sumMiddle := 0
	for i := updatesLine; i < len(lines); i++ {
		isValid := true
		strs := strings.Split(lines[i], ",")
		nums := make([]int, 0, len(strs))
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		for j := range len(nums) {
			for k := j + 1; k < len(nums); k++ {
				if (!prec[PrecKey{nums[j], nums[k]}]) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if !isValid {
			correctNums := make([]int, 0, len(nums))
			for j := range len(nums) {
				correctNums = append(correctNums, nums[j])
			}

			for j := range len(nums) - 1 {
				for k := range len(nums) - j - 1 {
					if prec[PrecKey{correctNums[k+1], correctNums[k]}] {
						correctNums[k], correctNums[k+1] = correctNums[k+1], correctNums[k]
					}
				}
			}
			sumMiddle += correctNums[len(correctNums)/2]
		}
	}

	fmt.Println(sumMiddle)
}
func main() {
	data, err := os.ReadFile("./input05.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
