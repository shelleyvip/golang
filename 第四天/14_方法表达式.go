package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}

func (tmp Shelley)Addr01(){
	fmt.Printf("Addr01:%p %v\n",&tmp,tmp)
}

func (tmp *Shelley)Addr02(){
	fmt.Printf("Addr02:%p %v\n",tmp,tmp)
}
func main(){
	p := Shelley{"周",25,"女"}
	fmt.Printf("main:%p %v\n",&p,p)

	//上一讲 的方法值 f := p.Assr01 //是隐藏了接受者
	//方法表达式：方法值：语法
	//现在的方法式 是把接受者显现出来
	f := (Shelley).Addr01
	f(p)//显示把接受者传过去==》p.Addr01

	f2 := (*Shelley).Addr02
	f2(&p)//显示把接受者传过去==》p.Addr02
}
