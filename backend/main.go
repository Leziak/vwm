package main

func main() {
	graph := newGraph()
	crawl(graph)
	graph.pagerank(0.85, 0.00001)
}