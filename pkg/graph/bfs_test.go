package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSWithLinearGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	vertexA, vertexB, vertexC := graph.vertices[0], graph.vertices[1], graph.vertices[2]
	graph.AddEdge(vertexA, vertexB)
	graph.AddEdge(vertexB, vertexC)
	expected := []Vertex{*vertexA, *vertexB, *vertexC}
	actual := graph.BFS(*vertexA)
	assert.Equal(t, expected, actual)
}

func TestBFSWithTree(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("Grandparent")
	graph.AddVertex("Father")
	graph.AddVertex("Uncle")
	graph.AddVertex("Me")
	graph.AddVertex("Cousin")
	GrandparentVertex, FatherVertex, UncleVertex := graph.vertices[0], graph.vertices[1], graph.vertices[2]
	MeVertex, CousinVertex := graph.vertices[3], graph.vertices[4]
	graph.AddEdge(GrandparentVertex, FatherVertex)
	graph.AddEdge(GrandparentVertex, UncleVertex)
	graph.AddEdge(FatherVertex, MeVertex)
	graph.AddEdge(UncleVertex, CousinVertex)
	expected := []Vertex{*GrandparentVertex, *FatherVertex, *UncleVertex, *MeVertex, *CousinVertex}
	actual := graph.BFS(*GrandparentVertex)
	assert.Equal(t, expected, actual)
}

func TestBFSWithACycle(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	vertex1, vertex2, vertex3 := graph.vertices[0], graph.vertices[1], graph.vertices[2]
	graph.AddEdge(vertex1, vertex2)
	graph.AddEdge(vertex2, vertex3)
	graph.AddEdge(vertex3, vertex1)
	expected := []Vertex{*vertex1, *vertex2, *vertex3}
	actual := graph.BFS(*vertex1)
	assert.Equal(t, expected, actual)
}
