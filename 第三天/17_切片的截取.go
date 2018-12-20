package main

import "fmt"

func main(){
	zhou :=[]int{1,2,3,4,5,6,7,8,9}
	//1：//不指定容量和长度
	s1 := [...]int{}
	fmt.Println("s1=",s1)
	fmt.Printf("len=%d,cap=%d\n",len(s1),cap(s1))



	//2：截取中间某个长度和容量 len：6-3 cap：7-3
	s2 :=zhou[3:6:7]
	fmt.Println("s2=",s2)
	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))

	//常用
	s3 := zhou [:6] //从0开始去掉6个元素，容量也是6（也就是:6  是打印6以前：6）
	fmt.Println("s3=",s3)
	fmt.Printf("s3:len=%d,cap=%d\n",len(s3),cap(s3))


	s4 :=zhou[3:]//从下标3开始算起，直到结尾 （3：三以后）
	fmt.Println("s4=",s4)
	fmt.Printf("s4:len=%d,cap=%d",len(s4),cap(s4))

}