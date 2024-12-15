package main

import (
	"errors"
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

func checkInBounds(cell Pos, rows int, cols int) bool {
	return cell.r >= 0 && cell.r < rows && cell.c >= 0 && cell.c < cols
}

func findStartingPos(grid []string, rows int, cols int) (Pos, error) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' {
				return Pos{r, c}, nil
			}
		}
	}
	return Pos{-1, -1}, errors.New("could not find starting pos")
}

func getMoveDir(move byte) Pos {
	switch move {
	case '^':
		return Pos{-1, 0}
	case 'v':
		return Pos{1, 0}
	case '<':
		return Pos{0, -1}
	case '>':
		return Pos{0, 1}
	}
	fmt.Println("invalid move", move)
	return Pos{0, 0}
}

func replaceInString(s string, pos int, c string) string {
	return s[:pos] + c + s[pos+1:]
}

func printGrid(grid []string, rows int, cols int, currPos Pos) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == currPos.r && c == currPos.c {
				fmt.Print("@")
			} else {
				fmt.Printf("%c", grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func sumGPSCoords(grid []string, rows int, cols int, symbol byte) int {
	sum := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == symbol {
				sum += r*100 + c
			}
		}
	}

	return sum
}

func solve1(origGrid []string, moves string) {
	grid := make([]string, len(origGrid))
	copy(grid, origGrid)

	rows, cols := len(grid), len(grid[0])
	currPos, err := findStartingPos(grid, rows, cols)
	check(err)

	grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
	for i := range moves {
		dir := getMoveDir(moves[i])
		next := Pos{currPos.r + dir.r, currPos.c + dir.c}
		if grid[next.r][next.c] == '.' {
			currPos = next
		} else if grid[next.r][next.c] == 'O' {
			for grid[next.r][next.c] == 'O' {
				next.r += dir.r
				next.c += dir.c
			}
			if grid[next.r][next.c] == '.' {
				// roll blocks
				grid[next.r] = replaceInString(grid[next.r], next.c, "O")
				currPos.r += dir.r
				currPos.c += dir.c
				grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
			}
		}
	}
	fmt.Println(sumGPSCoords(grid, rows, cols, 'O'))
}

func transformGrid(grid []string) []string {
	newGrid := []string{}
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		newRow := ""
		for c := 0; c < cols; c++ {
			if grid[r][c] == '.' {
				newRow += ".."
			} else if grid[r][c] == '#' {
				newRow += "##"
			} else if grid[r][c] == '@' {
				newRow += "@."
			} else {
				newRow += "[]"
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func isBox(c byte) bool {
	return c == '[' || c == ']'
}

func canRoll(grid []string, currPos Pos, dir Pos) bool {
	nextPos1 := Pos{currPos.r + dir.r, currPos.c + dir.c} // above/below left side of box
	nextPos2 := Pos{nextPos1.r, nextPos1.c + 1}           // above/below right side of box

	block1 := grid[nextPos1.r][nextPos1.c]
	block2 := grid[nextPos2.r][nextPos2.c]
	if block1 == '.' && block2 == '.' {
		return true
	}
	if block1 == '#' || block2 == '#' {
		return false
	}

	ret := true
	if block1 == '[' {
		ret = ret && canRoll(grid, nextPos1, dir)
	} else if block1 == ']' {
		ret = ret && canRoll(grid, Pos{nextPos1.r, nextPos1.c - 1}, dir)
	}

	if block2 == '[' {
		ret = ret && canRoll(grid, nextPos2, dir)
	} else if block2 == ']' {
		ret = ret && canRoll(grid, Pos{nextPos2.r, nextPos2.c - 1}, dir)
	}

	return ret
}

func rollBoxes(grid []string, currPos Pos, dir Pos) {
	nextPos1 := Pos{currPos.r + dir.r, currPos.c + dir.c} // above/below left side of box
	nextPos2 := Pos{nextPos1.r, nextPos1.c + 1}           // above/below right side of box

	block1 := grid[nextPos1.r][nextPos1.c]
	block2 := grid[nextPos2.r][nextPos2.c]

	if block1 == '.' && block2 == '.' {
		grid[nextPos1.r] = replaceInString(grid[nextPos1.r], nextPos1.c, "[")
		grid[nextPos2.r] = replaceInString(grid[nextPos2.r], nextPos2.c, "]")
		grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
		grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c+1, ".")
		return
	}
	if block1 == '[' {
		rollBoxes(grid, nextPos1, dir)
	} else if block1 == ']' {
		rollBoxes(grid, Pos{nextPos1.r, nextPos1.c - 1}, dir)
	}

	if block2 == '[' {
		rollBoxes(grid, nextPos2, dir)
	}

	isNowFree := grid[nextPos1.r][nextPos1.c] == '.' && grid[nextPos2.r][nextPos2.c] == '.'
	if !isNowFree {
		fmt.Println("blocks not freed correctly!")
	} else {
		grid[nextPos1.r] = replaceInString(grid[nextPos1.r], nextPos1.c, "[")
		grid[nextPos2.r] = replaceInString(grid[nextPos2.r], nextPos2.c, "]")
		grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
		grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c+1, ".")
	}
}

func solve2(origGrid []string, moves string) {
	grid := make([]string, len(origGrid))
	copy(grid, origGrid)

	rows, cols := len(grid), len(grid[0])
	currPos, err := findStartingPos(grid, rows, cols)
	check(err)

	grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
	for i := range moves {
		dir := getMoveDir(moves[i])
		next := Pos{currPos.r + dir.r, currPos.c + dir.c}
		if grid[next.r][next.c] == '.' {
			currPos = next
		} else if isBox(grid[next.r][next.c]) {
			if moves[i] == '<' || moves[i] == '>' {
				for isBox(grid[next.r][next.c]) {
					next.r += dir.r
					next.c += dir.c
				}

				if grid[next.r][next.c] == '.' {
					// roll blocks
					for next != currPos {
						prevPos := Pos{next.r - dir.r, next.c - dir.c}
						prevSym := fmt.Sprintf("%c", grid[prevPos.r][prevPos.c])
						grid[next.r] = replaceInString(grid[next.r], next.c, prevSym)
						next = prevPos
					}
					grid[currPos.r] = replaceInString(grid[currPos.r], currPos.c, ".")
					currPos.r += dir.r
					currPos.c += dir.c
				}
			} else {
				boxPos := Pos{next.r, next.c}
				if grid[next.r][next.c] == ']' {
					boxPos.c--
				}
				if canRoll(grid, boxPos, dir) {
					rollBoxes(grid, boxPos, dir)
					currPos.r += dir.r
					currPos.c += dir.c
				}
			}
		}
	}
	fmt.Println(sumGPSCoords(grid, rows, cols, '['))
}

func part1(grid []string, moves string) {
	solve1(grid, moves)
}

func part2(grid []string, moves string) {
	newGrid := transformGrid(grid)
	solve2(newGrid, moves)
}

func main() {
	data, err := os.ReadFile("./input15_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	breakpoint := 0
	for {
		if lines[breakpoint] == "" {
			break
		}
		breakpoint++
	}

	start := time.Now()
	part1(lines[:breakpoint], strings.Join(lines[breakpoint+1:], ""))
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(lines[:breakpoint], strings.Join(lines[breakpoint+1:], ""))
	fmt.Printf("part2: %s\n", time.Since(start))
}
