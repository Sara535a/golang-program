package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	for {
		var n int
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		s = append(s, n)
	}
	sort.Ints(s)
	lent := len(s) - 1
	if s[lent/2] == s[lent/2+1] {
		fmt.Println(s[lent/2])
	} else {
		fmt.Println("-1")
	}
}
