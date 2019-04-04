package main

import "fmt"

func Add01(a int)int  {
	if a == 1{
		return 1
	}
	return a + Add01(a - 1)

}


func main() {
	var sum int
	sum = Add01(100)
	fmt.Println("sum=",sum)

}

