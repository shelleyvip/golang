package main
func main(){
	//创建一个channel，双向的 正常的
	ch := make(chan int)

	//双向channel 能隐式转换为单向channel
	var wreleCh  chan<- int = ch //只能写，不能读
	var readch <-chan int = ch //只能读，不能写

	wreleCh <- 666//写

	<- readch //读



	//单项无法转换为双向
	//var ch2 chan int = wreleCh
}
