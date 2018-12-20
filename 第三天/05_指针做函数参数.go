package main

import "fmt"

func swap(p1,p2 *int) {
	*p1,*p2 = *p2,*p1
	fmt.Printf("swap: a=%d b=%d\n",*p1,*p2)

	//fmt.Printf("swap: a=%d b=%d\n",a,b)
}
func main()  {
	a,b := 10,20
	swap(&a,&b) // 地址传递 取地址
	fmt.Printf("main: a=%d b=%d",a,b)

}


