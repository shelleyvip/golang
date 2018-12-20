package main

import (
	"fmt"
	"regexp"
)

func main(){
	rege02 := "3.14 2.3 asd 6. 22.33 df 4."

	smu := regexp.MustCompile(`/d+/./d+/`)
	if smu == nil{
		fmt.Println("MustCompile err")
		return
	}
	result := smu.FindAllStringSubmatch(rege02,-1)
	fmt.Println("result=",result)
}