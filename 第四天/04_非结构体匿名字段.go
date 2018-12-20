package main
import"fmt"

type xue string//自定义类型 给一个类型改名


type shelley struct {
	name string
	age int
	addr string
}
type zhou struct {
	shelley //结构体匿名字段  //这个比较常用
	 int    //基础类型的匿名字段 //这个不是很常用
	xue     //
}

func main(){           //后面这两个是上面int和 自定义类型xue 也就是string的内容
	s1 := zhou{shelley{"周",26,"北京"},666666,"呵呵呵"}
	fmt.Printf("s1=%+v\n",s1)

	//还可以整体给shelley做一个赋值；方法如下
	s1.shelley = shelley{"李",88,"北京"}

	                                    //在这里直接打印int,和xue就可以了
	fmt.Println(s1.name,s1.age,s1.addr,s1.int,s1.xue)

	             //也可以把shelley当作一个整体来操作,直接套进去就可以了
	fmt.Println(s1.shelley,s1.int,s1.xue )


}