package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	num  int
	list string
)

func init() {
	flag.IntVar(&num, "num", 0, "A number to guess")
	flag.StringVar(&list, "list", "", "List of integers as a space separates string")
}

func main() {
	flag.Parse()

	fmt.Printf("--> %s\n", list)

	sl := strings.Split(list, " ")
	var il []int
	for i := 0; i < len(sl); i++ {
		n, err := strconv.Atoi(sl[i])
		if err != nil {
			fmt.Printf("err=%v\n", err)
			continue
		}
		il = append(il, n)
	}
	sort.Ints(il)
	fmt.Printf("--> %v\n", il)

	// guessing ...
	guess(num, il)
}

func guess(num int, list []int) {
	var iMin, iMax, mid, g int

	iMin = 0
	iMax = len(list) - 1

	for {
		if iMin < 0 || iMax > len(list)-1 || iMax < iMin {
			fmt.Printf("you num %d was not found\n", num)
			break
		}

		mid = (iMax-iMin)/2 + iMin
		g = list[mid]
		fmt.Printf("Guessing  min=%d, mid=%d,max=%d, guess=%d\n", iMin, mid, iMax, g)

		switch {
		case g == num:
			fmt.Printf("you num is >> %d <<\n", g)
			return
		case num > g:
			iMin = mid + 1
		default:
			iMax = mid - 1
		}
		time.Sleep(500 * time.Millisecond)
	}
}
