package main

import "fmt"

func swap(a,b string)(string,string)  {
	return a,b
	
}

func main() {
      x,y := swap("hello","world")
      fmt.Println(x,y)
}
