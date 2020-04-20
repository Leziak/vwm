package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"math"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

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

func (self *graph) link(source, target string, weight float64) {
	if _, ok := self.nodes[source]; ok == false {
		self.nodes[source] = &node{
			weight: 0,
			outbound: 0,
		}
	}

	self.nodes[source].outbound += weight

	if _, ok := self.nodes[target]; ok == false {
		self.nodes[target] = &node{
			weight: 0,
			outbound: 0,
		}
	}

	if _, ok := self.edges[source]; ok == false {
		self.edges[source] = map[string]float64{}
	}

	self.edges[source][target] += weight
}

func (self *graph) rank(α, ε float64, callback func(id string, rank float64)) {
	Δ := float64(1.0)
	inverse := 1 / float64(len(self.nodes))

	// Normalize all the edge weights so that their sum amounts to 1.
	for source := range self.edges {
		if self.nodes[source].outbound > 0 {
			for target := range self.edges[source] {
				self.edges[source][target] /= self.nodes[source].outbound
			}
		}
	}

	for key := range self.nodes {
		self.nodes[key].weight = inverse
	}

	for Δ > ε {
		leak := float64(0)
		nodes := map[string]float64{}

		for key, value := range self.nodes {
			nodes[key] = value.weight

			if value.outbound == 0 {
				leak += value.weight
			}

			self.nodes[key].weight = 0
		}

		leak *= α

		for source := range self.nodes {
			for target, weight := range self.edges[source] {
				self.nodes[target].weight += α * nodes[source] * weight
			}

			self.nodes[source].weight += (1-α)*inverse + leak*inverse
		}

		Δ = 0

		for key, value := range self.nodes {
			Δ += math.Abs(value.weight - nodes[key])
		}
	}

	for key, value := range self.nodes {
		callback(key, value.weight)
	}
}

func getRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func crawl(graph *graph) {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.MaxDepth(3), colly.Async(true))
	outlink := ""
	cnt := 0
	input := "https://en.wikipedia.org/wiki/United_States"

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		inlink := e.Attr("href")
		outlink = e.Request.URL.String()
		cnt++

		/* Tento big brain hack znizi pocet moznych hran v grafe na max zopar tisic tym,
		* ze budem robit request iba raz za 50-100 moznych requestov.
		* Navyse sa bude vystupny subor vzdy lisit od toho predchadzajuceho.
		 */
		if n := getRand(80, 81); cnt%n != 0 {
			return
		}

		/*
		* Pre istotu takto trapne exhaustivne, aby som neodignoroval legit wiki stranky
		 */
		if strings.HasPrefix(inlink, "/wiki/") &&
			!strings.HasPrefix(inlink, "/wiki/Wikipedia") &&
			!strings.HasPrefix(inlink, "/wiki/Category") &&
			!strings.HasPrefix(inlink, "/wiki/File") &&
			!strings.HasPrefix(inlink, "/wiki/Help") &&
			!strings.HasPrefix(inlink, "/wiki/User") &&
			!strings.HasPrefix(inlink, "/wiki/Template") &&
			!strings.HasPrefix(inlink, "/wiki/ISBN") &&
			!strings.HasPrefix(inlink, "/wiki/Portal") &&
			!strings.HasPrefix(inlink, "/wiki/Talk") &&
			!strings.HasPrefix(inlink, "/wiki/Special") {
			decodedOutlink, _ := url.PathUnescape(outlink)
			decodedInlink, _ := url.PathUnescape("https://en.wikipedia.org" + inlink)
			fmt.Println(cnt, decodedOutlink, decodedInlink)
			// file.WriteString(decodedOutlink + "\t" + decodedInlink + "\n")
			graph.link(decodedOutlink, decodedInlink, 1.0)
			e.Request.Visit(inlink)
		}

	})

	c.Visit(input)

	c.Wait()
}

func main() {
	file, _ := os.Create("links.txt")
	graph := newGraph()
	crawl(graph)
	graph.rank(0.85, 0.00001, func(node string, rank float64) {
		file.WriteString(node + "\t" + fmt.Sprintf("%f", rank) + "\n")
		// fmt.Println("Node", node, "has a rank of" , rank)
	})
}
