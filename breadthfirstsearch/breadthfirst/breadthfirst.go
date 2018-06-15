package breadthfirst

type Tree map[string][]string

// BreadthFirst method finds the first node in the given tree, that
// satisfies required criteria, given by the provided function.
func (t Tree) BreadthFirst(top string, selectFunc func(string) bool) string {
	var person string
	var conns []string
	var ok bool
	queue := []string{}

	// Collecting conns of the top person
	conns = t[top]
	queue = append(queue, conns...)
	//fmt.Printf("initial queue=%#v\n", queue)

	// Search one level first before checking the next one
	// Avoid redundant processing by storing already seen connections
	processed := map[string]struct{}{}
	for len(queue) > 0 {
		person, queue = queue[0], queue[1:]

		// fmt.Printf("person=%s\n", person)
		// fmt.Printf("current queue=%#v\n", queue)

		if ok, queue = t.checkSuitability(processed, person, selectFunc, queue); ok {
			return person
		}
	}

	return ""
}

func (t Tree) checkSuitability(processed map[string]struct{}, person string, selectFunc func(string) bool, queue []string) (bool, []string) {
	// check first degree node first
	if _, ok := processed[person]; ok {
		return false, queue
	}

	if selectFunc(person) {
		return true, nil
	}
	// if the first degree node did not satisfy given criteria,
	// collect all its connections for later processing
	for _, con := range t[person] {
		queue = append(queue, con)
	}

	processed[person] = struct{}{}

	return false, queue
}
