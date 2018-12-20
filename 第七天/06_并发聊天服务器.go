package main

import (
	"fmt"
	"net"
)

func main()  {
	listener,err := net.Listen("tcp",":8000")
	if err != nil{
		fmt.Println("net.Listen err",err)
		return
	}
	defer listener.Close()

	go Manager()

	for{
		cnnn,err := listener.Accept()
		if err != nil{
			fmt.Println("listener.Accept err",err)
			continue
		}
		go Handleconn(conn)

	}
}
func Handleconn(conn net.Conn)  {
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	cli := Celient(make(chan string),cliAddr,cliAddr)

	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli,conn)

	messaage <- "[" + cli.Addr+']'+c

}

type clint struct {
	C chan string
	Name string
	Addr string
}
var onlineMap = make(map[string]clint)
var messaage = make(chan string)

func Manager()  {
	for{
		msg := <-messaage

		for _,cli := range onlineMap{
			cli.C <= msg
		}
	}

}
func WriteMsgToClient(cli clint,conn net.Conn){
	for msg:= range cli.C{
		conn.Write([]byte(msg+"\n"))

	}
}

func MakeMsg(cli Celient,msg string)(buf string){


}
