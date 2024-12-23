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
}

func (graph *Graph[T]) AddNode(val T) *Node[T] {
	node := Node[T]{Val: val}
	graph.Nodes = append(graph.Nodes, node)
	return &node
}

func (graph *Graph[T]) RemoveNode(val T) {
	delete(graph.AdjList, val)
	for i, lst := range graph.AdjList {
		idx := -1
		for j, e := range lst {
			if e.To.Val == val {
				idx = j
				break
			}
		}
		if idx != -1 {
			graph.AdjList[i] = slices.Concat(lst[:idx], lst[idx+1:])
		}
	}
}

func (graph *Graph[T]) AddEdge(from T, to T, weight int) *Edge[T] {
	toNode := graph.GetNode(to)
	if toNode == nil {
		fmt.Println("Error: toNode is nil")
		return nil
	}
	edge := Edge[T]{Weight: weight, To: toNode}
	graph.Edges = append(graph.Edges, edge)
	graph.AdjList[from] = append(graph.AdjList[from], edge)
	return &edge
}

func (graph *Graph[T]) AddUndirectedEdge(from T, to T, weight int) {
	graph.AddEdge(from, to, weight)
	graph.AddEdge(to, from, weight)
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

func (graph *Graph[T]) ShortestPath(from T, to T) (int, map[T]bool, []T) {
	start := graph.GetNode(from)
	end := graph.GetNode(to)
	dist := make(map[T]int, len(graph.Nodes))
	path := map[T]bool{}
	orderedPath := []T{}
	nodes := make([]T, 0, len(graph.Nodes))
	for _, n := range graph.Nodes {
		dist[n.Val] = math.MaxInt
		nodes = append(nodes, n.Val)
	}

	dist[start.Val] = 0
	for len(path) < len(nodes) {
		curr := graph.getClosest(nodes, dist, path)
		if path[curr] {
			continue
		}

		path[curr] = true
		orderedPath = append(orderedPath, curr)

		if curr == end.Val {
			break
		}

		for _, e := range graph.AdjList[curr] {
			// relaxation on the edge
			if !path[e.To.Val] && dist[curr]+e.Weight < dist[e.To.Val] {
				dist[e.To.Val] = dist[curr] + e.Weight
			}
		}
	}
	return dist[end.Val], path, orderedPath
}

func (graph *Graph[T]) GetOnlyPath(from T, to T) (map[T]bool, []T) {
	start := graph.GetNode(from)
	end := graph.GetNode(to)
	path := map[T]bool{}
	orderedPath := []T{}

	curr := start.Val
	for curr != end.Val {
		path[curr] = true
		orderedPath = append(orderedPath, curr)
		numAdj := len(graph.AdjList[curr])
		if numAdj > 2 {
			panic("too many adjacents")
		}
		for _, e := range graph.AdjList[curr] {
			if !path[e.To.Val] {
				curr = e.To.Val
				break
			}
		}
	}
	path[curr] = true
	orderedPath = append(orderedPath, curr)
	return path, orderedPath
}

func (graph *Graph[T]) getClosest(nodes []T, dist map[T]int, visited map[T]bool) T {
	var minNode T
	minDist := math.MaxInt

	for _, n := range nodes {
		if !visited[n] && dist[n] <= minDist {
			minDist = dist[n]
			minNode = n
		}
	}
	return minNode
}

func (graph *Graph[T]) GetFullPathsEdges(start T, end T) []Edge[T] {
	allEdges := []Edge[T]{}
	min := math.MaxInt
	graph.getBestPathsRec(start, &map[T]bool{}, []Edge[T]{}, &allEdges, end, 0, &min)
	return allEdges
}

func (graph *Graph[T]) getBestPathsRec(currNode T, visited *map[T]bool, currPath []Edge[T], edges *[]Edge[T], end T, currVal int, currMin *int) {
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
			if currVal+e.Weight <= *currMin {
				graph.getBestPathsRec(next, visited, append(currPath, e), edges, end, currVal+e.Weight, currMin)
			}
		}
	}
	(*visited)[currNode] = false
}

func (graph *Graph[T]) GetNodesVals() []T {
	vals := make([]T, 0, len(graph.Nodes))
	for _, n := range graph.Nodes {
		vals = append(vals, n.Val)
	}
	return vals
}

func (graph *Graph[T]) GetConnectedComponents() [][]T {
	visited := map[T]bool{}
	components := make([][]T, 0, len(graph.Nodes))

	for _, node := range graph.Nodes {
		if visited[node.Val] {
			continue
		}
		visitedCC := map[T]bool{}
		graph.dfs(node.Val, &visitedCC)

		component := []T{}
		for n := range visitedCC {
			if visited[n] {
				fmt.Println("already visited, error!")
			}
			visited[n] = true
			component = append(component, n)
		}
		components = append(components, component)
	}
	return components
}

func (graph *Graph[T]) dfs(curr T, visited *map[T]bool) {
	(*visited)[curr] = true
	for _, adj := range graph.AdjList[curr] {
		if !(*visited)[adj.To.Val] {
			graph.dfs(adj.To.Val, visited)
		}
	}
}
