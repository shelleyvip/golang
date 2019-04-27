package main

import (
	"fmt"
)

//定义一个Account结构体
type Account struct {
	AccountNo string  //账目
	Pwd string     //密码
	Blance float64 //余额
}
//方法
//实现的功能是存款 //给它绑定一个方法叫Deposite
func (account *Account) Deposite(money float64,pwd string) {
	//1：看下输入的密码是否正确  看他是否等于account结构体里面的pwd
	if pwd != account.Pwd{
		fmt.Println("您输入的密码错误")
		return
	}
	//2：检测存款金额是否正确
	if money <= 0{
			fmt.Println("您输入的金额有误")
		return
	}
	//3：密码正确 金额正确 就开始存款
	account.Blance += money
	fmt.Println("存款成功")
}
//取款
func (account *Account) WithDraw(money float64,pwd string) {
	//1：看下输入的密码是否正确  看他是否等于account结构体里面的pwd
	if pwd != account.Pwd{
		fmt.Println("您输入的密码错误")
		return
	}
	//2：检测取款金额是否正确  //不能大于现有的金额 条件如下
	if money <= 0 || money > account.Blance{
		fmt.Println("您输入的金额有误")
		return
	}
	//3：密码正确 金额正确 就开始取款
	account.Blance -= money
	fmt.Println("取款成功")
}
//查询

func (account *Account)Query(pwd string)  {
        //看下输入的密码是否正确//是否等于上面Account里面的pwd
        if pwd != account.Pwd{
        	fmt.Println("您输入的密码不正确")
			return
		}
        fmt.Printf("您的账户为:%v 余额为:%v\n",account.AccountNo,account.Blance)
}


func main()  {


}
