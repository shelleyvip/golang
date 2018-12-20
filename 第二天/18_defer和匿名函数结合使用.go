package main

import "fmt"

func main(){
	x := 60
	m := 90
	l := 80
	defer func(hu,zai,ling int) {
		fmt.Printf("hu=%d\n zai=%d\n ling=%d\n",hu,zai,ling)
	}(92,02,06)
	fmt.Printf("x=%d\n m=%d\n l=%d\n",x,m,l)
}