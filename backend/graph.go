package main

type node struct {
	weight   float64
	outbound float64
}

type graph struct {
	edges map[string](map[string]float64)
	nodes map[string]*node
}

func newGraph() *graph {
	return &graph{
		edges: make(map[string](map[string]float64)),
		nodes: make(map[string]*node),
	}
}

func (graph *graph) link(source, target string, weight float64) {
	if _, ok := graph.nodes[source]; ok == false {
		graph.nodes[source] = &node{
			weight:   0,
			outbound: 0,
		}
	}

	graph.nodes[source].outbound += weight

	if _, ok := graph.nodes[target]; ok == false {
		graph.nodes[target] = &node{
			weight:   0,
			outbound: 0,
		}
	}

	if _, ok := graph.edges[source]; ok == false {
		graph.edges[source] = map[string]float64{}
	}

	graph.edges[source][target] += weight
}