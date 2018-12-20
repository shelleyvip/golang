package main

import (
	"fmt"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com")
		resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status=", resp.Status)  //状态
	fmt.Println("StatusCode=", resp.StatusCode)//状态码
	fmt.Println("Header=", resp.Header) //头部信息
	//fmt.Println("Body=", resp.Body)

		//读取Body的内容👇
	buf := make([]byte, 1024*8)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body err", err)
			break
		}
		tmp += string(buf[:n])

	}
	fmt.Println("tmp=",tmp)
}
