package main

import (
	"aoc23/graph"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildGraph(lines []string) *graph.Graph[string] {
	g := graph.Graph[string]{}
	g.Init(len(lines)*2, len(lines)*2)

	for _, l := range lines {
		nodes := strings.Split(l, "-")
		if !g.HasNode(nodes[0]) {
			g.AddNode(nodes[0])
		}
		if !g.HasNode(nodes[1]) {
			g.AddNode(nodes[1])
		}

		g.AddUndirectedEdge(nodes[0], nodes[1], 1)
	}

	return &g
}

func setToString(nodes []string) string {
	sort.Strings(nodes)
	return strings.Join(nodes, "-")
}

func solve(lines []string) {
	g := buildGraph(lines)
	nodes := g.GetNodesVals()
	distinct := map[string]bool{}
	for _, n := range nodes {
		if n[0] != 't' {
			continue
		}
		adj := []string{}
		for _, e := range g.AdjList[n] {
			adj = append(adj, e.To.Val)
		}

		for i := 0; i < len(adj); i++ {
			for j := i + 1; j < len(adj); j++ {
				if g.HasEdge(adj[i], adj[j]) || g.HasEdge(adj[j], adj[i]) {
					distinct[setToString([]string{n, adj[i], adj[j]})] = true
				}
			}
		}
	}
	fmt.Println(len(distinct))
}

func part1(lines []string) {
	solve(lines)
}

func part2(lines []string) {
	// solve(lines)
}

func main() {
	data, err := os.ReadFile("./input23_def.txt")
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
