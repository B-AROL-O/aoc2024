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

func checkInBounds(cell Cell, rows int, cols int) bool {
	return cell.r >= 0 && cell.r < rows && cell.c >= 0 && cell.c < cols
}

type Cell struct {
	r, c int
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

type vertKey struct {
	col, dir int
}

type horizKey struct {
	row, dir int
}

func countSides(r []Cell, grid []string, rows int, cols int) int {
	symbol := grid[r[0].r][r[0].c]
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // N, E, S, W
	vertSides := make(map[vertKey][]int)
	horizSides := make(map[horizKey][]int)
	for _, cell := range r {
		for d, dir := range dirs {
			next := Cell{cell.r + dir[0], cell.c + dir[1]}
			if !checkInBounds(next, rows, cols) || grid[next.r][next.c] != symbol {
				// perimetral cell
				if d == N {
					horizSides[horizKey{cell.r, d}] = append(horizSides[horizKey{cell.r, d}], cell.c)
				} else if d == S {
					horizSides[horizKey{cell.r + 1, d}] = append(horizSides[horizKey{cell.r + 1, d}], cell.c)
				} else if d == W {
					vertSides[vertKey{cell.c, d}] = append(vertSides[vertKey{cell.c, d}], cell.r)
				} else {
					vertSides[vertKey{cell.c + 1, d}] = append(vertSides[vertKey{cell.c + 1, d}], cell.r)
				}
			}
		}
	}

	cnt := 0
	for _, v := range horizSides {
		sort.Ints(v)
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] > 1 {
				cnt++
			}
		}
		cnt++
	}

	for _, v := range vertSides {
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
	regions := make([][]Cell, 0)
	rows, cols := len(grid), len(grid[0])
	for r := range rows {
		for c := range cols {
			if !visited[Cell{r, c}] {
				region := make([]Cell, 0)
				dfs(Cell{r, c}, grid, rows, cols, &visited, &region)
				regions = append(regions, region)
			}
		}
	}

	res := 0
	for _, r := range regions {
		area := len(r)
		perimeter := 0
		if !useSides {
			perimeter = computePerimeter(r, grid, rows, cols)
		} else {
			perimeter = countSides(r, grid, rows, cols)
			fmt.Printf("- A region of %c plants with price %d * %d = %d\n", grid[r[0].r][r[0].c], area, perimeter, area*perimeter)
		}
		res += area * perimeter
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
	data, err := os.ReadFile("./input12.txt")
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
