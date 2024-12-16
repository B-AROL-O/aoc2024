package graph

import "slices"

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
	To     *Node[T]
}

func (graph *Graph[T]) AddNode(val T) *Node[T] {
	node := Node[T]{Val: val}
	graph.Nodes = append(graph.Nodes, node)
	return &node
}

func (graph *Graph[T]) AddEdge(from T, to T, weight int) *Edge[T] {
	var toNode *Node[T]

	for _, n := range graph.Nodes {
		if n.Val == to {
			toNode = &n
			break
		}
	}

	edge := Edge[T]{Weight: weight, To: toNode}
	graph.Edges = append(graph.Edges, edge)
	graph.AdjList[from] = append(graph.AdjList[from], edge)
	return &edge
}

func (graph *Graph[T]) HasNode(val T) bool {
	return slices.ContainsFunc(graph.Nodes, func(n Node[T]) bool { return n.Val == val })
}

func (graph *Graph[T]) Init(nodes int, edges int) {
	graph.Nodes = make([]Node[T], 0, nodes)
	graph.Edges = make([]Edge[T], 0, edges)
	graph.AdjList = make(map[T][]Edge[T], nodes)
}
