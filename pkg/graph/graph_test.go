package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmptyGraph(t *testing.T) {
	actual := NewGraph()
	expectedVerticesLength := 0
	expectedEdgeLength := 0
	assert.Equal(t, expectedVerticesLength, len(actual.vertices))
	assert.Equal(t, expectedEdgeLength, len(actual.edges))
}

func TestAddVertex(t *testing.T) {
	newGraph := &Graph{}
	newGraph.AddVertex("A")
	expected := "A"
	actual := newGraph.vertices[0]
	assert.Equal(t, expected, actual.value)
}

func TestAddEdge(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	vertexA, vertexB := graph.vertices[0], graph.vertices[1]
	graph.AddEdge(vertexA, vertexB)
	expectedVertexFromA := &Vertex{value: "B"}
	actual := graph.edges[Vertex{value: "A"}][0]
	assert.Equal(t, expectedVertexFromA, actual)
}
