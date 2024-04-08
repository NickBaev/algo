package contraction

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Vertex struct {
	Label      int
	Neighbours []*Vertex
}

type Graph struct {
	MaxKey         int
	VerticesLabels []int
	Vertices       map[int]*Vertex
}

func ReadGraphFromFile(path string) Graph {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = '\t'
	csvReader.FieldsPerRecord = -1
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	graph := Graph{
		MaxKey:   0,
		Vertices: make(map[int]*Vertex),
	}
	for _, line := range data {
		var label int
		for i, field := range line {
			fieldInt, _ := strconv.Atoi(field)
			if i == 0 {
				label = fieldInt
				_, ok := graph.Vertices[label]
				if !ok {
					vertex := Vertex{label, make([]*Vertex, 0)}
					graph.Vertices[label] = &vertex
					graph.VerticesLabels = append(graph.VerticesLabels, label)
				}

				continue
			}

			_, ok := graph.Vertices[fieldInt]
			if !ok {
				vertex := Vertex{fieldInt, make([]*Vertex, 0)}
				graph.Vertices[fieldInt] = &vertex
				graph.VerticesLabels = append(graph.VerticesLabels, fieldInt)
			}
			graph.Vertices[label].Neighbours = append(graph.Vertices[label].Neighbours, graph.Vertices[fieldInt])
		}
	}
	graph.MaxKey = len(graph.Vertices)

	return graph
}

func (g *Graph) GetRandomEdge() (*Vertex, *Vertex) {
	label := g.VerticesLabels[rand.Intn(len(g.VerticesLabels)-1)]
	head := g.Vertices[label]
	tail := g.Vertices[label].Neighbours[rand.Intn(len(head.Neighbours)-1)]

	return head, tail
}

func (g *Graph) MergeVertices(head, tail *Vertex) {
	//Merge neighbours
	neighbours := make([]*Vertex, 0)
	neighbours = append(neighbours, head.Neighbours...)
	neighbours = append(neighbours, tail.Neighbours...)

	//Delete self loops
	neighbours = DeleteSelfLoops(neighbours, head.Label)
	neighbours = DeleteSelfLoops(neighbours, tail.Label)

	//Create new Vertex
	g.MaxKey++
	newVertexLabel := g.MaxKey
	newVertex := Vertex{newVertexLabel, neighbours}
	g.Vertices[newVertexLabel] = &newVertex
	g.VerticesLabels = append(g.VerticesLabels, newVertexLabel)

	//Reconnect new Vertex to neighbours
	for _, neighbour := range neighbours {
		g.Vertices[neighbour.Label].Neighbours = append(g.Vertices[neighbour.Label].Neighbours, &newVertex)
	}

	g.DeleteVertex(head)
	g.DeleteVertex(tail)

}

func DeleteSelfLoops(neighbours []*Vertex, label int) []*Vertex {
	newNeighbours := make([]*Vertex, 0)
	for _, val := range neighbours {
		if val.Label != label {
			newNeighbours = append(newNeighbours, val)
		}
	}

	return neighbours
}

func (g *Graph) DeleteVertex(vertex *Vertex) {
	//Delete vertex references from neighbours
	for _, neighbor := range vertex.Neighbours {
		newNeighbors := make([]*Vertex, 0)
		for _, val := range neighbor.Neighbours {
			if val.Label == vertex.Label {
				continue
			}
			newNeighbors = append(newNeighbors, val)
		}
		g.Vertices[neighbor.Label].Neighbours = newNeighbors
	}

	//Delete Vertex
	g.Vertices[vertex.Label] = nil
	for i, label := range g.VerticesLabels {
		if label == vertex.Label {
			g.VerticesLabels[i] = g.VerticesLabels[len(g.VerticesLabels)-1]
			g.VerticesLabels = g.VerticesLabels[:len(g.VerticesLabels)-1]
		}
	}
}

func (g *Graph) Contract() {
	for len(g.VerticesLabels) > 2 {
		g.MergeVertices(g.GetRandomEdge())
	}
}
