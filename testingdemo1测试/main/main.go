package main

func addUpper(n int)int {
	res := 0
	for i :=1;i<= n;i++{
		res += i
	}
	return res
}

func addUpper2(n int)int {
	res := 0
	for i :=1;i<= n;i++{
		res += i
	}
	return res
}



//func main() {
//	//传统的测试方法  就是在main函数中使用  看看结果是否正确
//	res := addUpper(10)  //从1++..10= 55
//	if res != 55{
//		fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n",res,55)
//	}else{
//		fmt.Println("addUpper正确")
//	}
//
//}
