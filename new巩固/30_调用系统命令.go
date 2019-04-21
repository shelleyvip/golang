package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os/exec"
	"time"
)

func main() {     //exec.Command调用系统命令
	cmd := exec.Command("df","-tH")  //-l是以树桩的形式显示  ls是当前目录下的文件
	out,_ := cmd.StdoutPipe()
	if err := cmd.Start();err != nil{
		log.Fatal(err)
	}
	f := bufio.NewReader(out)
	for{
		line,err := f.ReadString('\n')
		if err != nil{
			break
		}
		fmt.Print(line)
	}
	time.Sleep(time.Hour)
	cmd.Wait()
}