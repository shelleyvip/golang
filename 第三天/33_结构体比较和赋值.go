package main
import "fmt"
//定义以个结构体 用type 和struct
type Shelley struct {
	name string
	age int
	sex string    ///定义类型
	qq int
	addr string
}

func main() {
	///相对等的类型
	zhou1 := Shelley{name: "周雪莉", age: 22, sex: "中", qq: 128, addr: "朝阳"}
	zhou2 := Shelley{name: "周雪莉", age: 22, sex: "中", qq: 128, addr: "朝阳"}
	zhou3 := Shelley{name: "周雪莉", age: 25, sex: "中", qq: 128, addr: "朝阳"}

	fmt.Println("zhou1==zhou2",zhou1==zhou2)
	fmt.Println("zhou1==zhou3",zhou1==zhou3)
//同类型的结构体可以赋值
 var nannan Shelley
	 nannan =  zhou1
	fmt.Println("nannan=",nannan)
}