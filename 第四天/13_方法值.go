package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}
func (tmp Shelley)Addr01(){
	fmt.Printf("Addr01=%p %v\n",&tmp,tmp)
}

func(tmp *Shelley)Addr02(){
	fmt.Println("tmp=",tmp)
}

func main(){
	p := Shelley{"周",25,"女"}
	fmt.Printf("main:p=%p %v\n",&p,p)
	p.Addr01()//👈这个是传统调用方法

	//👇这个就是方法值 //保存方法的语法 //
	Func111 := p.Addr01//👈 调用函数时，无须在传递接收者，隐蔽了接收者
	Func111() //这里等价于p.Addr01()

	//👇这个就是方法值 //保存方法的语法 //
	Func222 := p.Addr02//👈调用函数时，无须在传递接收者，隐蔽了接收者
	Func222()//那么这里就等价于p.Addr02
}