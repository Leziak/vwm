package main

import (
	"fmt"
	"os"
	"math"
)


func (graph *graph) pagerank(α, ε float64) {
	pagerankfile, _ := os.Create("pagerankfile.txt")
	Δ := float64(1.0)
	inverse := 1 / float64(len(graph.nodes))

	// Normalize all the edge weights so that their sum amounts to 1.
	for source := range graph.edges {
		if graph.nodes[source].outbound > 0 {
			for target := range graph.edges[source] {
				graph.edges[source][target] /= graph.nodes[source].outbound
			}
		}
	}

	for key := range graph.nodes {
		graph.nodes[key].weight = inverse
	}

	for Δ > ε {
		leak := float64(0)
		nodes := map[string]float64{}

		for key, value := range graph.nodes {
			nodes[key] = value.weight

			if value.outbound == 0 {
				leak += value.weight
			}

			graph.nodes[key].weight = 0
		}

		leak *= α

		for source := range graph.nodes {
			for target, weight := range graph.edges[source] {
				graph.nodes[target].weight += α * nodes[source] * weight
			}

			graph.nodes[source].weight += (1-α)*inverse + leak*inverse
		}

		Δ = 0

		for key, value := range graph.nodes {
			Δ += math.Abs(value.weight - nodes[key])
		}
	}

	for key, value := range graph.nodes {
		pagerankfile.WriteString(key + "\t" + fmt.Sprintf("%f", value.weight) + "\n")
	}
}