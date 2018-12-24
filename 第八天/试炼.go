package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start,end int
	fmt.Printf("请输入起始页(>= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>= start):")
	fmt.Scan(&end)
	
	Woking(start,end)
}
func Woking(start,end int)  {
	fmt.Printf("正在爬去%d到%d的页面",start,end)
	page := make(chan int)
	for i :=start;i<= end;i++{
		Spiderpage1(i,page)
	}
	for i := start;i <= end;i++{
		fmt.Printf("第%d个页面爬取完成",<-page)
	}
}
func Spiderpage1(i int,page chan int)  {
	url := "https://tieba.baidu.com/f?kw=%E4%B8%89%E5%92%8C%E5%A4%A7%E7%A5%9E&ie=utf-8&pn=" +
		strconv.Itoa((i-1)*50)
	result,err := HttpGet1(url)
	if err != nil{
		fmt.Println(" HttpGet err",err)
		return
	}
	f,err := os.Create("第"+strconv.Itoa(i)+"页"+".html")
	if err != nil{
		fmt.Println("os.Create err",err)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}
func HttpGet1(url string)(result string,err error)  {
	resp,err1 := http.Get(url)
	if err1 != nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte,4096)
	for{
		n,err2 := resp.Body.Read(buf)
		if n == 0{
			break
		}
		if err2 != nil && err2 == io.EOF{
           return
		}
		result += string(buf[:n])
	}
	return
}