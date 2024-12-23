package main

import (
	"fmt"
	"math"
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

type Pos struct {
	r, c int
}

var dirSymbols map[Pos]string = map[Pos]string{
	{0, 1}:  ">",
	{0, -1}: "<",
	{1, 0}:  "v",
	{-1, 0}: "^",
}

type PosPath struct {
	pos  Pos
	path string
}

var numPad map[byte]Pos = map[byte]Pos{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'-': {3, 0},
	'0': {3, 1},
	'A': {3, 2},
}

var dirPad map[byte]Pos = map[byte]Pos{
	'-': {0, 0},
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

func getNextDirs(from Pos, to Pos) []Pos {
	nextDirs := []Pos{}
	if to.r > from.r {
		nextDirs = append(nextDirs, Pos{1, 0})
	}
	if to.r < from.r {
		nextDirs = append(nextDirs, Pos{-1, 0})
	}
	if to.c > from.c {
		nextDirs = append(nextDirs, Pos{0, 1})
	}
	if to.c < from.c {
		nextDirs = append(nextDirs, Pos{0, -1})
	}
	return nextDirs
}

func getPadPaths(start byte, end byte, byteToPos map[byte]Pos) []string {
	posToByte := map[Pos]byte{}
	for k, v := range byteToPos {
		posToByte[v] = k
	}
	paths := []string{}
	startPosPath := PosPath{byteToPos[start], ""}
	q := []PosPath{startPosPath}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.pos == byteToPos[end] {
			paths = append(paths, curr.path+"A")
			continue
		}

		for _, dir := range getNextDirs(curr.pos, byteToPos[end]) {
			nextPos := Pos{curr.pos.r + dir.r, curr.pos.c + dir.c}
			dirSym := dirSymbols[dir]
			path := curr.path + dirSym
			nextPosPath := PosPath{nextPos, path}
			if posToByte[nextPos] != '-' && posToByte[nextPos] != start {
				q = append(q, nextPosPath)
			}
		}
	}
	return paths
}

type MemoKey struct {
	path string
	lvl  int
}

func getShortestPath(seq string, lvl int, maxLvl int, memo *map[MemoKey]int, useNumPad bool) int {
	if lvl == maxLvl {
		return len(seq)
	}

	if val, ok := (*memo)[MemoKey{seq, lvl}]; ok {
		return val
	}

	start := 'A'
	totMinLen := 0
	pad := dirPad
	if useNumPad {
		pad = numPad
	}
	for _, c := range seq {
		paths := getPadPaths(byte(start), byte(c), pad)
		min := math.MaxInt
		for _, p := range paths {
			l := getShortestPath(p, lvl+1, maxLvl, memo, false)
			if l < min {
				min = l
			}
		}
		totMinLen += min
		start = c
	}
	(*memo)[MemoKey{seq, lvl}] = totMinLen
	return totMinLen
}

func solve(lines []string, levels int) {
	sum := 0
	memo := map[MemoKey]int{}
	for _, l := range lines {
		totMinLen := getShortestPath(l, -1, levels, &memo, true)
		numPart, _ := strconv.Atoi(l[:len(l)-1])
		fmt.Printf("%d * %d\n", totMinLen, numPart)
		sum += totMinLen * numPart
	}

	fmt.Println(sum)
}

func part1(lines []string) {
	solve(lines, 2)
}

func part2(lines []string) {
	solve(lines, 25)
}

func main() {
	data, err := os.ReadFile("./input21_def.txt")
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
