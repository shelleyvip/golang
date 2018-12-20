package main

import "fmt"

func main()  {
	shelley := []int{1,2,3,4,5}  //原
	zailing := []int{7,7,7,7,7,7,7}  //目标
	copy(zailing,shelley) //利用copy函数 原覆盖目标的值
	fmt.Printf("zailing=%d\n",zailing)

}
