package main

import "fmt"

type customerView struct {
	//定义必要的字段
	key string   //接收用户的输入
	loop bool //表示是否循环的显示主菜单
}

func (this *customerView)mainMenu(){
	for{
		fmt.Println("----------客户信息管理软件-----------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Println("请选择1-5:")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误请重新输入...")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("你退出客户关系管理系统")
}
func main() {
	//在main函数中创建一个customerView，并运行显示主菜单。。
	customerView := customerView{
		key : "",
		loop : true ,
	}

	//显示主菜单
	customerView.mainMenu()
}
