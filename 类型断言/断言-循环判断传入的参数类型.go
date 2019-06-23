package main

import "fmt"

func TypeJudge(items ...interface{})  {
	for x,y := range items{
		switch y.(type) { //这里的type是一个关键字 固定写法
		case bool:
			fmt.Printf("param %d is boll 值是%v\n",x,y)
		case float32:
			fmt.Printf("param %d is float32 值是%v",x,y)
		case float64,int:
			fmt.Printf()

		}
	}
}

func main()  {

}
