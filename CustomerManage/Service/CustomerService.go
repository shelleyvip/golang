package Service

import (
	"fmt"
	"golang/CustomerManage/model"
)

//该CustomerService 完成对Customer的操作 包括增删改查
type CustomerService struct {
	 //cstomers []model.Customer 你 少写一个U  字母
	customers []model.Customer

	//声明一个字段 表示当前切片含有多少个客户
	//该字段还可以作为  新客户的编号 ID+1
	customerNum int

}
//编写一个方法可以返回*CustomerService方法
func NewCustomerService() *CustomerService  {
	//为了可以看到有客户在切片中，我们初始化一个客户
	costomerService := &CustomerService{}
	costomerService.customerNum = 1
	costomer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	costomerService.customers = append(costomerService.customers,costomer)
	return costomerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到customer切片

func (this *CustomerService)Add(customer model.Customer)bool  {
	this.customerNum++
	customer.Id = this.customerNum
	//我们确定一个分配ID的规则，就是添加的顺序
	this.customers = append(this.customers,customer)
	return true
	
}
//根据ID删除客户（从切片中删除）
func (this *CustomerService)Delete(id int)bool {
	index := this.FindById(id)

	//如果index == -1 说明没有这个客户ID号
	if  index == -1{
    return false
	}
	//如果找到的话就删除
	//如何从切片中删除一个元素  this.customers[:index]第一个切片的意思是：从0的位置开始截取到index，
   this.customers = append(this.customers[:index],this.customers[index+1:]...)
                        //this.customers[index+1:]... 从index+1 到后面全部写完
                        return true
}

//修改客户
func (this *CustomerService)Update(id int ) bool {

}

//根据ID查找客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) FindById (id int)int{

    //返回的-1说明没有找到
	index := -1

	//遍历this.Customer 切片    （找到的就进来）
	for i :=0;i< len(this.customers);i++{
		if this.customers[i].Id == id{
			//找到
			index = i
		}
	}
	return index

}