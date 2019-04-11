package main

import "fmt"

func swap(a,b string)(string,string)  {
	return a,b
	
}

func split(sub int)(x,y int)  {
	x = sub / 10
	y = sub * 10
    return
}

func main() {
      x,y := swap("hello","world")
      fmt.Println(x,y)
      fmt.Println(split(17))
}
