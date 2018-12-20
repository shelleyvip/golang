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

func xue (zhou1 Shelley){
 zhou1.qq = 999
 fmt.Println("xue=",zhou1)
}
  ////值传递形参是无法改变实参的
  ///只是里面变了，外面main函数还是一样的

func main() {
	///相对等的类型
	zhou1 := Shelley{name: "周雪莉", age: 22, sex: "中", qq: 128, addr: "朝阳"}
	xue(zhou1)//条件 值传递  形参里面的值 无法改变实参
	fmt.Println("zhou1=",zhou1)
}