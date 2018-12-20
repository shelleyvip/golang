package main

import (
	"fmt"
	"net"
)

func main(){
	listener,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Listen err=",err)
		return
	}
	defer func() {
		err = listener.Close()
		if err != nil{
			fmt.Println("err=",err)
		}
	}()

	conn,err := listener.Accept()//接收
	if err != nil{
		fmt.Println("listener.Accept err",err)
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil{
			fmt.Println("err=",err)
		}
	}()


	buf := make([]byte,1024)
	for{
		n,err := conn.Read(buf) //
		if err != nil{
				return
			}
			fmt.Println("buf=",string(buf[:n]))
		}
	}
