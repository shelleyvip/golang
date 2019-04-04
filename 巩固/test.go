package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	var flag bool
	for num := 1;num <8000;num++{
        flag = true
        for i :=2;i<num;i++{
        	if num %i == 0{
        		flag = false
        		break
			}
        	if flag{

			}
        	//end := time.Now().Unix()
        	//fmt.Println("普通的方法耗时",end-start)
		}
		//fmt.Println("普通的方法耗时",end-start)

	}
	end := time.Now().Unix()

	fmt.Println("普通的方法耗时",end-start)

}
