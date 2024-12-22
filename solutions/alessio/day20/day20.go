package main

import (
	"aoc20/graph"
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

func printPath(path map[Pos]bool, obstacles map[Pos]bool, rows int, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if path[Pos{r, c}] && obstacles[Pos{r, c}] {
				panic("anomaly")
			}
			if obstacles[Pos{r, c}] {
				fmt.Print("#")
			} else if path[Pos{r, c}] {
				fmt.Print("O")
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
			p := Pos{r, c}
			if obstacles[p] {
				continue
			}
			for _, d := range dirs {
				adj := Pos{r + d.r, c + d.c}
				if checkInBounds(adj, rows, cols) && !obstacles[adj] {
					graph.AddEdge(p, adj, 1)
					graph.AddEdge(adj, p, 1) // undirected graph
				}
			}
		}
	}
}

func addGraphNodes(g *graph.Graph[Pos], obstacles map[Pos]bool, grid []string, rows int, cols int) (Pos, Pos) {
	var start, end Pos
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !obstacles[Pos{r, c}] {
				g.AddNode(Pos{r, c})
				if grid[r][c] == 'S' {
					start = Pos{r, c}
				} else if grid[r][c] == 'E' {
					end = Pos{r, c}
				}
			}
		}
	}
	return start, end
}
func getGraph(grid []string, obstacles map[Pos]bool, rows int, cols int) (*graph.Graph[Pos], Pos, Pos) {
	g := graph.Graph[Pos]{}
	g.Init(rows*cols, rows*cols)
	start, end := addGraphNodes(&g, obstacles, grid, rows, cols)
	addGraphEdges(&g, obstacles, rows, cols)
	return &g, start, end
}

func getObstacles(grid []string) map[Pos]bool {
	obstacles := map[Pos]bool{}
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '#' {
				obstacles[Pos{r, c}] = true
			}
		}
	}
	return obstacles
}

func isOnGridEdge(pos Pos, rows int, cols int) bool {
	return pos.r == 0 || pos.r == rows-1 || pos.c == 0 || pos.c == cols-1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(grid []string) {
	rows, cols := len(grid), len(grid[0])
	obstacles := getObstacles(grid)
	g, start, end := getGraph(grid, obstacles, len(grid), len(grid[0]))
	// printGrid(obstacles, len(grid), len(grid[0]))
	// path := g.GetOnlyPath(start, end)
	_, _, orderedPath := g.ShortestPath(start, end)
	// printPath(path, obstacles, rows, cols)
	pathIndexes := map[Pos]int{}
	for i, p := range orderedPath {
		pathIndexes[p] = i
	}

	cnt := 0
	saveCheat := map[int]int{}
	for o := range obstacles {
		if isOnGridEdge(o, rows, cols) {
			continue
		}
		from := Pos{}
		to := Pos{}
		if grid[o.r+1][o.c] != '#' && grid[o.r-1][o.c] != '#' {
			from.r = o.r + 1
			from.c = o.c
			to.r = o.r - 1
			to.c = o.c
		} else if grid[o.r][o.c+1] != '#' && grid[o.r][o.c-1] != '#' {
			from.r = o.r
			from.c = o.c + 1
			to.r = o.r
			to.c = o.c - 1
		} else {
			continue
		}

		origSubpathLen := abs(pathIndexes[to] - pathIndexes[from])

		obstacles[o] = false
		g.AddNode(o)
		g.AddEdge(from, o, 1)
		g.AddEdge(o, from, 1)
		g.AddEdge(to, o, 1)
		g.AddEdge(o, to, 1)
		// fmt.Printf("\nRemoving obstacle %d,%d\n", o.r, o.c)
		saving := origSubpathLen - 2
		// fmt.Printf("Saved %d picoseconds.\n", saving)
		saveCheat[saving]++
		if saving >= 100 {
			// fmt.Println(saving)
			cnt++
		}
		obstacles[o] = true
		g.RemoveNode(o)
	}

	// saves := []int{}
	// for k := range saveCheat {
	// 	saves = append(saves, k)
	// }

	// sort.Ints(saves)

	// for _, s := range saves {
	// 	fmt.Printf("There are %d cheats that save %d picoseconds.\n", saveCheat[s], s)
	// }
	fmt.Println(cnt)
}

func part1(grid []string) {
	solve(grid)
}

func part2(grid []string) {
}

func main() {
	data, err := os.ReadFile("./input20_def.txt")
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
