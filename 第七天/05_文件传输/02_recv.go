package main
import (
	"fmt"
	"io"
	"net"
	"os"
)
//使用方法：build 或者run
//先编译成两个可执行文件，直接./运行（先recv等待接收）在运行./send 后面跟上绝对路径的有效文件
func main() {
	receiver()
}
func receiver() {
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close() //不要忘记关闭

	conn, err := listener.Accept() //创建用于通信的socket
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close() //不要忘记关闭

	buf := make([]byte, 4096)
	n, err := conn.Read(buf) //接收发送端发来的文件名
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	filename := string(buf[:n]) //转成字符串

	conn.Write([]byte("ok")) //回发“ok”信号给客户端

	receiveFilm(filename, conn) //接收和保存文件由该函数来完成
}
func receiveFilm(filename string, conn net.Conn) {
	fl, err := os.Create("./" + filename) //创建新文件（需要指定 保存路径 和 文件名）
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer fl.Close() //不要忘记关闭

	buf := make([]byte, 4096)
	for { //conn是一个文件，里面缓存了发送端发送的所有内容，所以需要用for循环来判断是否读完
		n, err := conn.Read(buf) //读取到buf缓冲区中
		if err != nil {
			if err == io.EOF { //io.EOF是文件末尾，当err读到它，代表已经读完
				fmt.Println("read finish 接受完毕✅")
			} else {
				fmt.Println("conn.Read err:", err)
			}
			return
		}

		fl.Write(buf[:n]) //把读到的内容写到新创建的文件中，读多少写多少
	}
}

