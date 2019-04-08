package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abc123"
	n,err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println(n)
	}
}
