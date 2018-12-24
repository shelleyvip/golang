package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)
//爬虫概念：
    //访问web服务器，获取指定信息的一段程序
//工作流程
    //1：明确目标URl
    //2：发送请求 获取应答数据包
    //3：保存过滤数据，提取有用的信息
    //4 :使用 分析得到的数据
//百度贴吧爬虫的实现：操作流程如下
        //1：提示用户指定起始，终止页。创建Dowork()函数
        //2：使用start，end循环爬取每一页数据
        //3：获取每一页的URL---下一页=前一页+50
        //4：封装 实现HttpGet（）函数，爬取一个网页的数据内容，通过result 返回
      // http.Get/resp.Body.Close/buf:=make(1024*4)/for{resp.Body.Read(buf)/result += string(buf[:n])return
        //5:创建 .html文件。使用循环因子i命名
        //6:result写入 文件WrieString(result).f.close 不推荐使用 defer

//爬去网页内
func HttpGet3(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误，传出给调用者，（调用者也就是HttpGet）
		return
	}
	defer resp.Body.Close()

	//循环读取数据，传出给调用者
	buf := make([]byte, 1024*4)
	for {
		n, err1 := resp.Body.Read(buf)
		if n == 0 { //读取结束或者出问题
		if err1 == io.EOF{
			break
		}
			fmt.Println("resp.Body.Read err", err1)
		}
		//累加每一次循环 读到的 buf数据，存入result 一次性返回。
		result += string(buf[:n])

	}
	return
}
//爬取页面操作
func DoworkB(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//1)循环爬取每一页数据 明确目标（要知道你准备在哪个范围或者哪个网站去搜索）
	for i := start; i < end; i++ { ///第二步✌️然后根据这个范围来个循环来生成网页 网址，
		url := "https://tieba.baidu.com/f?kw=%E4%B8%89%E5%92%8C%E5%A4%A7%E7%A5%9E&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50) //// 由于它从一开始输入，所以要-1 *50
		fmt.Println("url =", url)

		//2）//把爬取的内容，返回到result 用result去接收
		result, err := HttpGet3(url) //
		if err != nil {
			fmt.Println("HttpGet err", err)
			continue
		}

		//把内容写入到文件
		FileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(FileName)   //第四步 ：：往这个文件写内容
		if err1 != nil {
			fmt.Println("os.Create err1", err1)
			continue
		}

		f.WriteString(result) //写内容
		f.Close()//关闭文件
	}
}
func main() {
	//第一步 首先输入一个范围（指定爬取起始页到终止页 一个1-多少的范围）
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoworkB(start, end)
}