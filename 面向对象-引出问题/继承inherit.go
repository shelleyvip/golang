package main

import "fmt"

//继承的优点
//代码的复用性提高了，维护性和扩展性也提高了
//可以一个方法N个使用

//编写一个学生考试系统
type Student struct {
	Name string
	Age int       //共有属性提取出来
	Score int
}
//小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}
//大学生
type Graduate struct {
	Student    //嵌入了Student匿名结构体
}

//将pupil和Graduate共有的字段 也绑定到*Student
func (stu *Student)ShowInfo()  {  //显示信息
	fmt.Printf("学生名字=%v 学生年龄=%v,学生分数=%v\n",stu.Name,stu.Age,stu.Score)
}
//设置成绩
func (stu *Student)SetScore(Score int)  {
	//业务判断
	stu.Score = Score
}
//给Student增加一个方法 让pupil和Graduate都可以使用该方法
func (stu * Student)GetSum(n1 int,n2 int)int  {
	return n1 + n2
	
}
//显示考试状态
//这是Pupil结构体特有的方法 保留
func (p *Pupil)testing()  {
	fmt.Println("小学生正在考试....")
}

//这是Graduate结构体特有的方法 保留
func (p *Graduate)testing()  {
	fmt.Println("大学生正在考试....")
}

func main() {
	//当我们对结构体嵌套了匿名结构体 使用方法会发生变化
	pupil := Pupil{}
    pupil.Student.Name = "小铁战士"
    pupil.Student.Age = 8
    pupil.testing()
    pupil.Student.SetScore(90)
    pupil.Student.ShowInfo()
    fmt.Println("result=",pupil.GetSum(10,20))


    graduate := Graduate{}      //定义一个变量去接收 调用Graduate的内容
    graduate.Student.Name = "胡老师"  //赋值
    graduate.Student.Age = 27     //赋值
    graduate.testing()    //执行testing内容信息
    graduate.Student.SetScore(288)  //设置分数
    graduate.Student.ShowInfo()  //获取信息
	fmt.Println("result=",graduate.GetSum(50,20))

}



