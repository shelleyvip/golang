package  main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//编写一个函数接收两个文件的路径

func CopyFile(dstFileName string,srcFileName string)(writtne int64,err error)  {
	srcFile,err := os.Open(srcFileName)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
	}
	defer  srcFile.Close()
	//通过srcFile 获取reader
	reader := bufio.NewReader(srcFile)

	//打开dasFileName                         //判断dasFile文件是否存在 os.O_WRONLY //不存在则创建os.O_CREATE
	dstFile,err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,666)
      if  err != nil{
      	fmt.Printf("open file err=%v\n",err)
      	return
	  }
	//通过dasFile 获取writer
	writne := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writne,reader)

}

func main() {
     srcFile := "/Users/golang/desktop/ascii.jpg" //源文件
     dstFile := "/Users/golang/work/go_path/src/golang/acs.jpg"  //目标文件

     _,err := CopyFile(dstFile,srcFile)
     if err == nil{
		 fmt.Println("copy complete")

	 }else {
		 fmt.Printf("copy file err =%v\n",err)

	 }
     }
