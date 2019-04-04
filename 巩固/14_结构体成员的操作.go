package main
import "fmt"

type Person struct {
	name string
	sex byte
	age int

}
type Student struct {
	Person
	id   int
	addr string
}

func main()  {
	s1 := Student{Person{"shlley",'a',25},2,"bj"}
	s1.sex = 'n'
	s1.name = "zhouxueli"
	s1.id = 666
	s1.age =18
	s1.addr = "bj"


	s1.Person = Person{	"米露",'m',3}
	fmt.Println()
	fmt.Println(s1.age,s1.addr,s1.id,s1.name,s1.sex)

}


