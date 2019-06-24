package model

import "fmt"

//声明一个customer结构体表示客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string

}

//使用工厂模式，返回一个customer的实例
func Newcustomer(id int,name string,gender string,age int,
	phone string,email string) Customer {
		return Customer{
			Id : id,
			Name : name,
			Gender : gender,
			Age:age,
			Phone:phone,
			Email:email,
		}
	
}

//返回用户的信息 格式化后的
func (this Customer)GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,
		this.Name,this.Gender,this.Age, this.Phone,this.Email)
	return info
	
}

