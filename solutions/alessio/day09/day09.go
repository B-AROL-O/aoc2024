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

func part1(line string) {
	res := make([]int, 0, len(line))
	for i := range len(line) {
		length := int(line[i] - '0')
		for range length {
			if i%2 == 0 {
				res = append(res, i/2)
			} else {
				res = append(res, -1)
			}
		}
	}

	l, r := 0, len(res)-1
	for l < r {
		if res[l] != -1 {
			l++
			continue
		}
		if res[r] == -1 {
			r--
			continue
		}

		res[l], res[r] = res[r], res[l]
		r--
		l++
	}

	checksum := int64(0)
	for pos, id := range res {
		if id == -1 {
			break
		}
		checksum += int64(pos) * int64(id)
	}

	fmt.Println(checksum)
}

func part2(line string) {
	files := make([]Block, 0, len(line)/2)
	free := make([]Block, 0, len(line)/2)
	pos := 0
	for i := 0; i < len(line); i++ {
		length := int(line[i] - '0')
		if i%2 == 0 {
			files = append(files, Block{pos, length})
		} else {
			free = append(free, Block{pos, length})
		}
		pos += length
	}

	for i := len(files) - 1; i >= 0; i-- {
		file := &files[i]
		for j := 0; j < len(free) && free[j].start < file.start; j++ {
			if free[j].length >= file.length {
				file.start = free[j].start
				free[j].start += file.length
				free[j].length -= file.length
			}
		}
	}

	checksum := int64(0)
	for i := range len(files) {
		for j := range files[i].length {
			checksum += int64(i) * int64(files[i].start+j)
		}
	}

	fmt.Println(checksum)
}

func main() {
	data, err := os.ReadFile("./input09_def.txt")
	check(err)

	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	line := strings.Trim(dataStr, "\n")

	// part1WithIntervals(line)
	start := time.Now()
	part1(line)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(line)
	fmt.Printf("part2: %s\n", time.Since(start))
}
