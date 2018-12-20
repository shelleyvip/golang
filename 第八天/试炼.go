package main
import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)
func main() {
	var start, end int
	fmt.Printf("请输入起始页(>= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>= 终止页）:")
	fmt.Scan(&end)
	Dowork(start, end) //工作函数
}
func Dowork(start, end int) {
	fmt.Printf("正在爬取第%d页面%d \n", start, end)
	//page := make(chan int)
	sync := make(chan bool)
	for i := start; i <= end; i++ {
		//定义一个函数爬取主页面
		go SpiderPape(i, sync)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("正在爬取出%d到%d%v \n", start, end, <-sync)
	}
}
func SpiderPape(i int, sync chan bool) {
	//明确爬取的Url
	//https://www.pengfu.com/xiaohua_1.html
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页面%s \n", i, url)
	//开始爬取内
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	// fmt.Println("r=",result)
	//取<h1 class="dp-b"><a href="一个段子URL连接
	//解释表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls =",joyUrls)
	//取网址
	//第一个返回的是下标，第二个返回内容
	for _, data := range joyUrls {
		fmt.Println("url=", data[1])
	}
	// 2）爬 将网站的所有的内容全部爬下来
	//把内容放在文件夹
	FileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(FileName)
	if err1 != nil {
		fmt.Println("os.Create", err1)
		return
	}
	_, err = f.WriteString(result) //写内容
	if err != nil {
		fmt.Println("err", err)
	}
	err = f.Close() //关闭
	if err != nil {
		fmt.Println("err", err)
	}
	sync <- true
	//page <- i
}
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //发送get请求
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4) //读取网页内容
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n]) //累加读取的内容
	}
	return
}