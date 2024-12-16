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

func nextDirs(curr Pos) []Pos {
	left := Pos{-curr.c, curr.r}
	right := Pos{curr.c, -curr.r}
	return []Pos{left, right}
}

var dirs []Pos = []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // E, W, S, N

func checkInBounds(pos Pos, rows int, cols int) bool {
	return pos.r >= 0 && pos.r < rows && pos.c >= 0 && pos.c < cols
}

type PosDir struct {
	pos, dir Pos
}

func getDirSym(dir Pos) byte {
	if dir == dirs[0] {
		return '>'
	} else if dir == dirs[1] {
		return '<'
	} else if dir == dirs[2] {
		return 'v'
	} else {
		return '^'
	}
}

func printPath(grid []string, path []PosDir) {
	pathmap := make(map[Pos]Pos)
	for _, p := range path {
		if _, exists := pathmap[p.pos]; exists {
			fmt.Println("path is not unique")
		}
		pathmap[p.pos] = p.dir
	}
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dir, exists := pathmap[Pos{r, c}]; exists {
				fmt.Printf("%c", getDirSym(dir))
			} else {
				fmt.Printf("%c", grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findPathRec(grid []string, currPos Pos, currDir Pos, currScore int, currMin *int, visited *map[Pos]bool, path []PosDir) {
	if grid[currPos.r][currPos.c] == '#' {
		return
	}
	if grid[currPos.r][currPos.c] == 'E' { // reached end of path
		if currScore < *currMin {
			fmt.Printf("path with score %d\n", currScore)
			// printPath(grid, path)
			*currMin = currScore
		}
		return
	}
	(*visited)[currPos] = true
	// go along currDir
	nextPos := Pos{currPos.r + currDir.r, currPos.c + currDir.c}
	if !(*visited)[nextPos] && currScore+1 < *currMin {
		findPathRec(grid, nextPos, currDir, currScore+1, currMin, visited, append(path, PosDir{currPos, currDir}))
	}
	// turn clockwise and counter-clockwise
	for _, dir := range nextDirs(currDir) {
		nextPos = Pos{currPos.r + dir.r, currPos.c + dir.c}
		if !(*visited)[nextPos] && currScore+1001 < *currMin {
			findPathRec(grid, nextPos, dir, currScore+1001, currMin, visited, append(path, PosDir{currPos, dir}))
		}
	}
	(*visited)[currPos] = false // backtrack
}

func isIntersection(grid []string, pos Pos, currDir Pos) bool {
	for _, dir := range nextDirs(currDir) {
		next := Pos{pos.r + dir.r, pos.c + dir.c}
		if grid[next.r][next.c] == '.' {
			return true
		}
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

func buildGraphRec(grid []string, currPos Pos, currDir Pos, lastNode Pos, currWeight int, graph *graph.Graph[Pos]) {
	if grid[currPos.r][currPos.c] == '#' {
		return
	}

	visited := false
	if isIntersection(grid, currPos, currDir) {
		visited = graph.HasNode(currPos)
		if !visited {
			graph.AddNode(currPos)
		}
		if currPos != lastNode {
			graph.AddEdge(lastNode, currPos, currWeight)
		}
		if grid[currPos.r][currPos.c] == 'E' {
			return
		}
		if !visited || grid[currPos.r][currPos.c] == 'S' { // if already visited, no sense recurring on it.
			// restart counting from this node
			// turn clockwise and counter-clockwise
			for _, dir := range nextDirs(currDir) {
				nextPos := Pos{currPos.r + dir.r, currPos.c + dir.c}
				// lastNode = currPos and weight = 1001 (restart counting from this node)
				buildGraphRec(grid, nextPos, dir, currPos, 1001, graph)
			}
		}
	}

	// go along currDir
	nextPos := Pos{currPos.r + currDir.r, currPos.c + currDir.c}
	buildGraphRec(grid, nextPos, currDir, lastNode, currWeight+1, graph)
}

func getGraph(grid []string, rows int, cols int) *graph.Graph[Pos] {
	g := graph.Graph[Pos]{}
	g.Init(rows*cols, rows*cols)
	start, end := findStartAndEndPos(grid, rows, cols)
	g.AddNode(start)
	g.AddNode(end)
	buildGraphRec(grid, start, dirs[0], start, 0, &g)
	return &g
}

func solve1(grid []string) {
	rows, cols := len(grid), len(grid[0])
	g := getGraph(grid, rows, cols)
	printMazeWithNodes(grid, rows, cols, g)
	// fmt.Println(g.Nodes)

	fmt.Println(g.AdjList)
	// currPos, endPos := findStartAndEndPos(grid, rows, cols)
	// fmt.Println(currPos, endPos)
	// currDir := dirs[0] // start facing east

	// currMin := math.MaxInt
	// visited := make(map[Pos]bool)
	// findPathRec(grid, currPos, currDir, 0, &currMin, &visited, []PosDir{})
	// fmt.Println(currMin)
}

func solve2(origGrid []string) {

}

func part1(grid []string) {
	solve1(grid)
}

func part2(grid []string) {
	solve2(grid)
}

func main() {
	data, err := os.ReadFile("./input16.txt")
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
