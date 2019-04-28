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
}
