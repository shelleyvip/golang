package main

import "fmt"

func main() {
	s :=[]int{1,2,3,4,5,6,7}
	s = s [1:4]
	fmt.Println(s)

	s =s[:2]
	fmt.Println(s)

	s=s[1:]
	fmt.Println(s)
   a :=[...]int{2,3,5,7,9,11}
   s = a[:0]
   fmt.Println(s,len(s),cap(s))
   s =s[:4]
	fmt.Println(s,len(s),cap(s))
  s=s [:2]
	fmt.Println(s,len(s),cap(s))


  var  s1[]int  //空切片为nil
  if s1 == nil{
  	fmt.Println("nil")
  }
  fmt.Println(len(s1))
  s1 = a[:0]
  fmt.Println(s1 == nil)


}
