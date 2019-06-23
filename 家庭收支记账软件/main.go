package main

import "fmt"

func main() {
    //声明一个变量 接收用户的输入的选项
    key := ""

    //声明一个变量控制是否退出for
    loop := true

    //定义账户的余额
    balance := 10000.00

    //用来记录每次收支的金额
    money := 0.0

    //用来记录每次支出的说明
    note := ""

    //定义一个变量 记录是否有收支行为
    flag := false
    //支出的详情 使用字符串来说明
    //当有收支时，只需要对这个details进行拼接处理即可
    details :=  "收支\t账户余额\t收支金额\t收支说明"

	//显示这个主菜单
	for {
		fmt.Println("\n------------家庭收支记账软件-------------")
		fmt.Println("                             1 收支明细")
		fmt.Println("                             2 登记收入")
		fmt.Println("                             3 登记支出")
		fmt.Println("                             4 退出软件")
		
		fmt.Print("请选择 (1-4):")
		fmt.Scanln(&key)   //提取输入

		switch key {
		case "1":
			fmt.Println("---------当前收支明细记录-----------")
		if flag{
			fmt.Println(details)
		}else {
			fmt.Println("当前没有收支明细...请来一笔吧")
		}
		case "2":
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			balance += money //修改账户余额  有收入所以余额肯定有变动

			fmt.Println("本次收入说明:")
		   fmt.Scanln(&note)

			//将这个收入情况拼接到details变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note	)
			flag =  true

		case "3":
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			//这里需要做一个必要的判断
			if money > balance {
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money

			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note	)
			flag = true
		case "4":
			//退出提示
			fmt.Println("您确定要退出吗 y/n")
			choice := ""
			for{
			fmt.Scanln(&choice)
			if choice == "y" || choice == "n"{
				break
			}
			fmt.Println("您输入的有误 请重新输入:")
			}
			if choice == "y"{
				loop = false
			}
			//loop 默认为true 如果输入的数字不再这个范围之内 那么就是false并提示下列一段话
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop{
			break  //
		}
	}
    fmt.Println("您已退出家庭收支软件的使用")
}
