package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	s,err := os.OpenFile("a.txt",os.O_CREATE|os.O_RDWR,0644)
	if err != nil{
		log.Fatal(err)
	}
	//s.WriteString("hello\n")
	//s.Seek(1,os.SEEK_SET)
	//s.WriteString("$$$$")
	//s.Close()


	//	这种方式属于块读取
	buf := make([]byte,1024)
	s.Read(buf)
	fmt.Println(string(buf))
	s.Close()
}
