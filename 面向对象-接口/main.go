package main

import "fmt"

//定义个一个Bsb接口
type Usb interface {
	Start()
	Stop()
}
//定义一个手机 实现USB接口的方法
type Phone struct {

}
//让phone去实现USB接口的方法
func (p Phone)Start()  {
  fmt.Println("手机开始工作")
}

func (p Phone)Stop(){
	fmt.Println("手机停止工作")
}

//定义一个相机 实现USB的方法
type Camera struct {
	
}
//让Camera去实现USB接口的方法
func (c Camera)Start()  {
  fmt.Println("相机开始工作")
}

func (c Camera)Stop()  {
	fmt.Println("相机结束工作")
}

//计算机
type Computer struct {

} 
//编写一个方法working方法，接收一个USB接口类型的变量
//只要实现了USB接口（所谓实现USB接口 就是指实现了 就是实现了USB接口声明的所有方法）
func (d Computer)Working(usb Usb)  {
	//通过USB接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main()  {
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点 去调working接口里面的USB的方法
	computer.Working(phone)
	computer.Working(camera)


	
}
