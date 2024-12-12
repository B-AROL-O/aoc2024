package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Cell struct {
	r, c int
}

func checkInBounds(cell Cell, rows int, cols int) bool {
	return cell.r >= 0 && cell.r < rows && cell.c >= 0 && cell.c < cols
}

func dfs(curr Cell, lines []string, rows int, cols int, visited *map[Cell]bool, region *[]Cell) {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	(*visited)[curr] = true
	*region = append(*region, curr)
	for _, dir := range dirs {
		next := Cell{curr.r + dir[0], curr.c + dir[1]}
		if checkInBounds(next, rows, cols) &&
			lines[next.r][next.c] == lines[curr.r][curr.c] &&
			!(*visited)[next] {
			dfs(next, lines, rows, cols, visited, region)
		}
	}
}

func computePerimeter(r []Cell, grid []string, rows int, cols int) int {
	symbol := grid[r[0].r][r[0].c]

	perimeter := 0
	for _, cell := range r {
		for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Cell{cell.r + dir[0], cell.c + dir[1]}
			if !checkInBounds(next, rows, cols) || grid[next.r][next.c] != symbol {
				perimeter++
			}
		}
	}

	return perimeter
}

const (
	N = iota
	E
	S
	W
)

type sideKey struct {
	row, col, dir int
}

func countSides(r []Cell, grid []string, rows int, cols int) int {
	symbol := grid[r[0].r][r[0].c]
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // N, E, S, W
	sides := make(map[sideKey][]int)
	for _, cell := range r {
		for d, dir := range dirs {
			next := Cell{cell.r + dir[0], cell.c + dir[1]}
			if !checkInBounds(next, rows, cols) || grid[next.r][next.c] != symbol {
				// perimetral cell
				if d == N {
					sides[sideKey{cell.r, -1, d}] = append(sides[sideKey{cell.r, -1, d}], cell.c)
				} else if d == S {
					sides[sideKey{cell.r + 1, -1, d}] = append(sides[sideKey{cell.r + 1, -1, d}], cell.c)
				} else if d == W {
					sides[sideKey{-1, cell.c, d}] = append(sides[sideKey{-1, cell.c, d}], cell.r)
				} else {
					sides[sideKey{-1, cell.c + 1, d}] = append(sides[sideKey{-1, cell.c + 1, d}], cell.r)
				}
			}
		}
	}

	cnt := 0
	for _, v := range sides {
		sort.Ints(v)
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] > 1 {
				cnt++
			}
		}
		cnt++
	}

	return cnt
}

func solve(grid []string, useSides bool) {
	visited := make(map[Cell]bool)
	rows, cols := len(grid), len(grid[0])
	res := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[Cell{r, c}] {
				region := []Cell{}
				dfs(Cell{r, c}, grid, rows, cols, &visited, &region)
				area := len(region)
				perimeter := 0
				if !useSides {
					perimeter = computePerimeter(region, grid, rows, cols)
				} else {
					perimeter = countSides(region, grid, rows, cols)
					// fmt.Printf("- A region of %c plants with price %d * %d = %d\n", grid[region[0].r][region[0].c], area, perimeter, area*perimeter)
				}
				res += area * perimeter
			}
		}
	}

	fmt.Println(res)
}

func part1(grid []string) {
	solve(grid, false)
}

func part2(grid []string) {
	solve(grid, true)
}

func main() {
	data, err := os.ReadFile("./input12_def.txt")
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
