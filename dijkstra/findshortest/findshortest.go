package findshortest

import (
	"fmt"
	"math"
)

type NodeMap map[string]int
type Graph map[string]NodeMap

type weight map[string]int
type parent map[string]string

// FindShortest implements Dijkstra's algorithm
// It uses four data structures, all maps:
//		Graph ... holds the description of the (weighted) graph, ie node and edge
//		information
func (g Graph) FindShortest() {
	cost := make(weight)
	par := make(parent)

	// setup ----------------------------------
	for k, v := range g {
		//fmt.Printf("%s:\n", k)

		if k == "Fin" {
			par[k] = ""
		}

		for nn, nw := range v {
			if k == "Start" {
				par[nn] = k
				cost[nn] = nw
			} else {
				if _, ok := cost[nn]; ok {
					continue
				}
				cost[nn] = math.MaxInt32
				par[nn] = ""
			}

		}
		//fmt.Printf("NODE %s: cost -> %+v\n\n", k, cost)
		//fmt.Printf("NODE %s: parent -> %+v\n\n", k, par)
	}

	// implementation ----------------------------------
	// processed holds already processed nodes
	processed := map[string]struct{}{}

	// cost holds distance to a node from the start.
	node := findLowestCostNode(processed, cost)
	for node != "" {
		children := g[node]
		for chn, chw := range children {
			newWeight := cost[node] + chw
			if cost[chn] > newWeight {
				cost[chn] = newWeight
				par[chn] = node
			}
		}
		processed[node] = struct{}{}
		node = findLowestCostNode(processed, cost)
	}

	fmt.Printf("Shortest distance from the Start to each node: %#v\n", cost)
}

// findLowestCostNode finds the node with the shortest path from start
// it disregards already processed nodes
func findLowestCostNode(processed map[string]struct{}, cost weight) string {
	min := math.MaxInt32
	var node string

	for k, v := range cost {
		_, ok := processed[k]
		if v < min && !ok {
			min = v
			node = k
		}
	}

	return node
}
