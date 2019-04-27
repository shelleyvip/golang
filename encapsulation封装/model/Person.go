package model

import "fmt"

type person struct {
	Name string     //大写的可以访问
	age int         //字段小写 是 不可以导出的
	sal float64     //其它包不可以访问的
}
//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string )*person  {
	return &person{
		Name :name,
	}

}
//为了访问age和sal 我们编写一个SetXXX的方法 和一个GetXXX的方法
func (p *person)Setage(age int)  {
	if age >0 || age < 150 {
		p.age = age
	}else {
		fmt.Println("年龄范围不正确")
		//给程序员一个默认值
	}
}
//获取age的方法
func (p *person)Getage()int  {
	return p.age
}


//薪水
func (p *person)Setsal(sal float64)  {
   if sal >= 3000 && sal <=30000{
   	 p.sal = sal
   }else {
   	fmt.Println("薪水的范围不对")
   }
}

//获取薪水的方法
func (p *person)Getsal()float64  {
	return p.sal
}

