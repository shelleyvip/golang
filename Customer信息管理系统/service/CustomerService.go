package service

import "golang/Customer信息管理系统/model"

//该CustomerService,完成对Customer的操作
//增删改查
type CustomerService struct {
    customer []model.Customer
    //声明一个字段表示当前切片含有多少个客户
    //该字段后面，还可以作为新客户的id
    customerNum int
}
//编写一个方法可以返回*CustomerService
func NewCustomerService() *CustomerService {
    //为了能看到有客户在切片中，我们初始化一个客户
    customerService := &CustomerService{}
    customerService.customerNum = 1
    customer := model.Newcustomer(1,"张三","妖怪",20,"123","zs@sohu.com")
    customerService.customer = append(customerService.customer,customer)
    return customerService
}

