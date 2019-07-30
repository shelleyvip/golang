package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file,err := os.Open("/Users/golang/desktop/复习秘籍.html")
	if err != nil{
		fmt.Println("open file err",err)
	}
	//当函数退出时要及时退出file
	defer file.Close()     //要及时关闭file句柄 否则会有内存泄露



	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for{
      str,err := reader.ReadString('\n')//按行去读 读到一个换行就结束
      if err == io.EOF{  //io.EOF 表示文件的末尾
      	break
	  }
      //输出内容
      fmt.Print(str)
	}
  fmt.Println("文件读取结束......a ")
}
