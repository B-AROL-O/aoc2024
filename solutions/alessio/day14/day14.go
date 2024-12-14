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

type coord struct {
	x, y int
}

func solveEquation(p coord, v coord, t int) coord {
	return coord{p.x + v.x*t, p.y + v.y*t}
}

func shiftPoint(p coord, height int, width int) coord {
	p.x = p.x % width
	p.y = p.y % height
	if p.x < 0 {
		p.x = p.x + width
	}
	if p.y < 0 {
		p.y = p.y + height
	}
	return p
}

func solve(lines []string, time int, height int, width int) {
	quads := map[int]int{}
	for _, l := range lines {
		var p, v coord
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)

		pf := solveEquation(p, v, time)
		pf = shiftPoint(pf, height, width)

		if pf.x < width/2 {
			if pf.y < height/2 {
				quads[0]++
			} else if pf.y > height/2 {
				quads[3]++
			}
		} else if pf.x > width/2 {
			if pf.y < height/2 {
				quads[1]++
			} else if pf.y > height/2 {
				quads[2]++
			}
		}
	}

	mult := 1
	for i := range len(quads) {
		fmt.Printf("quad %d: %d\n", i, quads[i])
		mult *= quads[i]
	}

	fmt.Println(mult)
}

const (
	WIDTH_EXAMPLE  = 11
	HEIGHT_EXAMPLE = 7
	WIDTH_DEF      = 101
	HEIGHT_DEF     = 103
)

func part1(lines []string) {
	// solve(lines, 100, HEIGHT_EXAMPLE, WIDTH_EXAMPLE)
	solve(lines, 100, HEIGHT_DEF, WIDTH_DEF)
}

func part2(lines []string) {
	startPoints := []coord{}
	velocities := []coord{}
	for _, l := range lines {
		var p, v coord
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)
		startPoints = append(startPoints, p)
		velocities = append(velocities, v)
	}

	time := 1
	for time <= 101*103 {
		fmt.Printf("\n\n****** TIME %d ******\n\n", time)
		currPoints := map[coord]bool{}
		for i, p := range startPoints {
			pf := solveEquation(p, velocities[i], time)
			pf = shiftPoint(pf, HEIGHT_DEF, WIDTH_DEF)
			currPoints[pf] = true
		}

		for y := 0; y < HEIGHT_DEF; y++ {
			for x := 0; x < WIDTH_DEF; x++ {
				if currPoints[coord{x, y}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		time++
	}
}

func main() {
	data, err := os.ReadFile("./input14_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	// start = time.Now()
	part2(lines)
	// fmt.Printf("part2: %s\n", time.Since(start))
}
