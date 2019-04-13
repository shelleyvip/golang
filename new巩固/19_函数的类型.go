package main

import "fmt"

func println()  {
fmt.Println("hello world")
}
func Println1()  {
	fmt.Println("shelley")
}

type fstsct struct {
	Func func()
}

func func1(n int)int  {
       return n + 1
}
func main() {
	var f func()
	var flist[3] func(int)int
	flist[0] = func1
	flist[0](10)
	fmt.Println(flist)
	f = println
	f()

	f = Println1
	f()


}