package main

import (
	"fmt"
)

func main()  {
	var xueli int
	fmt.Println("请输入你的考试的分数: ")
	fmt.Scan(&xueli)

	switch  {
	case xueli  > 90:
		fmt.Println("优秀")
	case xueli  > 80:
		fmt.Println("良好")
	case xueli > 70:
		fmt.Println("一般")
	case xueli <= 60:
		fmt.Println("及格")

	default :
		fmt.Println("不合格，请重新复读，并且考试去")
	}

}

