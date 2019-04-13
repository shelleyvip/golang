package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1,3,5,7,9,2,4,6,8,0}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}