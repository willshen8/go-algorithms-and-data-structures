package graph

import "sync"

type ID interface{}

// Vertice represent a node on the graph
type Vertex struct {
	value interface{}
}

// Graph is constructed using adjacency list method with a map
// This implementation is undirected with no weight
// All the vertices are stored in a pointer array
// Edge is defined a vertex leading to other vertices
type Graph struct {
	vertices []*Vertex
	edges    map[Vertex][]*Vertex
	lock     sync.RWMutex
}

// NewGraph is a constructor that initializes a new graph
func NewGraph() *Graph {
	newGraph := &Graph{}
	newGraph.edges = make(map[Vertex][]*Vertex)
	return newGraph
}

// AddVertex will add a new vertex
func (g *Graph) AddVertex(val interface{}) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.vertices = append(g.vertices, &Vertex{value: val})
}

// AddEdge will add a directed edge from A to B
func (g *Graph) AddEdge(a *Vertex, b *Vertex) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.edges[*a] = append(g.edges[*a], b)
}

func (g *Graph) RemoveVertex(v *Vertex) {
	g.lock.Lock()
	defer g.lock.Unlock()
	delete(g.edges, *v) // remove it's edges
	for index, vertex := range g.vertices {
		if vertex == v {
			g.vertices[index] = g.vertices[len(g.vertices)-1] // move the last element to i
			g.vertices = g.vertices[:len(g.vertices)-1]       // truncate the slice to remove v
		}
	}
}

func (g *Graph) RemoveEdge(a *Vertex, b *Vertex) bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	if len(g.edges[*a]) == 0 {
		return false
	}
	for index, vertex := range g.edges[*a] {
		if vertex == b {
			g.edges[*a][index] = g.edges[*a][len(g.edges)-1]
			g.edges[*a] = g.edges[*a][:len(g.edges)-1]
		}
	}
	return true
}
