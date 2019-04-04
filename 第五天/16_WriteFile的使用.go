package main

import (
	"fmt"
	"os"
)

func WriteFile(path string){
	//打开文件，新建文件
	//这里有两个文件一个是文件另一个是错误
	f,err := os.Create(path)
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	//使用完毕需，要关闭文件
	defer f.Close()

	var buf string
	for i :=0;i<10;i++{
		buf = fmt.Sprintf("i=%d\n",i)
		fmt.Println("buf=",buf)
		          //WriteString(buf)系统内置函数写入string
		n,err := f.WriteString(buf)
		if err != nil{
			fmt.Println("err=",err)
			return
		}
		fmt.Println("n=",n)
	}
}
/// 文件路径  /Users/golang/Work/go_path/工程管理不同目录调用/第五天/
//cd  进入文件夹    ； cd  /Users/golang/Work/go_path/工程管理不同目录调用/第五天/
//ls 查看文件
//cat 查看文件的的；比如：  cat 01_试炼.go
func main(){
	//先创建一个路径 ./代表当前路径  //后面跟上要创建的名称xxx.txt
	//path := "./读文件.txt" //相对路径读取
	path := "/Users/golang/Work/读文件.txt"   //////绝对路径读取
	WriteFile(path)///把这个文件传进来
}
