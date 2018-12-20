package main

import (
	"fmt"
	"net/http"
)

func main()  {

	//注册处理函数，用户连接，自动调用指定的处理函数
	//"/"是根路径的意思
	http.HandleFunc("/",HandConn)

	http.ListenAndServe(":8000",nil)
}
//w 是给客户端回复写入数据
//r 是读客户端发送的数据
func HandConn(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("hello world!"))//给客户端回复数据
fmt.Println("r.Method=请求方法",r.Method)
	fmt.Println("r.url=请url",r.URL)
	fmt.Println("r.Header=头",r.Header)
	fmt.Println("r.host=请主机",r.Host)
	fmt.Println("r.host=body",r.Body)






}
