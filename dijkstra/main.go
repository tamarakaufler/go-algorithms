package main

import (
	"github.com/tamarakaufler/go-algorithms/dijkstra/findshortest"
)

func main() {

	graph := make(findshortest.Graph)

	graph["Start"] = findshortest.NodeMap{
		"Two":   5,
		"Three": 6,
		"Four":  2,
	}
	graph["Two"] = findshortest.NodeMap{
		"Six":  2,
		"Five": 3,
	}
	graph["Three"] = findshortest.NodeMap{
		"Five": 2,
	}
	graph["Four"] = findshortest.NodeMap{
		"Five":  7,
		"Three": 2,
	}
	graph["Five"] = findshortest.NodeMap{
		"Fin":   3,
		"Seven": 1,
	}
	graph["Six"] = findshortest.NodeMap{
		"Fin": 3,
	}
	graph["Seven"] = findshortest.NodeMap{
		"Fin": 1,
	}
	graph["Fin"] = findshortest.NodeMap{}

	graph.FindShortest()
}
