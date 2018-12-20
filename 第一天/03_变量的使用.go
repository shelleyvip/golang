package main       //必须要导入一个包
import "fmt"

func main(){       //导入包必须要使用
//var 变量在程序运行期间是可以改变的量
//1.声明格式：var 变量名 类型 变量声明了必须要使用
//2.只是声明没有初始化的变量默认值为：0
//3.同一个{}里 声明的变量名是唯一的（不能重名）
var a int
fmt.Println("a=",a)
//可以声明多个变量
var c,b int
fmt.Println("c=",c,"b=",b)
//变量同时赋值
var s int
  s = 10
fmt.Println("s=",s)

  //变量初始化，声明变量时 同时赋值
  var d int = 10 ///变量初始化，声明变量时 同时赋值（一步到位）
  d = 20          //赋值；先声明后赋值
  fmt.Println("d=",d)

  ///自动推倒类型 必须初始化 它是通过初始化的值来推导类型
  e := 88
  fmt.Printf("e type is %T\n",e)
}