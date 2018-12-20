package main

import (
	"errors"
	"fmt"
)

func main(){
	err01 := fmt.Errorf("%v","this is normal err01")
	fmt.Println("err01=",err01)
  //生成两个提示错误的 接口的变量

	err02 := fmt.Errorf("this is normal err02")
	fmt.Println("err02=",err02)

	err03 := errors.New("this is normal err03")
	fmt.Println("err03=",err03)


}