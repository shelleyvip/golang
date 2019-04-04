
package main

import "fmt"

func and(a int){
	if a == 1 {
		fmt.Println("a=",a)
		return
	}

	and (a -1)
	fmt.Println("a=",a)
}

func main() {
	and(10)
	fmt.Println("main")
}


