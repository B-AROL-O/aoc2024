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

func part1(lines []string) {
	keys := [][]int{}
	locks := [][]int{}
	for i := 0; i < len(lines); i += 8 {
		heights := []int{0, 0, 0, 0, 0}
		for j := i + 1; j < i+6; j++ {
			for k := 0; k < 5; k++ {
				if lines[j][k] == '#' {
					heights[k]++
				}
			}
		}
		if lines[i] == "....." {
			keys = append(keys, heights)
		} else {
			locks = append(locks, heights)
		}
	}

	cnt := 0
	for _, l := range locks {
		for _, k := range keys {
			fits := true
			for c := 0; c < 5; c++ {
				if l[c]+k[c] > 5 {
					fits = false
					break
				}
			}
			if fits {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

func part2(lines []string) {
}

func main() {
	data, err := os.ReadFile("./input25_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	fmt.Printf("part2: %s\n", time.Since(start))
	part2(lines)
}
