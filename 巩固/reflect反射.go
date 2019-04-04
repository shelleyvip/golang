package main

import (
	"fmt"
	"reflect"
)
//专门演示反射
func reflectTest01(b interface{})  {
	//通过反射获取的传入的变量的type，king 和值
	//1先获取到reflect.Type  类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)

	//2获取reflect.values
	rVlue := reflect.ValueOf(b)

	n2 := 2 +rVlue.Int()
	fmt.Println("n2=",n2)

	fmt.Printf("rVlue=%v rVlue Type=%T\n",rVlue,rVlue)
	
}

//对结构体演示的反射
func reflectTst02(b interface{}){
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)
	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("IV=%v Typr=%T\n",iV,iV)

}

type student struct {
	name string
	age int
}


func main() {

	//需求
	//编写一个案例
	//演示对（基本数据类型，interface{}reflect.values）进行反射的基本操作

	// 先定义一个int
	var num int = 100
	reflectTest01(num)

	stu := student{"壮壮",0}

	reflectTst02(stu)
	fmt.Println("stu=",stu)



}
