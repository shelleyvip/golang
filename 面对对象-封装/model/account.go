package model

import (
	"fmt"
)

//定义一个Account结构体
type account struct {
	accountNo string  //账目
	pwd string     //密码
	balance float64 //余额
}
//创建一个工厂模式的函数-构造函数
func NewAccount(accountNo string ,pwd string,balance float64) *account {
	//if len(accountNo) < 6 || len(accountNo) > 10{
	//	fmt.Println("账号的长度错误")
	//	return nil
	//}
	//if len(pwd) != 6{
	//	fmt.Println("密码长度不对")
	//	return nil
	//}
	//if balance < 20{
	//	fmt.Println("金额不对")
	//	return nil
	//}
	return &account{accountNo: accountNo, pwd: pwd, balance: balance}
}
//方法
//实现的功能是存款 //给它绑定一个方法叫Deposite
func (account *account) Deposite(money float64,pwd string) {
	//1：看下输入的密码是否正确  看他是否等于account结构体里面的pwd
	if pwd != account.pwd{
		fmt.Println("您输入的密码错误")
		return
	}
	//2：检测存款金额是否正确
	if money <= 0{
		fmt.Println("您输入的金额有误")
		return
	}
	//3：密码正确 金额正确 就开始存款
	account.balance += money
	fmt.Println("存款成功")
}
//取款
func (account *account) WithDraw(money float64,pwd string) {
	//1：看下输入的密码是否正确  看他是否等于account结构体里面的pwd
	if pwd != account.pwd{
		fmt.Println("您输入的密码错误")
		return
	}
	//2：检测取款金额是否正确  //不能大于现有的金额 条件如下
	if money <= 0 || money > account.balance{
		fmt.Println("您输入的金额有误")
		return
	}
	//3：密码正确 金额正确 就开始取款
	account.balance -= money
	fmt.Println("取款成功")
}
//查询
func (account *account)Query(pwd string)  {
	//看下输入的密码是否正确//是否等于上面Account里面的pwd
	if pwd != account.pwd{
		fmt.Println("您输入的密码不正确")
		return
	}
	fmt.Printf("您的账户为:%v 余额为:%v\n",account.accountNo,account.balance)
}


//通过Setaccount的方法给account里的字段赋值
func (p *account)SetaccountNo(accountNo string) {
	if len(accountNo) > 6 || len(accountNo) <10{
		p.accountNo = accountNo
	}else {
		fmt.Println("您输入的账户长度不对")
	}
}
//在通过Getaccount来获取字段的值
func (p *account)Getaccount()string  {
     return p.accountNo
}


func (p *account)Setpwd(pwd string) {
	if len(pwd) == 6{
		p.pwd = pwd
	}else {
		fmt.Println("您输入的密码有误")
	}
}

func (p *account)Getpwd()string  {
	return p.pwd
}

func (p *account)Setbalance(balance float64)  {
	if balance < 20 {
		p.balance = balance
	}else {
		fmt.Println("您输入的金额不对")
	}
}

func (p *account)Getbalance()float64  {
	return p.balance
}
