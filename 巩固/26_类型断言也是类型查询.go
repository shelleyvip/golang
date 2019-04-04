package main

import "fmt"

type Studenta struct {
	id int
	name string
}

func main() {
    i := make([]interface{},3)
	i[0] = 666
	i[1] = "hello"
	i[2] = Studenta{22,"zhouzhou"}
      //第一个是下标 第二个是下标对应的值
	for index,data := range i{
		//第一个是值 第二个是判断结果的真假
		if values,ok := data.(int); ok == true{
			fmt.Printf("[%d] 类型为int 内容为%d\n",index,values)
		}else if values,ok := data.(string);ok == true {
			fmt.Printf("[%d] 类型为string 内容为%s\n",index,values)
		}else if values,ok := data.(Studenta);ok == true{
			fmt.Printf("[%d] 类型是id =int name =string 内容是%d %s\n",index,values.id,values.name)
		}
	}

}

