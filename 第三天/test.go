package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)

func countline(url string)  {
	fmt.Println("代码读取完毕")
	fp,err:=os.Open(url)
	//打开.go文件
	if err!=nil {
		log.Fatal(err)
	}
	defer fp.Close()
	//使用bufio包下的newscannner方法来数行
	input:=bufio.NewScanner(fp)
	for input.Scan(){
		//每多一行则全局变量a加一
		a++
	}
}

func findgo(url string)  {
	dp,err:=os.OpenFile(url,os.O_RDONLY,os.ModeDir)
	//打开根目录，并扫描其中所有文件
	if err!=nil {
		log.Fatal(err)
		return
	}
	defer dp.Close()
	fileinfo,err:=dp.Readdir(-1)
	if err!=nil {
		log.Fatal(err)
		return
	}
	for _,info:=range fileinfo{
		//遍历所有文件，若是文件夹，则将此文件夹目录作为参数，重复调用此函数
		if info.IsDir() {
			findgo(url+"/"+info.Name())
		}else if strings.HasSuffix(info.Name(),".go") {
			//若是.go文件，则调用数行countline函数
			countline(url+"/"+info.Name())
		}
	}
}
var a int=0
func main()  {
	var url string
	fmt.Println("请输入您go代码所在根目录：")
	fmt.Scan(&url)
	//此函数用于寻找用户提供的目录下的所有.go文件
	findgo(url)
	score:=float32(a)/100000.0*100.0
	//结束后进行人性化输出
	fmt.Printf("您已经敲了%d行代码\n十万行代码的目标已经完成%.2f%%了，继续努力哦!!!",a,score)
}

