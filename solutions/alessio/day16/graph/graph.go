package graph

import (
	"fmt"
	"math"
	"slices"
)

type Graph[T comparable] struct {
	Nodes   []Node[T]
	Edges   []Edge[T]
	AdjList map[T][]Edge[T]
}

type Node[T any] struct {
	Val T
}

type Edge[T any] struct {
	Weight int
	From   T
	To     *Node[T]
	Dir    T
}

func (graph *Graph[T]) AddNode(val T) *Node[T] {
	node := Node[T]{Val: val}
	graph.Nodes = append(graph.Nodes, node)
	return &node
}

func (graph *Graph[T]) AddEdge(from T, to T, weight int, dir T) *Edge[T] {
	toNode := graph.GetNode(to)
	if toNode == nil {
		fmt.Println("Error: toNode is nil")
		return nil
	}
	edge := Edge[T]{Weight: weight, To: toNode, Dir: dir}
	graph.Edges = append(graph.Edges, edge)
	graph.AdjList[from] = append(graph.AdjList[from], edge)
	return &edge
}

func (graph *Graph[T]) HasNode(val T) bool {
	return slices.ContainsFunc(graph.Nodes, func(n Node[T]) bool { return n.Val == val })
}

func (graph *Graph[T]) HasEdge(from T, to T) bool {
	for _, e := range graph.AdjList[from] {
		if e.To.Val == to {
			return true
		}
	}
	return false
}

func (graph *Graph[T]) GetNode(val T) *Node[T] {
	for _, n := range graph.Nodes {
		if n.Val == val {
			return &n
		}
	}
	return nil
}

func (graph *Graph[T]) Init(nodes int, edges int) {
	graph.Nodes = make([]Node[T], 0, nodes)
	graph.Edges = make([]Edge[T], 0, edges)
	graph.AdjList = make(map[T][]Edge[T], nodes)
}

func (graph *Graph[T]) ShortestPath(from T, to T, startDir T) int {
	start := graph.GetNode(from)
	end := graph.GetNode(to)
	dist := make(map[T]int, len(graph.Nodes))
	minDir := make(map[T]T, len(graph.Nodes))
	path := map[T]bool{}
	nodes := make([]T, 0, len(graph.Nodes))
	for _, n := range graph.Nodes {
		dist[n.Val] = math.MaxInt
		nodes = append(nodes, n.Val)
	}

	dist[start.Val] = 0
	minDir[start.Val] = startDir
	for len(path) < len(nodes) {
		curr := graph.getClosest(nodes, dist, path)
		if path[curr] {
			fmt.Println("Error: curr is already in path")
		}
		path[curr] = true

		if curr == end.Val {
			return dist[curr]
		}

		for _, e := range graph.AdjList[curr] {
			// relaxation on the edge
			weight := e.Weight
			if minDir[curr] != e.Dir {
				weight += 1000
			}
			if dist[curr]+weight < dist[e.To.Val] {
				dist[e.To.Val] = dist[curr] + weight
				minDir[e.To.Val] = e.Dir
			}
		}
	}
	return dist[end.Val]
}

func (graph *Graph[T]) getClosest(nodes []T, dist map[T]int, visited map[T]bool) T {
	var minNode T
	minDist := math.MaxInt

	for _, n := range nodes {
		if !visited[n] && dist[n] < minDist {
			minDist = dist[n]
			minNode = n
		}
	}
	return minNode
}

func (graph *Graph[T]) GetFullPathsEdges(start T, end T, dir T, minFound *int) []Edge[T] {
	allEdges := []Edge[T]{}
	// oldMin := minFound
	graph.getBestPathsRec(start, dir, &map[T]bool{}, []Edge[T]{}, &allEdges, end, 0, minFound)
	// if minFound != oldMin {
	// 	fmt.Println("min changed from", oldMin, "to", minFound, ". Should not happen.")
	// }
	return allEdges
}

func (graph *Graph[T]) getBestPathsRec(currNode T, currDir T, visited *map[T]bool, currPath []Edge[T], edges *[]Edge[T], end T, currVal int, currMin *int) {
	if currNode == end {
		fmt.Println("found path")
		if currVal < *currMin {
			*currMin = currVal
			(*edges) = append([]Edge[T]{}, currPath...) // replace all
		} else if currVal == *currMin {
			(*edges) = append((*edges), currPath...)
		}
		return
	}

	(*visited)[currNode] = true
	for _, e := range graph.AdjList[currNode] {
		next := e.To.Val
		if !(*visited)[next] {
			e.From = currNode
			weight := e.Weight
			if currDir != e.Dir {
				weight += 1000
			}
			if currVal+weight <= *currMin {
				graph.getBestPathsRec(next, e.Dir, visited, append(currPath, e), edges, end, currVal+weight, currMin)
			}
		}
	}
	(*visited)[currNode] = false
}
