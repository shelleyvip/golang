package main
import "fmt"
type Zhouxueli interface {
	Shelley()
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

func (pointer *Huzailing)Shelley(){
	fmt.Printf("Huzailing[%s] Shelley\n ", *pointer)
}

func (pointer *Loop)Shelley(){
	fmt.Printf("Loop[%d %s] Shelley\n ", pointer.id,pointer.addr)
}
func (pointer  *Connit) Shelley(){
	fmt.Printf("Connit[%v %v %v]\n",pointer.name,pointer.age,pointer.sex)
}
//在这里定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同的表现，多态
func Addr01 (i Zhouxueli)  {
	i.Shelley()

}
func main()  {
	a := &Connit{"周🍐",25,'m'}
	b :=  &Loop{29,"北京朝阳"}
	var c Huzailing = "hell"

    //调用统一函数，可以有不同表现，这就是多态，多种形态
	Addr01(a)
	Addr01(b)
	Addr01(&c)

	//第二种方式😁自研发
		///a.Shelley()
		//b.Shelley()
		//c.Shelley()
}