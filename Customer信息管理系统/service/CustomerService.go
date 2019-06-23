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