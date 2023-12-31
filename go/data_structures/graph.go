// graph.go

package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{}
	}
}

func (g *Graph) AddEdge(vertex1, vertex2 int) {
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

func (g *Graph) Display() {
	for vertex, neighbors := range g.vertices {
		fmt.Printf("%d -> %v\n", vertex, neighbors)
	}
}

func main() {
	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1)

	fmt.Println("Graph:")
	graph.Display()
}
