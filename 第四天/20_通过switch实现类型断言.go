package main

import "fmt"

type Shelley struct {
	name string
	age int
}
func main(){
	s1 :=make([]interface{},3)
	s1[0]= "米露"
	s1[1]= 20160628
	s1[2]= Shelley{"米露出生于",20160628}
	for index,data := range s1{
		switch value := data.(type){
		case string:
			fmt.Printf("[%d]内容是string 类型是%s\n",index,value)
		case int:
			fmt.Printf("[%d]内容是int 类型是%d\n",index,value)
		case Shelley:
			fmt.Printf("[%d]内容是Shelley 类型是%s %d\n",index,value.name,value.age)
		}
	}
}