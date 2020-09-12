package kosaraju

type Vertex struct {
	Value int
	IsExplored bool
	Topo int // topological order
	SCC int // number of strongly connected component
	Edges []Edge
}

type Edge struct {
	Tail, Head *Vertex
}

type Graph map[int]*Vertex

//simple edges representation
type SimpleEdge struct {
	Tail, Head int
}


func Populate(vertexes []int, edges []SimpleEdge) Graph {
	graph := Graph{}
	for _, value := range vertexes {
		graph[value] = &Vertex{
			Value: value,
			IsExplored: false,
		}
	}
	for _, sedge := range edges {
		graph[sedge.Tail].Edges = append(graph[sedge.Tail].Edges, Edge{
			Tail: graph[sedge.Tail],
			Head: graph[sedge.Head],
			})
	}
	return graph
}

func Reverse(edges []SimpleEdge) []SimpleEdge {
	var result []SimpleEdge
	for _, value := range edges {
		result = append(result, SimpleEdge{value.Head, value.Tail})
	}
	return result
}

func sortedByTopo(graph Graph) Graph {
	sorted := Graph{}
	for _, value := range graph {
		sorted[value.Topo] = value
	}
	return sorted
}

func TopoSort(graph Graph) {
	max := len(graph)
	for _, vertex := range graph {
		if !vertex.IsExplored {
			topoSort(vertex, &max)
		}
	}
}

func topoSort(vertex *Vertex, counter *int) {
	vertex.IsExplored = true
	for _, value := range vertex.Edges {
		if !value.Head.IsExplored {
			topoSort(value.Head, counter)
		}
	}
	vertex.Topo = *counter
	*counter--
}

func DfsScc(vertex *Vertex, SCC int) {
	vertex.IsExplored = true
	vertex.SCC = SCC
	for _, value := range vertex.Edges {
		if !value.Head.IsExplored {
			DfsScc(value.Head, SCC)
		}
	}
}

func FindSCC (vertexes []int, edges []SimpleEdge) Graph {
	graph := Populate(vertexes, edges)
	revGraph := Populate(vertexes, Reverse(edges))
	TopoSort(revGraph)
	sorted := sortedByTopo(revGraph)
	SCC := 0
	for i := 1; i <= len(vertexes); i++ {
		current := sorted[i].Value
		if !graph[current].IsExplored {
			SCC++
			DfsScc(graph[current], SCC)
		}
	}
	return graph
}

func GroupBySCC (graph Graph) [][]int {
	var result [][]int
	resultmap := map[int][]int{}
	for _, value := range graph {
		resultmap[value.SCC] = append(resultmap[value.SCC], value.Value)
	}
	for _, value := range resultmap {
		if value != nil {
			result = append(result, value)
		}
	}
	return result
}