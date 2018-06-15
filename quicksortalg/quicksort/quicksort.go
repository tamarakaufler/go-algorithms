package quicksort

import (
	"log"
	"math/rand"
)

// SortMeQuick is a recursive function for sorting
// an array of integers
func SortMeQuick(si []int) []int {

	// base case
	if len(si) <= 1 || isEqual(si) {
		return si
	}

	// recursive case
	pi := rand.Intn(len(si))
	pivot := si[pi]

	log.Printf(">>> pivot index: %d, pivot value: %d\n", pi, pivot)

	smaller := []int{}
	bigger := []int{}
	for _, el := range si {
		if el <= pivot {
			smaller = append(smaller, el)
		} else {
			bigger = append(bigger, el)
		}
	}

	// log.Printf("\tsmaller: %#v\n", smaller)
	// log.Printf("\tbigger: %#v\n-----------------------\n\n", bigger)

	return append(SortMeQuick(smaller), SortMeQuick(bigger)...)
}

func isEqual(l []int) bool {
	lng := len(l)
	if lng <= 1 {
		return true
	}
	v := l[0]
	for i := 0; i <= lng-1; i++ {
		if l[i] != v {
			return false
		}
	}
	return true
}
