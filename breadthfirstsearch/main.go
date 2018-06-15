package main

import (
	//"github.com/tamarakaufler/go-algorithms/breadthfirstsearch/breadthfirst"
	"fmt"

	"github.com/tamarakaufler/go-algorithms/breadthfirstsearch/breadthfirst"
)

func main() {
	//tree := make(breadthfirst.Tree)

	tree := breadthfirst.Tree{
		"Tamara": []string{"Mirko", "Lucien", "Dominic"},

		"Mirko":   []string{"Zuzana", "MirkoJr", "Marek"},
		"Lucien":  []string{"Annie", "Will", "Matt"},
		"Dominic": []string{"Jane", "Coral", "Lauraine"},

		"Zuzana":   []string{"Martina"},
		"MirkoJr":  []string{"Martin", "Tomas"},
		"Marek":    []string{"Anicka", "Petr"},
		"Annie":    []string{"Melanie", "Petra"},
		"Will":     []string{"Tom"},
		"Matt":     []string{},
		"Lauraine": []string{"Nadine", "Keira"},
		"Coral":    []string{"Hugo", "Stephanie", "David", "Sarah", "Claire"},
		"Jane":     []string{"Phil", "Samantha", "Christopher"},
	}

	selected := tree.BreadthFirst("Tamara", isFree)
	fmt.Printf("\n 👍 Found %s 👍\n\n", selected)
}

func isFree(g string) bool {
	if len(g) > 7 {
		return true
	}
	return false
}
