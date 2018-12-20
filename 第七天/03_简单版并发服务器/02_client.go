package main
import (
	"fmt"
	"io"
	"net"
	"os"
)
//使用方法：先服务 在客户端
//
	func main(){
	//主动拨号连接服务器
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Dial err=",err)
		return
	}
	defer func() { //main调用完毕，关闭连接
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//新建一个协程
	//接收服务器回复的数据
	go func() {
		//切片缓冲
		buf := make([]byte,1024)
		for{
			n,err := conn.Read(buf)//接收服务器的请求
			if err != nil{
				fmt.Println("conn.Read err=",err)
				return
			}
			fmt.Println("buf=",string(buf[:n]))//打印接收到的内容，有多少读多少
		}
	}()
	//从键盘输入的内容给服务器发送
	str := make([]byte,1024)
	for{    //用系统包os下面的stdin的方法去读取（标准输入的意思）
		n,err1 := os.Stdin.Read(str)//从键盘读取 放在str
		if err1 != nil{
			if err1 == io.EOF{
				fmt.Println("os.Stdin. err=",err)
			}
			return
		}
		//把输入的内容给服务器发送
		conn.Write(str[:n])
	}
}
