package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go",Myhand)
	//在指定的地址进行监听
	http.ListenAndServe("127.0.0.1:8000",nil)
	}

//w 是给客户端回复写入数据
//r 是读客户端发送的数据
func Myhand(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello world !!!")
}



//上面是一个web服务器给网址提供服务的
