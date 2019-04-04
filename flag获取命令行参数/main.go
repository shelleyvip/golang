package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s","","sss")

var bbb = flag.Bool("n",false,"方法")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if *bbb{
		fmt.Print()
	}


}
