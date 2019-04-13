package main

import "fmt"

func addn(m int)func(int)int  {
	return func(n int)int{
		return m + n
	}
}

func main() {
	f := addn(100)
	fmt.Println(f(20))
}