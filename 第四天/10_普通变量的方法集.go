package main

import "fmt"

type Shelley struct {
	name string
	sex string
	age int
	addr string
}

func (tmp Shelley)SetinfoValue(){
	fmt.Println("SetinfoValue")
}

func (tmp *Shelley)SetinfoPointer(){
	fmt.Println("SetinfoPointer")

}
func main(){
	p1 :=Shelley{"周","随便",999,"星球"}
	p1.SetinfoPointer()//这里是调用func (tmp *Shelley)SetinfoPointer()指针的内容
	//它中间做了转换处理   先把p1 转换为（&p）在调用（&p).SetinfoPointer()

}
