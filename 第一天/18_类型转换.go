package main
func main(){
	var ch byte
	ch = 'a' ////字符本质上就是整型
	var t int
	t = int (ch)///把ch的值转 取出来后转换给int 在给t赋值
	println("t=",t)
}//////在这里不能互相转换的类型交不兼容类型
///编译通过说名兼容
////编译不通过说明不兼容
////字符'a'本质上就是整型所以 以上 两种可以相互转换

