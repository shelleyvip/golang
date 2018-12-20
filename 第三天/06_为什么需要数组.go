package main

import "fmt"

func main(){
	var ia [50]int
	for i:=0;i<50;i++{
		ia[i] = i + 1
		fmt.Printf("ia=%d\n",ia[i])
	}

}
