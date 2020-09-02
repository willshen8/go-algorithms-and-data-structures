package graph

import (
	"github.com/willshen8/go-algorithms-and-data-structures/pkg/queue"
)

//BFS = breath first search. It will take a start node and returns a list of nodes visited
func (g *Graph) BFS(start Vertex) []Vertex {
	var visited []Vertex
	q := queue.NewQueue()
	q.Enqueue(start)
	visited = append(visited, start)

	for q.Length() > 0 {
		qVertex := q.Dequeue()
		for _, vertex := range g.edges[qVertex.(Vertex)] {
			var exist bool
			for _, v := range visited {
				if v == *vertex {
					exist = true
				}
			}
			if !exist {
				visited = append(visited, *vertex)
				q.Enqueue(*vertex)
			}
		}
	}
	return visited
}
