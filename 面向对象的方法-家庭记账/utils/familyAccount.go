package utils

import "fmt"

type FamilyAccount struct {
	//声明必须的字段

	//声明一个变量 接收用户的输入的选项
	key string

	//声明一个变量控制是否退出for
	loop bool

	//定义账户的余额
	balance float64

	//用来记录每次收支的金额
	money float64

	//用来记录每次支出的说明
	note string

	//定义一个变量 记录是否有收支行为
	flag bool
	//支出的详情 使用字符串来说明
	//当有收支时，只需要对这个details进行拼接处理即可
	details string
}
//编写工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount()*FamilyAccount{
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.00,
		money : 0.0,
		note : "",
		flag : false,//在收支情况下flag是纪律有米有收支行为的 默认情况下是false
		details:"收支\t账户余额\t收支金额\t收支说明",

	}

}
//将显示明细写成一个方法
func (this *FamilyAccount)showDetails()  {
	fmt.Println("---------当前收支明细记录-----------")
	if this.flag{
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有收支明细...请来一笔吧")
	}
}
//将登记收入写成一个方法 和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额  有收入所以余额肯定有变动

	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)

	//将这个收入情况拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note	)
	this.flag =  true
}
//将登记支出 写一个方法和*FamilyAccount绑定
func (this *FamilyAccount)pay()  {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额的金额不足")

	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note	)
	this.flag = true
}

//将退出系统也写成一个方法 和FamilyAccount绑定
func (this *FamilyAccount)exit() {
	//退出提示
	fmt.Println("您确定要退出吗 y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("您输入的有误 请重新输入:")
	}
	if choice == "y" {
		this.loop = false
	}

}
//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount)MainMenu() {

	//显示这个主菜单
	for {
		fmt.Println("\n------------家庭收支记账软件-------------")
		fmt.Println("                             1 收支明细")
		fmt.Println("                             2 登记收入")
		fmt.Println("                             3 登记支出")
		fmt.Println("                             4 退出软件")

		fmt.Print("请选择 (1-4):")
		fmt.Scanln(&this.key) //提取输入

		switch this.key {
		case "1":
			this.showDetails()

		case "2":
			this.income()
		case "3":
             this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
}
		if !this.loop{
			break
		}
}
	}