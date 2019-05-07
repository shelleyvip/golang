package main

import "fmt"

type Monkey struct {
	Name string
}

type LittleMonkey struct {
	Monkey
}

func (m *Monkey)climbing()  {
    fmt.Println(m.Name,"生来会爬树")
}

type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimin()
}

//让LittleMonkey 实现BirdAble //实现他的方法就等于实现BirdAble
func (m *LittleMonkey)Flying()  {
	fmt.Println(m.Name,"通过学习 学会了飞翔")
}


//让LittleMonkey 实现/实现FishAble 实现他的Swimin方法就等于实现FishAble
func (m *LittleMonkey) Swimin() {
	fmt.Println(m.Name,"通过学习 学会了游泳")
}

func main() {

	//创建一个LittleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "孙悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimin()
}
