package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个已经存在的文件，将原来的内容覆盖诚心的十句话"你好 大米露"
   filepath := "/Users/golang/work/go_path/src/golang/filedemo/shelley.txt"
                                                  //os.O_TRUNC, 这个是把原先的内容先晴空覆盖掉
       file,err := os.OpenFile(filepath,os.O_WRONLY |os.O_TRUNC,0666)
	if err != nil{
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()


	str := "你好 大米露\n"
	//写入时使用低缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0;i<10;i++{
		writer.WriteString(str)
	}

	//因为writer时代缓存的 因此要调用writerString方法时，其实
	//内容是先写到缓存的，所以需要调用flush方法，将缓存的数据真正写到文件中，否则文件中没有数据
	writer.Flush()

}
