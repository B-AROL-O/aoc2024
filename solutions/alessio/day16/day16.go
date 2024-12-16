package main

import (
	"aoc16/graph"
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

func findStartAndEndPos(grid []string, rows int, cols int) (Pos, Pos) {
	start, end := Pos{-1, -1}, Pos{-1, -1}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				start = Pos{r, c}
			} else if grid[r][c] == 'E' {
				end = Pos{r, c}
			}
			if start.r != -1 && end.r != -1 {
				return start, end
			}
		}
	}
	return start, end
}

func turnDirs(curr Pos) []Pos {
	left := Pos{-curr.c, curr.r}
	right := Pos{curr.c, -curr.r}
	return []Pos{left, right}
}

var dirs []Pos = []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // E, W, S, N

type PosDir struct {
	pos, dir Pos
}

func isIntersection(grid []string, pos Pos) bool {
	e := grid[pos.r+dirs[0].r][pos.c+dirs[0].c]
	w := grid[pos.r+dirs[1].r][pos.c+dirs[1].c]
	s := grid[pos.r+dirs[2].r][pos.c+dirs[2].c]
	n := grid[pos.r+dirs[3].r][pos.c+dirs[3].c]

	if (n != '#' || s != '#') && (e != '#' || w != '#') {
		return true
	}

	return false
}

func printMazeWithNodes(grid []string, rows int, cols int, graph *graph.Graph[Pos]) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if graph.HasNode(Pos{r, c}) {
				fmt.Print("+")
			} else {
				fmt.Printf("%c", grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func addGraphEdgesIter(grid []string, start Pos, graph *graph.Graph[Pos]) {
	queue := []PosDir{{start, dirs[0]}}
	visited := make(map[PosDir]bool)

	if grid[start.r-1][start.c] != '#' {
		queue = append(queue, PosDir{start, dirs[3]})
	}

	if grid[start.r+1][start.c] != '#' {
		queue = append(queue, PosDir{start, dirs[2]})
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		nextPos := Pos{curr.pos.r + curr.dir.r, curr.pos.c + curr.dir.c}
		dist := 0
		for grid[nextPos.r][nextPos.c] != '#' {
			dist++
			if isIntersection(grid, nextPos) || grid[nextPos.r][nextPos.c] == 'E' {
				if !graph.HasEdge(curr.pos, nextPos) {
					graph.AddEdge(curr.pos, nextPos, dist, curr.dir)
				}
				for _, dir := range turnDirs(curr.dir) {
					if !visited[PosDir{nextPos, dir}] {
						queue = append(queue, PosDir{nextPos, dir})
					}
				}
				queue = append(queue, PosDir{nextPos, curr.dir})
				break
			}
			nextPos.r += curr.dir.r
			nextPos.c += curr.dir.c
		}
	}
}

func addGraphNodes(g *graph.Graph[Pos], grid []string, rows int, cols int) {
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] != '#' && (isIntersection(grid, Pos{r, c}) || grid[r][c] == 'E' || grid[r][c] == 'S') {
				g.AddNode(Pos{r, c})
			}
		}
	}
}
func getGraph(grid []string, rows int, cols int) (*graph.Graph[Pos], Pos, Pos) {
	g := graph.Graph[Pos]{}
	g.Init(rows*cols, rows*cols)
	start, end := findStartAndEndPos(grid, rows, cols)
	addGraphNodes(&g, grid, rows, cols)
	addGraphEdgesIter(grid, start, &g)
	return &g, start, end
}

func solve1(grid []string) {
	rows, cols := len(grid), len(grid[0])
	g, start, end := getGraph(grid, rows, cols)
	// printMazeWithNodes(grid, rows, cols, g)
	fmt.Println(g.ShortestPath(start, end, dirs[0]))
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func getDir(from Pos, to Pos) Pos {
	diffR := to.r - from.r
	diffC := to.c - from.c
	if diffR == 0 {
		dirC := diffC / abs(diffC)
		return Pos{0, dirC}
	} else if diffC == 0 {
		dirR := diffR / abs(diffR)
		return Pos{dirR, 0}
	} else {
		fmt.Println("wtf")
		return Pos{0, 0}
	}
}

func solve2(grid []string) {
	rows, cols := len(grid), len(grid[0])
	g, start, end := getGraph(grid, rows, cols)
	edges := g.GetFullPathsEdges(start, end, dirs[0])

	tiles := map[Pos]bool{}
	for _, e := range edges {
		dir := getDir(e.From, e.To.Val)
		curr := Pos{e.From.r, e.From.c}
		for curr.r != e.To.Val.r || curr.c != e.To.Val.c {
			tiles[curr] = true
			curr.r += dir.r
			curr.c += dir.c
		}
		tiles[curr] = true
	}

	// for r := 0; r < rows; r++ {
	// 	for c := 0; c < cols; c++ {
	// 		if tiles[Pos{r, c}] {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Printf("%c", grid[r][c])
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
	fmt.Println(len(tiles))
}

func part1(grid []string) {
	solve1(grid)
}

func part2(grid []string) {
	solve2(grid)
}

func main() {
	data, err := os.ReadFile("./input16_def.txt")
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
