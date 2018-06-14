package main

import (
	"github.com/tamarakaufler/go-algorithms/dijkstra/findshortest"
)

func main() {

	graph := make(findshortest.Graph)

	graph["Start"] = findshortest.NodeMap{
		"Two":   5,
		"Three": 7,
		"Four":  2,
	}

	graph["Two"] = findshortest.NodeMap{
		"Five": 3,
		"Six":  2,
	}
	graph["Three"] = findshortest.NodeMap{
		"Six":   2,
		"Seven": 2,
	}
	graph["Four"] = findshortest.NodeMap{
		"Three": 5,
		"Seven": 7,
		"Eight": 6,
	}
	graph["Five"] = findshortest.NodeMap{
		"Nine": 3,
	}
	graph["Six"] = findshortest.NodeMap{
		"Nine": 3,
	}
	graph["Seven"] = findshortest.NodeMap{
		"Nine": 2,
		"Ten":  1,
	}
	graph["Eight"] = findshortest.NodeMap{
		"Ten":    3,
		"Eleven": 2,
	}
	graph["Nine"] = findshortest.NodeMap{
		"Ten": 2,
		"Fin": 7,
	}
	graph["Ten"] = findshortest.NodeMap{
		"Fin": 1,
	}
	graph["Eleven"] = findshortest.NodeMap{
		"Ten": 3,
		"Fin": 3,
	}
	// graph["Two"] = findshortest.NodeMap{
	// 	"Six":  2,
	// 	"Five": 3,
	// }
	// graph["Three"] = findshortest.NodeMap{
	// 	"Five": 2,
	// }
	// graph["Four"] = findshortest.NodeMap{
	// 	"Five":  7,
	// 	"Three": 2,
	// }
	// graph["Five"] = findshortest.NodeMap{
	// 	"Fin":   3,
	// 	"Seven": 1,
	// }
	// graph["Six"] = findshortest.NodeMap{
	// 	"Fin": 3,
	// }

	graph["Fin"] = findshortest.NodeMap{}

	graph.FindShortest()
}
