package main

import "fmt"

func zhou(a,b int) int{
	return (a - b)
}
//函数自身也是一种类型，通过type个函数起名
type Shelley func (int,int)int

func main(){
	var abc int
	abc = zhou(20,10)
	var nihao Shelley
	 nihao = zhou
	 abc = nihao(88,77)
	 fmt.Println("abc=",abc)

}
