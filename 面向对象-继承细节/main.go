package main

import "fmt"

type A struct {
    Name string
    age int
}

func (a *A)SayOk()  {
	fmt.Println("A SayOk",a.Name)
}

func (a *A)Hello()  {
	fmt.Println("A Hello",a.Name)
}

type B struct {
	A
	Name string
}

func (b *B)SayOk()  {
	fmt.Println("B SayOk",b.Name)
}


//第三：有名结构体
type C struct {
	a A     //嵌套一个有名结构体           //组合关系
	//如果你要是想访问A就必须把a带上去 不然会报错
}

//第四小案例
type Goods struct {   //goods产品
	Name string
	Price float64
}
type Brand struct {   //brand品牌
	Name string
	Address string
}
type Tv struct{
	Goods
	Brand
}

//第五小案例  也可以指针传递 指针传递过去的是地址
type Tv2 struct {
	*Goods
	*Brand
}


func main()  {
   //var b B
   //b.A.Name = "shelley"
   //b.A.age = 26
   //b.A.Hello()
   //b.A.SayOk()

   //细节1☝️
   ////上买的写法可以简化
   //b.age = 27                 //对这里的代码小结
   //b.Name = "zxl"         //当我们直接通过b访问字段或方法时，其执行如下比如b.Name字段
   //b.SayOk()             //如果B没有此字段就去看B潜入的匿名结构体A有没有声明这个Name字段，
   //b.Hello()            //如果有就调用，如果没有就继续查找，如果都找不到 就会报错

   //细节✌️
   var b B
   b.Name = "tom" //ok 是给b.B里面的Name赋的值
   b.A.Name = "jack" //如果非要给A的字段Name赋值 那就显示调用
	b.age = 88     //ok 同理  是给b.B里面的age赋的值
   b.SayOk()    //B SayOk tom  调用了b的方法
   b.Hello()    //就近原则如果本作用域找不到这个字段 那就去找嵌入的结构体中去找 如果也没有赋值那么就是空串

   //细节☝️✌️
   var c C
	//c.Name error// //嵌套一个有名结构体           //组合关系
	//	//如果你要是想访问A就必须把a带上去 不然会报错
	//如果C里面有Name的字段就不会报错
	c.a.Name = "Nan"

	//细节四✌️✌️
	//嵌套结构体后，可是在创建结构体变量实例时，可以直接指定各个匿名函数的值
	tv := Tv{Goods{"电视机001",7777.77,},Brand{"海尔","山东"}}

	//两种方法都是可以的
	tv1 := Tv{
		Goods{
			Name:  "电视机002",
			Price: 8888.88,
		},
		Brand{
			Name : "爱立信",
			Address: "青岛",
		},
	}
   fmt.Println("tv",tv)
	fmt.Println("tv1=",tv1)

	tv2 := Tv2{&Goods{"电视机003",9999.99},&Brand{"创维","上海"}}
	fmt.Println("tv2=",tv2)//这样取得话打印出来的是两个地址
	fmt.Println("t2=",*tv2.Goods,*tv2.Brand)

}
