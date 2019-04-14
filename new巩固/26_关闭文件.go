package main

import (
	"github.com/labstack/gommon/log"
	"os"
)

func Print()  {
   f,err := os.Open("a.txt")
   if err != nil{
   	log.Fatal(err)
   }
   defer f.Close()
}
func main() {
	print()
}
