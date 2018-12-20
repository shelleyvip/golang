package main

import "fmt"

type Shelley struct {
	name string
	sex string
	age int
}//修改普通变量，非指针 值语义，就等于数是拷贝了一份过来（和值传递是一样的）
// *值语义*  这种方法是值语义也就是（所谓的值传递）内容是里面改了外面是改不了的
func (tmp Shelley)Addr01(a string,b string,c int){
	tmp.name = a
	tmp.sex = b
	tmp.age = c
	fmt.Println("tmp=",tmp)
}
//这里接受者为指针变量 也就是是引用传递（指针传递）
//func (tmp *Shelley)Pointer(a string,b string,c int){
//tmp.name = a
//tmp.sex = b
//tmp.age = c
//fmt.Println("tmp=",tmp)
//}

func main(){
	p1 :=Shelley{"你","d",88}
	fmt.Printf("p1=%p\n",&p1)

	p1.Addr01("nn","bb",99)//值语义传递修改Addr01的值
	fmt.Println("p1=",p1)//里面变了外面还是没有变

	//(&p1).Pointer("傻逼","变态",888)
	//fmt.Println("p1=",p1)//&地址 是里面变了 外面也了

}
