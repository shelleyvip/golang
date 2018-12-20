package main

import "fmt"

type setinfo interface {//这个是子集
	object()
}
type student interface {//这个是超集 因为它范围比较多的
	setinfo  //匿名字段，继承了object（）的接口
	result(zhou string)//	在这里我自己student还有一个方法，这个方法带有一个参数
}
type Shelley struct{
	name string
	id int
}

func (tmp *Shelley)object(){
	fmt.Printf("Shelley[%s %d]\n",tmp.name,tmp.id)
}
func(tmp *Shelley)result(zhou string){
	fmt.Println("周武军在唱歌:",zhou)
}
func main(){
	//超集可以转换给子集，反过来不可以
	var s1 student   //超级
	s1 = &Shelley{"周武军在ktv唱歌",666}

///     结论是超级可以转换为子集

	var i setinfo     //子集
	i = s1
	i.object()



}