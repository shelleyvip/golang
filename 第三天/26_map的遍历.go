package main

import "fmt"

func main()  {
	zhouzhou := map[int]string{1:"周",2:"雪",3:"莉",4:"我",5:"要",6:"雄起"}
	fmt.Println("zhouzhou=",zhouzhou)
	for key,value := range zhouzhou{
		fmt.Printf("key=%v value=%v\n",key,value) //第一个返回的是key 第二个返回的值
	}
//如何判断一个value是否存在
  value, ok := zhouzhou[6]
  	if  ok == true {   //如果上列判断的key键存在那么就打印出key键
  		fmt.Println("zhouzhou=",value)
	}else {     //如果没有这个key 就执行并打印出else的这段话
		fmt.Println("这个key 不存在！！！！！！")
	}

	}


