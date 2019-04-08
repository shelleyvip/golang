package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	f,err := os.Create("fmt.txt")
	if err != nil{
		log.Fatal(err)
	}
	//Fprintln 把写的内容放进在文件里
	fmt.Fprint(f,"hello")   //Fprint 把写的内容放进在文件里
	fmt.Fprintln(f,"helloln")
	s := "hello world"
	d := 123
	fmt.Fprintf(f,"s=%s d=%d",s,d)
	f.Close()
}