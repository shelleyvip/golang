package main

import "fmt"

func main()  {
	var a[20]int
	for i:=0;i<len(a);i++ {
		a[i] = i + 1
		fmt.Printf(" a=[%d]\n",a[i])

	}
}