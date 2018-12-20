package main

import "fmt"

func testa(){
	fmt.Println("aaaaa")
}

func testb(x int){
	//设置recover
	defer func(){
		if err := recover();err !=nil{//产生panic异常
			fmt.Println(err)
		}
	}()//别忘了这后面的（）是调用匿名函数


	var a [10]int
	a[x] = 111//当x为20 的时候，导致数组越界，产生一个panic,导致程序崩溃

}
func testc(){
	fmt.Println("ccccc")
}

func main(){
	testa()
	testb(2)
	testc()
}
//////设置recover 语法
//	defer func(){
//		if err := recover();err !=nil{//产生panic异常
//			fmt.Println(err)
//		}
//	}()//别忘了这后面的（）是调用匿名函数