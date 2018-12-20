package main
import "fmt"

//定义接口类型
type method interface {//定义类型用interface
	//方法，只有声明，没有实现。将由别的类型（分自定义类型）实现
	Shelley()//这就是方法
}
type Connit struct {
	name string
	age  int
	sex  byte
}
type Loop struct {
	id int
	addr string
}
type Huzailing string

//在这里接受者都必须是指针//这里是huzailing实现了此方法
func (pointer *Huzailing)Shelley(){//开始引入Shelley的这个接口实现
	fmt.Printf("Huzailing[%s] Shelley\n ", *pointer)
}
//这里是loop实现了此方法
func (pointer *Loop)Shelley(){//开始引入Shelley的这个接口实现
	fmt.Printf("Loop[%d %s] Shelley\n ", pointer.id,pointer.addr)
}

func (pointer  *Connit) Shelley(){//开始引入Shelley的这个接口实现
	fmt.Printf("Connit[%v %v %v]\n",pointer.name,pointer.age,pointer.sex)


}
func main(){
	var s   method  //定义接口类型的变量
	//在这里只要实现了此接口方法的类型，那么这个类型的变量（接受者的类型）就可以给 /s/ 赋值
	a := &Connit{"周🍐",25,'m'}
    s  = a
    s.Shelley()

    t :=  &Loop{88,"bj"}
    s = t
    s.Shelley()

    var str Huzailing = "hello"
    s = &str
   s.Shelley()
}