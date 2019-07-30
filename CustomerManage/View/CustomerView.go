package main

import (
	"fmt"
	"golang/CustomerManage/model"
)
import "golang/CustomerManage/Service"

type customerView struct {
     //定义必要的字段
     key string //接收用户的输入
     loop bool //表示是否循环的显示主菜单
     //增加一个字段CustomerService
     customerService *Service.CustomerService
}
//显示客户所有信息
func (this *customerView)list()  {
	//首先获取到当前所有的客户信息（其实就在切片中）
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0;i<len(customers);i++{
       fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("--------------------客户列表完成----------------\n\n")

}

//得到用户的输入信息，构建新的客户 并完成添加
func (this *customerView)Add()  {
	 fmt.Println("----------------添加客户-----------------")
	 fmt.Println("姓名:")
	 name := ""
	 fmt.Scanln(&name)

	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)


	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("电邮")
	email := ""
	fmt.Scanln(&email)


	//构建一个新的Customer的实例
	//注意ID号没有让用户输入，因为ID是唯一的 需要系统分配
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer){
		fmt.Println("--------------添加完成---------------")
	}else{
		fmt.Println("--------------添加失败---------------")
	}

}

//得到用户的输入的ID 删除对应的客户
func (this *customerView)delete()  {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Println("请选择待删除的客户编号（-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return  //放弃删除操作 返回
	}
	fmt.Println("确定是否删除（y/n）:")
	choice := ""

		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {//|| choice == "n" || choice == "N"{

		//调用customerService的Delete方法
		if this.customerService.Delete(id){
			fmt.Println("---------------删除完成----------------")
		}else{
			fmt.Println("---删除失败，输入的ID号不存在请重新输入---------")
		}
		}

	}


//退出软件
func (this *customerView)exit()  {
	fmt.Println("你确定要退出吗")
	for{
		fmt.Scanln(&this.key)
		if this.key == "y"||this.key == "Y"||this.key == "n" || this.key == "N"{
			break
		}else {
			fmt.Println("您的输入有误,请确认输入的是(y/n)")
		}
		if this.key == "y" ||this.key == "Y"{
			this.loop = false
		}
	}
}

//修改客户
func (this *customerView)update()  {
	fmt.Println("请选择待修改客户的编号:")

}

//显示主菜单
func (this *customerView)mainMenu()  {
      for{
      	  fmt.Println("---------------客户信息管理软件--------------")
		  fmt.Println("                        1 添加客户")
		  fmt.Println("                        2 修改客户")
		  fmt.Println("                        3 删除客户")
		  fmt.Println("                        4 客户列表")
      	  fmt.Println("                        5 退   出")
      	  fmt.Print("请选择1-5:")

      	  fmt.Scanln(&this.key)
		  switch this.key {
		  case "1":
		  	this.Add()
		  case "2":
		  	fmt.Println("修改客户")
		  case "3":
		  	this.delete()
		  case"4":
		this.list()
		  case "5":
		  	this.exit()
		  default:
			  fmt.Println("您的输入有误请重新输入")
		  }

         if !this.loop{
         	break
		 }
	  }
	fmt.Println("你退出了客户关系管理系统")

}

func main()  {
     //在main函数中 创建一个CustomerView，并运行主菜单
      customerView := customerView{
      	key : "",
      	loop : true,
	  }

      //这里完成customerView结构体的对customerService字段的初始化
      customerView.customerService = Service.NewCustomerService()

	//显示主菜单..
	customerView.mainMenu()

}
