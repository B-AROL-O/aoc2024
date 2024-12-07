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

func findStartingCell(lines []string, rows int, cols int) (int, int) {
	var startRow, startCol int
	found := false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
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
	return startRow, startCol
}

func part1(lines []string) {
	rows, cols := len(lines), len(lines[0])
	r, c := findStartingCell(lines, rows, cols)
	dirR, dirC := -1, 0 // up
	// see https://go.dev/blog/maps#key-types for more details about this implementation vs using a map of maps
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
			if lines[r+dirR][c+dirC] == '#' {
				fmt.Println("two cons block, breaking")
				break
			}
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
	rows, cols := len(lines), len(lines[0])
	startRow, startCol := findStartingCell(lines, rows, cols)

	loops := make(map[Cell]bool)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
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
					} else if lines[i][j] != '#' && isVisited {
						// fmt.Printf("Found loop with block in (%d, %d)\n", r, c)
						loops[Cell{r, c}] = true
						break
					}
					if lines[i][j] == '#' {
						fmt.Println(r, c, dirR, dirC)
					}
					nextR, nextC := i+dirR, j+dirC
					if nextR < 0 || nextC < 0 || nextR >= rows || nextC >= cols {
						break // loop not found, exit
					}

					cntChDir := 0
					// turn 90 degrees until there is free space in front
					for lines[i+dirR][j+dirC] == '#' {
						cntChDir++
						if cntChDir >= 5 {
							fmt.Println("wtf")
							break
						}
						dirR, dirC = dirC, -dirR // turn 90 degrees
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
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
