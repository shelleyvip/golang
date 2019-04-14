package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
)

func read(f *os.File)(string,error)  {
	var total []byte
	buf := make([]byte,1024)
	for{
		n,err := f.Read(buf)
		if err  == io.EOF{
			break
		}
		if err != nil{
			return "",err
		}
		total = append(total,buf[:n]...)
	}
	return string(total),nil

}

func main()  {
    f,err := os.Open("a.txt")
    if err != nil{
    	log.Fatalf("open error:%v\n ",err)
	}
    s,err := read(f)
    if err != nil{
    	log.Fatal("read error:%v\n",err)
	}
    fmt.Println(s)
}
