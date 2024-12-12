package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type memoKey struct {
	n   int64
	ttl int
}

func countLeaves(n int64, memo *map[memoKey]int, ttl int) int {
	if ttl == 0 { // leaf, return 1 to count just itself
		return 1
	}

	if v, ok := (*memo)[memoKey{n, ttl}]; ok {
		return v
	}

	res := 0
	nstr := strconv.FormatInt(n, 10)
	if n == 0 {
		res += countLeaves(1, memo, ttl-1)
	} else if len(nstr)%2 == 0 {
		l := len(nstr)
		half1, _ := strconv.ParseInt(nstr[:l/2], 10, 0)
		half2, _ := strconv.ParseInt(nstr[l/2:], 10, 0)
		res += countLeaves(half1, memo, ttl-1)
		res += countLeaves(half2, memo, ttl-1)
	} else {
		res += countLeaves(n*2024, memo, ttl-1)
	}
	(*memo)[memoKey{n, ttl}] = res
	return res
}

func bfs(nums []int, ttl int) int {
	queue := make([]memoKey, 0)
	for n := range nums {
		queue = append(queue, memoKey{int64(nums[n]), ttl})
	}

	cnt := 0
	currTtl := ttl
	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]
		if k.ttl == 0 {
			cnt++
		} else if k.ttl == -1 {
			continue
		}

		if k.ttl < currTtl {
			currTtl = k.ttl
			// fmt.Println()
		}
		// fmt.Printf("%d ", k.n)
		nstr := strconv.FormatInt(k.n, 10)
		if k.n == 0 {
			queue = append(queue, memoKey{1, k.ttl - 1})
		} else if len(nstr)%2 == 0 {
			l := len(nstr)
			half1, _ := strconv.ParseInt(nstr[:l/2], 10, 0)
			half2, _ := strconv.ParseInt(nstr[l/2:], 10, 0)
			queue = append(queue, memoKey{half1, k.ttl - 1})
			queue = append(queue, memoKey{half2, k.ttl - 1})
		} else {
			queue = append(queue, memoKey{k.n * 2024, k.ttl - 1})
		}
	}
	// fmt.Println()
	return cnt
}

func solve(line string, depth int) {
	numsStr := strings.Split(line, " ")
	nums := make([]int, 0, len(numsStr))
	for _, s := range numsStr {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	// fmt.Println(bfs(nums, depth))

	cnt := 0
	memo := make(map[memoKey]int)
	for _, n := range nums {
		cnt += countLeaves(int64(n), &memo, depth)
	}

	fmt.Println(cnt)
}

func part1(line string) {
	solve(line, 25)
}

func part2(line string) {
	solve(line, 75)
}

func main() {
	data, err := os.ReadFile("./input11_def.txt")
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	start := time.Now()
	part1(lines[0])
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(lines[0])
	fmt.Printf("part2: %s\n", time.Since(start))
}
