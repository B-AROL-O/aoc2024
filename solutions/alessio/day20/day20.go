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

type Pos struct {
	r, c int
}

var dirs []Pos = []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // E, W, S, N

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findStart(grid []string) Pos {
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				return Pos{r, c}
			}
		}
	}
	return Pos{}
}

func solve(grid []string, maxDist int) {
	start := findStart(grid)
	q := []Pos{start}
	visited := map[Pos]bool{}
	visited[start] = true
	orderedPath := []Pos{}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		orderedPath = append(orderedPath, p)
		for _, d := range dirs {
			adj := Pos{p.r + d.r, p.c + d.c}
			if !visited[adj] && grid[adj.r][adj.c] != '#' {
				q = append(q, adj)
				visited[adj] = true
			}
		}
	}

	cnt := 0
	for i := range len(orderedPath) {
		for j := i + 1; j < len(orderedPath); j++ {
			pathDist := j - i
			manDist := abs(orderedPath[i].r-orderedPath[j].r) + abs(orderedPath[i].c-orderedPath[j].c)
			if manDist <= maxDist && pathDist-manDist >= 100 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func part1(lines []string) {
	solve(lines, 2)
}

func part2(lines []string) {
	solve(lines, 20)
}

func main() {
	data, err := os.ReadFile("./input20_def.txt")
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
