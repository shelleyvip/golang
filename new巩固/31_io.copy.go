package main

import (
	"github.com/labstack/gommon/log"
	"os"
)

func v1()  {
	var f *os.File
	var err error
	if len(os.Args) >1{
		f,err = os.Open(os.Args[1])
		if err != nil{
			log.Fatal(err)
		}
	}else {
		f = os.Stdin
	}
	buf := make([]byte,1024)
  for{
  	n,err := f.Read(buf)
  	if err != nil{
		return
	}
  	os.Stdout.Write(buf[:n])
  }

}
func main()  {
	v1()
}