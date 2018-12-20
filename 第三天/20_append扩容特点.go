package main

import "fmt"

func main()  {
	s := make([]int,0,1)
	old := cap(s)
	for i := 0;i < 20;i++ {
		s = append(s,i)
		if new := cap(s); old < new {
			fmt.Printf("cap:  %d====>%d\n",old,new)
			old = new
		}
	}
}