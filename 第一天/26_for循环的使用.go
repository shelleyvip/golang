package main

import "fmt"

func main(){
	xueli := 0

	for  i := 0; i <= 10000; i++  {
		xueli = xueli + i
		fmt.Println("xueli =",xueli)
	}
}
