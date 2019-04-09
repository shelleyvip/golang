package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	f,err := os.Open(".")///.是指当前目录
	if err != nil{
		log.Fatal(err)
	}
	infos,_ :=f.Readdir(-1) //这里的参数是最大读取 -n是所有
	for _,info := range infos{
		fmt.Printf("%v %d %s\n",info.IsDir(),info.Size(),info.Name())//需要获取目录信息的 名字 大小 还有ladir 目录
	}
	f.Close()
}
