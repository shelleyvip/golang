package main

import "fmt"

func main() {
	func(a int, b string) {
		fmt.Printf("a=%d,b=%s\n", a, b)
	}(100, "雪莉")

}