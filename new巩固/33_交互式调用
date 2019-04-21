package main

import (
	"bufio"
	"fmt"

	"os/exec"
)

func main() {
	cmd := exec.Command("ls","-l")
	out,_ := cmd.StdoutPipe()
	cmd.Start()

	f := bufio.NewReader(out)
	for{
		line,err := f.ReadString('\n')
		if err != nil{
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
}
