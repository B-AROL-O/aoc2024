package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInBounds(cell Cell, rows int, cols int) bool {
	return cell.r >= 0 && cell.r < rows && cell.c >= 0 && cell.c < cols
}

type Cell struct {
	r, c int
}

func countTrails(curr Cell, lines []string, rows int, cols int, visited *map[Cell]bool, cnt int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	if lines[curr.r][curr.c] == '9' {
		return cnt + 1
	}
	for _, dir := range dirs {
		next := Cell{curr.r + dir[0], curr.c + dir[1]}
		if checkInBounds(next, rows, cols) &&
			lines[next.r][next.c] == lines[curr.r][curr.c]+1 &&
			!(*visited)[next] {
			(*visited)[next] = true
			cnt = countTrails(next, lines, rows, cols, visited, cnt)
		}
	}
	return cnt
}

func countDistinctTrails(curr Cell, lines []string, rows int, cols int, cnt int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	if lines[curr.r][curr.c] == '9' {
		return cnt + 1
	}
	for _, dir := range dirs {
		next := Cell{curr.r + dir[0], curr.c + dir[1]}
		if checkInBounds(next, rows, cols) &&
			lines[next.r][next.c] == lines[curr.r][curr.c]+1 {
			cnt = countDistinctTrails(next, lines, rows, cols, cnt)
		}
	}
	return cnt
}

func part1(lines []string) {
	rows, cols := len(lines), len(lines[0])
	cnt := 0
	for r := range rows {
		for c := range cols {
			if lines[r][c] == '0' {
				visited := make(map[Cell]bool)
				cnt += countTrails(Cell{r, c}, lines, rows, cols, &visited, 0)
			}
		}
	}

	fmt.Println(cnt)
}

func part2(lines []string) {
	rows, cols := len(lines), len(lines[0])
	cnt := 0
	for r := range rows {
		for c := range cols {
			if lines[r][c] == '0' {
				cnt += countDistinctTrails(Cell{r, c}, lines, rows, cols, 0)
			}
		}
	}

	fmt.Println(cnt)
}

func main() {
	data, err := os.ReadFile("./input10_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
