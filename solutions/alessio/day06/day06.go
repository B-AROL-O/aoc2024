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

type Cell struct {
	r, c int
}

func part1(lines []string) {
	// see https://go.dev/blog/maps#key-types for more details about this implementation vs using a map of maps
	rows := len(lines)
	cols := len(lines[0])
	var startRow, startCol int
	found := false
	for r := range rows {
		for c := range cols {
			if lines[r][c] == '^' {
				startRow = r
				startCol = c
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	r, c := startRow, startCol
	dirR, dirC := -1, 0 // up
	visited := make(map[Cell]bool)
	for {
		if lines[r][c] != '#' && !visited[Cell{r, c}] {
			visited[Cell{r, c}] = true
		}
		nextR, nextC := r+dirR, c+dirC
		if nextR < 0 || nextC < 0 || nextR >= rows || nextC >= cols {
			break
		}
		if lines[nextR][nextC] == '#' {
			dirR, dirC = dirC, -dirR
		}
		r += dirR
		c += dirC
	}

	fmt.Println(len(visited))
}

type VisitedCellKey struct {
	r, c, rDir, cDir int
}

func part2(lines []string) {
	rows := len(lines)
	cols := len(lines[0])
	var startRow, startCol int
	found := false
	for r := range rows {
		for c := range cols {
			if lines[r][c] == '^' {
				startRow = r
				startCol = c
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	loops := make(map[Cell]bool)
	for r := range rows {
		for c := range cols {
			if lines[r][c] == '.' {
				visited := make(map[VisitedCellKey]bool)
				// temporary put a block here
				lines[r] = lines[r][:c] + "#" + lines[r][c+1:]

				i, j := startRow, startCol
				dirR, dirC := -1, 0 // up
				for {
					isVisited := visited[VisitedCellKey{i, j, dirR, dirC}]
					if lines[i][j] != '#' && !isVisited {
						visited[VisitedCellKey{i, j, dirR, dirC}] = true
					}
					if lines[i][j] != '#' && isVisited {
						// fmt.Printf("Found loop with block in (%d, %d)\n", r, c)
						loops[Cell{r, c}] = true
						break
					}
					nextR, nextC := i+dirR, j+dirC
					if nextR < 0 || nextC < 0 || nextR >= rows || nextC >= cols {
						break
					}
					if lines[nextR][nextC] == '#' {
						dirR, dirC = dirC, -dirR
						isVisited = visited[VisitedCellKey{i, j, dirR, dirC}]
						if isVisited {
							loops[Cell{r, c}] = true
							break
						} else {
							visited[VisitedCellKey{i, j, dirR, dirC}] = true
						}
					}
					i += dirR
					j += dirC
				}
				// backtrack
				lines[r] = lines[r][:c] + "." + lines[r][c+1:]
			}
		}
	}
	fmt.Println(len(loops))
}

func main() {
	data, err := os.ReadFile("./input06.txt")
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	// part1(lines)
	part2(lines)
}
