package main

import (
	"aoc16/graph"
	"fmt"
	"math"
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

func solve1(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	g, start, end := getGraph(grid, rows, cols)
	// printMazeWithNodes(grid, rows, cols, g)
	min := g.ShortestPath(start, end, dirs[0])
	fmt.Println(min)
	return min
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

type PathPos struct {
	this  PosDir
	prev  PosDir
	score int
}

type PathEdge struct {
	from Pos
	to   Pos
}

type Ugly struct {
	from PosDir
	to   PosDir
}

func BFS(graph *graph.Graph[Pos], start Pos, end Pos) map[Pos]bool {
	startPosDir := PosDir{start, dirs[0]}
	queue := []PathPos{{startPosDir, PosDir{}, 0}}
	minScores := map[PosDir]int{}
	prevs := map[PosDir][]PosDir{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		currPosDir := curr.this
		_, exists := minScores[currPosDir]
		// if score not yet evaluated or the current score is lower than the previous one, update score and previous
		if !exists || curr.score < minScores[currPosDir] {
			minScores[currPosDir] = curr.score
			prevs[currPosDir] = []PosDir{curr.prev}
		} else if curr.score == minScores[currPosDir] {
			prevs[currPosDir] = append(prevs[currPosDir], curr.prev)
		}

		// if reached the end, continue to find potentially better (or equivalent) paths
		if currPosDir.pos == end {
			continue
		}

		// for every adjacent node, if the score for reaching it
		// in a certain direction is lower than the previous one, add it to the queue
		for _, e := range graph.AdjList[currPosDir.pos] {
			next := e.To.Val
			dir := getDir(currPosDir.pos, next)
			weight := e.Weight
			if dir != currPosDir.dir {
				weight += 1000
			}
			nextPosDir := PosDir{next, dir}
			if min, exists := minScores[nextPosDir]; !exists || curr.score+weight <= min {
				queue = append(queue, PathPos{nextPosDir, currPosDir, curr.score + weight})
			}
		}
	}

	fmt.Println("here 1")
	// find all the directions from where the best paths come to the end node
	minFinalScore := math.MaxInt
	minPosDirs := []PosDir{}
	for p, s := range minScores {
		if p.pos == end {
			if s < minFinalScore {
				minFinalScore = s
				minPosDirs = []PosDir{p}
			} else if s == minFinalScore {
				minPosDirs = append(minPosDirs, p)
			}
		}
	}

	fmt.Println("here 2")

	// walk down each path and mark all the visited tiles
	edges := []PathEdge{}
	nodeQueue := []PosDir{}
	visitedEdges := map[Ugly]bool{}
	nodeQueue = append(nodeQueue, minPosDirs...)
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		for _, adj := range prevs[node] {
			if adj.pos.r == 0 && adj.pos.c == 0 {
				//sentinel node, skip
				continue
			}
			e := Ugly{node, adj}
			if visitedEdges[e] {
				continue
			}
			visitedEdges[e] = true

			nodeQueue = append(nodeQueue, adj)
			edges = append(edges, PathEdge{node.pos, adj.pos})
		}
	}

	fmt.Println("here 3")

	visited := map[Pos]bool{}
	for _, e := range edges {
		if e.from == e.to {
			fmt.Println("found self-loop edge, should be end node. Exiting if not true.")
			fmt.Printf("found from node (%d, %d) to node (%d, %d)\n", e.from.r, e.from.c, e.to.r, e.to.c)
			if e.from != end {
				panic("IT WAS NOT END NODE. PANIC.")
			}
			continue
		}
		dir := getDir(e.from, e.to)
		r, c := e.from.r, e.from.c
		for r != e.to.r || c != e.to.c {
			visited[Pos{r, c}] = true
			r += dir.r
			c += dir.c
		}
		visited[e.to] = true
	}

	return visited
}

func solve2(grid []string, minFound int) {
	rows, cols := len(grid), len(grid[0])
	g, start, end := getGraph(grid, rows, cols)
	visited := BFS(g, start, end)

	// tiles := map[Pos]bool{}
	// for _, e := range edges {
	// 	dir := getDir(e.from, e.to)
	// 	curr := Pos{e.from.r, e.from.c}
	// 	for curr.r != e.to.r || curr.c != e.to.c {
	// 		tiles[curr] = true
	// 		curr.r += dir.r
	// 		curr.c += dir.c
	// 	}
	// 	tiles[curr] = true
	// }

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if visited[Pos{r, c}] {
				fmt.Print("O")
			} else {
				fmt.Printf("%c", grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println(len(visited))
}

func part1(grid []string) int {
	return solve1(grid)
}

func part2(grid []string, minFound int) {
	solve2(grid, minFound)
}

func main() {
	data, err := os.ReadFile("./input16_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	min := part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(lines, min)
	fmt.Printf("part2: %s\n", time.Since(start))
}
