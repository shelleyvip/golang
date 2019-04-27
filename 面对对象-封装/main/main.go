package main

import (
	"fmt"
	"golang/面对对象-封装/model"
)

func main() {
   p := model.NewAccount("zxl11111","888888",88.8)
	p.SetaccountNo("zxl11111") //账户
   p.Setbalance(88.8)//余额
   p.Setpwd("888888")   //密码
   fmt.Println(p)
   fmt.Println("account=",p.Getaccount(),"balance=",p.Getbalance(),"pwd=",p.Getpwd())
}