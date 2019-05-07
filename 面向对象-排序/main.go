package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1:声明一个结构体
type Hero struct {
	Name string
	Age int
}
//2:声明一个Hero结构体类型
type HeroSlice []Hero

//3:实现interface接口
func (hs HeroSlice)len()int  {
    return len(hs)
}
//less方法就是解决你是用什么标准进行排序
//1.按照hero的年龄进行从小到大排序 也就是升序排序
func (hs HeroSlice)lens(i,j int)bool  {
	//只要是hs[i].Age < hs[j].Age 为真 那么就是按升序排列
   return hs[i].Age < hs[j].Age
}

func (hs HeroSlice)Swap(i,j int)  {
	//交换
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	//上面三句话等价于下面的一句话
	hs[i],hs[j] = hs[j],hs[i]

}

func main() {
	//先定义一个数组切片
   var intSlice = []int{0,9,-1,8,4,6,7}

   //除了冒泡排序
   //还可以用系统提供的方法进行排序
   sort.Ints(intSlice)
   fmt.Println(intSlice)


   //对结构体切片进行排序
   //除了冒泡排序
   //还可以用系统提供的方法进行排序
   var heroes HeroSlice
   for i :=0;i <10;i++{
   	hero := Hero{Name: fmt.Sprintf("牛魔王%d",rand.Intn(100)),Age:rand.Intn(100)}  //会返回一个格式化的字符串

	   //将hero append到 heroes切片
	   heroes = append(heroes,hero)
	   }


   //看看排序前的顺序
   for _,v := range heroes{
   	fmt.Println(v)
   }

   sort.Sort(heroes)
   fmt.Println("---------排序后-----------")
	//看看排序前的顺序
	for _,v := range heroes{
		fmt.Println(v)
	}
   }





