package main

import (
	"fmt"
	"time"

	"github.com/tamarakaufler/go-algorithms/quicksortalg/quicksort"
)

func main() {
	ints := []int{99, 22, 11, 44, 55, 88, 77, 22, 66, 77, 7, 2, 5, 3, 4, 9, 3, 1, 99, 22, 11, 44, 55, 88, 77, 22, 66, 77, 7, 2, 5, 3, 4, 9, 3, 1, 99, 22, 11, 44, 55, 88, 77, 22, 66, 77, 7, 2, 5, 3, 4, 9, 3, 1}
	start := time.Now()
	fmt.Printf("--> %v\n", quicksort.SortMeQuick(ints))
	fmt.Printf("\ttook %s\n", time.Since(start))
}
