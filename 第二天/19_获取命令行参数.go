package main

import (
	"fmt"
	"os"
)

func main(){
	zhouxueli := os.Args
	ni := len(zhouxueli)
	fmt.Println("ni=",ni)
	for i, data := range zhouxueli{
		fmt.Printf("shelley=[%d] %s\n",i,data)
	}
}
