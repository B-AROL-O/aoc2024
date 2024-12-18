package main

import (
	"aoc18/graph"
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

func checkInBounds(pos Pos, rows int, cols int) bool {
	return pos.r >= 0 && pos.r < rows && pos.c >= 0 && pos.c < cols
}

var dirs []Pos = []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // E, W, S, N

func printGrid(obstacles map[Pos]bool, rows int, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if obstacles[Pos{r, c}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func addGraphEdges(graph *graph.Graph[Pos], obstacles map[Pos]bool, rows int, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, d := range dirs {
				p := Pos{r + d.r, c + d.c}
				if checkInBounds(p, rows, cols) && !obstacles[p] {
					graph.AddEdge(Pos{r, c}, p, 1)
				}
			}
		}
	}
}

func addGraphNodes(g *graph.Graph[Pos], obstacles map[Pos]bool, rows int, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !obstacles[Pos{r, c}] {
				g.AddNode(Pos{r, c})
			}
		}
	}
}
func getGraph(obstacles map[Pos]bool, rows int, cols int) (*graph.Graph[Pos], Pos, Pos) {
	g := graph.Graph[Pos]{}
	g.Init(rows*cols, rows*cols)
	start, end := Pos{0, 0}, Pos{rows - 1, cols - 1}
	addGraphNodes(&g, obstacles, rows, cols)
	addGraphEdges(&g, obstacles, rows, cols)
	return &g, start, end
}

func getObstacles(bytes []string) map[Pos]bool {
	obstacles := map[Pos]bool{}
	for _, s := range bytes {
		p := Pos{}
		fmt.Sscanf(s, "%d,%d", &p.r, &p.c)
		obstacles[p] = true
	}
	return obstacles
}

func solve(bytes []string, numFixedBytes int, size int) {
	obstacles := getObstacles(bytes[:numFixedBytes])
	g, start, end := getGraph(obstacles, size, size)
	// printGrid(obstacles, size, size)
	min := g.ShortestPath(start, end)
	fmt.Println(min)
}

func part1(bytes []string) {
	solve(bytes, 1024, 71)
	// solve(bytes, 12, 7)
}

func part2(bytes []string) {
	fixed, size := 2900, 71
	obstacles := getObstacles(bytes[:fixed])
	g, start, end := getGraph(obstacles, size, size)

	i := 0
	for fixed+i < len(bytes) {
		fmt.Println(i)
		newObstacle := Pos{}
		fmt.Sscanf(bytes[fixed+i], "%d,%d", &newObstacle.r, &newObstacle.c)
		obstacles[newObstacle] = true
		g.RemoveNode(newObstacle)
		if g.ShortestPath(start, end) == math.MaxInt {
			fmt.Printf("%d,%d\n", newObstacle.r, newObstacle.c)
			break
		}
		i++
	}
}

func main() {
	data, err := os.ReadFile("./input18_def.txt")
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
