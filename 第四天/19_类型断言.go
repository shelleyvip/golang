package main

import "fmt"

type student struct {
	name string
	id int
}

func main()  {
	//空接口类型 （interface{}）有3个元素
	i := make([]interface{},3)
	i[0] = "周雪莉"                        //string
	i[1] = 888                            //int
	i[2] = student{"女热播",89889}//student

	//类型查询 //类型断言
	//第一个返回下标，第二个返回下标对应的值，data其实局势空接口变量，分别是i[0],i[1],i[2]
	for index,data := range i {
		//在这里第一个返回的是值（也就是接口变量本身）第二是返回的是判断结果的真假
		if value,ok := data.(int);ok == true{
			fmt.Printf("x[%d] 类型为string,内容为%d\n",index,value)
		}else if value,ok := data.(string);ok == true{
			fmt.Printf("x[%d] 类型为int 内容为%s\n",index,value)
		}else if value,ok := data.(student);ok == true{
			fmt.Printf("x[%d] 类型为student 内容为%s %d\n",index,value.name,value.id)
		}
	}

}