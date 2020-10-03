package dijkstra

import (
	"container/heap"
	"math"
)

//Graph
type Vertex struct {
	Value int
	Length int
	Edges []Edge
}

type Edge struct {
	Length int
	Tail, Head *Vertex
}

type Graph map[int]*Vertex

//simple edges representation
type SimpleEdge struct {
	Tail, Head, Length int
}

func Populate(vertexes []int, edges []SimpleEdge) Graph {
	graph := Graph{}
	for _, value := range vertexes {
		graph[value] = &Vertex{
			Value: value,
			Length: 1000,
		}
	}
	for _, sedge := range edges {
		graph[sedge.Tail].Edges = append(graph[sedge.Tail].Edges, Edge{
			Length: sedge.Length,
			Tail: graph[sedge.Tail],
			Head: graph[sedge.Head],
		})
	}
	return graph
}

func FindPath(vertexes []int, edges []SimpleEdge, start int) Graph {
	graph := Populate(vertexes, edges)
	graph[start].Length = 0

	h := PriorityQueue{}
	h.pq = make([]*Item, len(graph))
	h.ValueIndex = make(map[int]int, len(graph))
	i := 0
	for id, value := range graph {
		h.pq[i] = &Item{
			id: id,
			len: value.Length,
			index: i,
		}
		h.ValueIndex[id] = i
		i++
	}
	heap.Init(&h)

	for len(h.pq) > 0 {
		vert := heap.Pop(&h).(*Item)
		graph[vert.id].Length = vert.len
		if len(h.pq) > 0 {
			for _, edge := range graph[vert.id].Edges {
				min := int(math.Min(float64(h.get(edge.Head.Value).len), float64(vert.len+edge.Length)))
				h.update(edge.Head.Value, min)
			}
		}
	}

	return graph
}


//Heap
type Item struct {
	id int // vertex id
	len int    // vertex len (priority)
	index int // The index of the item in the heap.
}

type PriorityQueue struct {
	ValueIndex map[int]int //map of indexes (for updating)
	pq []*Item
}

func (pq PriorityQueue) Len() int { return len(pq.pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.pq[i].len < pq.pq[j].len
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.ValueIndex[pq.pq[i].id], pq.ValueIndex[pq.pq[j].id] = pq.ValueIndex[pq.pq[j].id], pq.ValueIndex[pq.pq[i].id]
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
	pq.pq[i].index = i
	pq.pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.pq)
	item := x.(*Item)
	item.index = n
	pq.pq = append(pq.pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.pq = old[0 : n-1]
	delete(pq.ValueIndex, item.id)
	return item
}

func (pq *PriorityQueue) update(id int, len int) {
	index := pq.ValueIndex[id]
	pq.pq[index].len = len
	heap.Fix(pq, index)
}

func (pq *PriorityQueue) get(id int) *Item {
	return pq.pq[pq.ValueIndex[id]]
}

