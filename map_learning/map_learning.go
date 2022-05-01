package map_learning

type Vertex struct {
	height, weight float64
}

var m map[string]Vertex

func Map_vertex() Vertex {
	m = make(map[string]Vertex)
	m["me"] = Vertex{
		157, 44,
	}
	return m["me"]
}
