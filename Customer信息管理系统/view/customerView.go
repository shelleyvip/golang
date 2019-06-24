package main

import (
	"fmt"
	"golang/Customer信息管理系统/service"
)

type customerView struct {
	//定义必要的字段
	key string   //接收用户的输入
	loop bool //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView)list()  {
	//首先获取当前客户的所有信息（在切片中）
	customers := this.customerService.list()
	//显示
	fmt.Println("-------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i :=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------客户列表完成--------------")

	
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
			this.list()
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

	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
