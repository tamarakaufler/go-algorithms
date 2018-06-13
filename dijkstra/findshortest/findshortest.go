package findshortest

import (
	"fmt"
	"math"
	"sort"
)

type NodeMap map[string]int
type Graph map[string]NodeMap

type weight map[string]int

// FindShortest implements Dijkstra's algorithm
// It uses three data structures, all maps:
//		Graph ....... holds the description of the (weighted) graph, ie node and edge
//		information
//		cost ........ stores gradually updated path length for nodes
//		processed ... stores already processed nodes
func (g Graph) FindShortest() {

	// cost holds distance to a node from the start.
	// cost gets gradually updated with newly found shorted
	// paths of nodes.
	cost := make(weight)

	// setup ----------------------------------
	for k, v := range g {
		//fmt.Printf("%s:\n", k)

		for nn, nw := range v {
			if k == "Start" {
				cost[nn] = nw
			} else {
				if _, ok := cost[nn]; ok {
					continue
				}
				cost[nn] = math.MaxInt32
			}

		}
		fmt.Printf("NODE %s: current cost info -> %#v\n\n", k, cost)
	}

	// implementation ----------------------------------
	processed := map[string]struct{}{}

	node := findLowestCostNode(processed, cost)
	for node != "" {
		children := g[node]
		for chn, chw := range children {
			newWeight := cost[node] + chw
			if cost[chn] > newWeight {
				cost[chn] = newWeight
			}
		}
		processed[node] = struct{}{}
		node = findLowestCostNode(processed, cost)
	}

	// display results
	nodeNames := []string{}
	for k, _ := range cost {
		nodeNames = append(nodeNames, k)
	}
	sort.Strings(nodeNames)
	for _, node := range nodeNames {
		fmt.Printf("Shortest distance from the Start to each node: %s ... %d\n", node, cost[node])
	}
}

// findLowestCostNode finds the node with the shortest path from start from
// so far collected node distances. Already processed nodes are skipped.
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
