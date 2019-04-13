package main

import (
	"fmt"
	"strings"
)

func toupper(s string)string  {
	return strings.Map(func(r rune) rune {
		return r - ('a'- 'A') //小写转大写
	},s)
}

//第二种方式
func swap(r rune)rune  {
	fmt.Printf("r %c\n",r)
      return r
}

func main() {
	fmt.Println(toupper("hello"))

	s := strings.Map(func(r rune) rune {
		return r - 32
	},"hello")
	fmt.Println(s)
}
