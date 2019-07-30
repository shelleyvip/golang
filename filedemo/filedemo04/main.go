package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件 并写入五句话 hello Gardon
	//打开文件
	filepath := "/Users/golang/work/go_path/src/golang/filedemo/shelley.txt"
	file,err := os.OpenFile(filepath,os.O_WRONLY |os.O_CREATE,0666)
	if err != nil{
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()


	//准备写入五句话hello Gardon
	str := "hello Gardon\n"
	//写入时使用低缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0;i<5;i++{
		writer.WriteString(str)
	}

	//因为writer时代缓存的 因此要调用writerString方法时，其实
	//内容是先写到缓存的，所以需要调用flush方法，将缓存的数据真正写到文件中，否则文件中没有数据
	writer.Flush()

}


