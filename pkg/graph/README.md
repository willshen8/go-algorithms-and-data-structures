# Graphs 

## What is a graph

A graph represent a set of nodes/points/vertices that are connected via links/edges, and these edges connected vertices. The basic operations of a graph are the addition and removal of links and vertices.

Some of the different graphs are:

1. Directed graphs (graphs with directions)

    V1 -----> V2

    V1 and V2 are called edges and the link connecting the two is called the edge.
    Example could include people giving gifts.

2. Non-directed graphs (graphs without directions)

    V1 ------ V2 

    Two cities linked together is non-directed, as one can drive from city A to city B and vice versa. A tree is an undirected graph!

3. Connected graphs - given any two vertices, there is an edge connecting the two vertices

4. Non-connected graphs - if a graph doesn't satisfy the condition of connected graph, then it is a non-connected graph.

5. Simple graphs - A simple graph, also called a strict graph (Tutte 1998, p. 2), is an unweighted, undirected graph containing no graph loops or multiple edges.

6. Multi-graphs - a multi-graph is a graph which is permitted to have multiple edges, that is, edges that have the same end nodes. Thus two vertices may be connected by more than one edge.

7. Weight graphs - edges contains a certain arbitrary value such as cost, distance, time, quantity, etc. 

8. Directed Acyclic Graphs - graphs with no cycles.

## Possible graph representations

### Adjacency matrix

  e.g. The cost of going from A to B is 1

  - |  A 	|   B	|   C	|
   |---	|---	|---	|---	|
A  |   0	|   1	|   2	| 
B  |   2	|   0	|   3	| 
C  |   1	|   -1	|   0	|

 
 Pros: 
 1. Space efficient for dense graphs
 2. Edge weight look up is O(1)
 3. Simplest graph representation

 Cons:
 1. Requires O(N^2) spaces
 2. Iteracting over all edges takes O(N^2)

---
### Adjacency List

An adjacency list is a way to represent a graph as a map from nodes to list of edges.
e.g.
    A -> [(B,4), (C,1)]
    B -> [(C,4)]  - B can reach node C at the cost of 4

Pros:
1. Space efficient for sparse graphs
2. Iterating over all edges is efficient

Cons:
1. Less space efficient for denser graphs
2. Edge weight lookup is O(N)
3. Slightly more complex graph representations

---

### Edge list

An edge list is a way to represent a graph simple as an unordered list of edges.

[(C,A,4), (A,C,1)]

This is seldomly used for its lack of structure.

## Why is it useful?

Many real world problems can be represented as graphs, and solved. 
The internet can be represented as a complicated graph whereby each computer is a node. 

## Common Problems

1. Shortest path problem
Given a weighted graph, find the shortest path from node A to node B

Algorithm: BFS, Dijkstra's, Bellman-Ford, Floyd-Warshall, etc.

2. Connectivity
Can node A reach node B

Algorithm: Use union find data structure or any search algorithm

3. Negative cycles
Does the graph have negative cycles? (a path from A to B have negative weight)

Algorithm: Bellman-Ford and Floyd-Warshall

4. Strongly connected components

Self-contained cycles within a directed graph where every vertex in a given cycle can be reac every other vertex in the same cycle.

Algorithm: Tarjan's and Kosaraju's algorithm

5. Travelling Salesman's problem

Given a list of cities and the distances between each pair of cities, what is the shortest possible route that visits each city exactly once and return to the origin city?

Algorithm: Held-Karp, branch and bound and many approximation algorithm.

6. Minimum spanning tree (MST)

A MST is a subset of the edges of a connected, edge-weighted graph that connects all the vertices together, without any cycles and with the minimum possible total edge weight.

Algorithm: Kruskal's, Prim's & Boruvka's algorithm