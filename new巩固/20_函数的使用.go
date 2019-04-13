package main

import (
	"fmt"
	"os"
	"strconv"
)

func Add(m,n int)int  {
	return m *n
	
}

func Sub(m,n int)int  {
	return m / n
	
}

func main()  {
	funcmap := map[string]func(int,int)int{"*" :Add,"/":Sub}
	m,_ := strconv.Atoi(os.Args[1])
	n,_ := strconv.Atoi(os.Args[2])

	f := funcmap[os.Args[2]]
	if f != nil{
		fmt.Println(f(m,n))
	}
	
}
