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
