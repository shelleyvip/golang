package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Id int
}


func main() {
	s := []int{1,3,5,7,9,2,4,6,8,0}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)

	ss := []Student{}
	ss = append(ss,Student{Name:"cc",Id:3})

	ss = append(ss,Student{Name:"aa",Id:2})

	ss = append(ss,Student{Name:"bb",Id:1})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id > ss[j].Id
		return ss[i].Name < ss[j].Name
	})

	fmt.Println(ss)
}