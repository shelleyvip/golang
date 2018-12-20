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

func xue (zhou1 *Shelley){
	zhou1.qq = 999
	fmt.Println("xue=",*zhou1)
}
////指针传递是可以形参改变实参的

func main() {
	///相对等的类型
	zhou1 := Shelley{name: "周雪莉", age: 22, sex: "中", qq: 128, addr: "朝阳"}
	xue(&zhou1)//条件 指针地址传递（引用传递）形参可以改实参
	fmt.Println("zhou1=",zhou1)
}
