package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls","-1")
	out,_ := cmd.StdoutPipe()
	buf := make([]byte,1024)
	for{
		_,err :=out.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Printf("erad err %v",err)
			break
		}
		if err := cmd.Start();err != nil{
			log.Fatalf("start error :%v",err)
		}
	}
	fmt.Println(out)
}
