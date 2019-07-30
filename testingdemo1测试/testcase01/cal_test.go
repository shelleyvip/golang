package main

import (
	"fmt"
	"testing"
) //引入go 的testing轻量级 testing框架的包

//编写一个测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T)  {

	//调用cal.go里面的addUpper
	res := addUpper(10)
	if res != 55{
			t.Fatalf("AddUpper(10)执行错误 期望值=%v 实际值=%v\n",55,res)
	}
	t.Logf("AddUpper(10)执行正确...")
}

func TestHello(t *testing.T)  {
	fmt.Println("TestHello 被测试。。。。")
}


//测试单个文件 一定要带上被测试的源文件
 	//go test -v sab_test.go cal.go


 	//测试单个方法
 	// go test -v test.run TestHello