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

func checkInBounds(cell Cell, rows int, cols int) bool {
	return cell.r >= 0 && cell.r < rows && cell.c >= 0 && cell.c < cols
}

func part1(lines []string) {
	rows, cols := len(lines), len(lines[0])
	antennas := make(map[byte][]Cell)
	for r := range rows {
		for c := range cols {
			if lines[r][c] != '.' {
				antennas[lines[r][c]] = append(antennas[lines[r][c]], Cell{r, c})
			}
		}
	}

	antinodes := make(map[Cell]bool)
	for _, places := range antennas {
		for i := range len(places) {
			for j := i + 1; j < len(places); j++ {
				p1, p2 := places[i], places[j]
				dr, dc := p2.r-p1.r, p2.c-p1.c
				antinode1 := Cell{p2.r + dr, p2.c + dc}
				if checkInBounds(antinode1, rows, cols) {
					antinodes[antinode1] = true
				}
				antinode2 := Cell{p1.r - dr, p1.c - dc}
				if checkInBounds(antinode2, rows, cols) {
					antinodes[antinode2] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func part2(lines []string) {
	rows, cols := len(lines), len(lines[0])
	antennas := make(map[byte][]Cell)
	for r := range rows {
		for c := range cols {
			if lines[r][c] != '.' {
				antennas[lines[r][c]] = append(antennas[lines[r][c]], Cell{r, c})
			}
		}
	}

	antinodes := make(map[Cell]bool)
	for _, places := range antennas {
		for i := range len(places) {
			for j := i + 1; j < len(places); j++ {
				p1, p2 := places[i], places[j]
				// mark also antennas as antinodes
				antinodes[p1] = true
				antinodes[p2] = true

				dr, dc := p2.r-p1.r, p2.c-p1.c
				cnt := 1
				for {
					antinode := Cell{p2.r + cnt*dr, p2.c + cnt*dc}
					if checkInBounds(antinode, rows, cols) {
						antinodes[antinode] = true
						cnt++
					} else {
						break
					}
				}

				cnt = 1
				for {
					antinode := Cell{p1.r - cnt*dr, p1.c - cnt*dc}
					if checkInBounds(antinode, rows, cols) {
						antinodes[antinode] = true
						cnt++
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func main() {
	data, err := os.ReadFile("./input08.txt")
	check(err)

	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
