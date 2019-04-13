package main

import "fmt"

func main() {
	var filist []func()
	var i int
	for i := 0 ;i<3;i++{
		filist = append(filist, func() {
			fmt.Println("&i")
			fmt.Println(i)
		})
	}
	fmt.Println("i=",i)
	for _,f := range filist{
		f()
	}
}