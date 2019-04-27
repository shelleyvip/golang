package model

////定义一个结构体
//type Student struct {
//	Name string
//	Score float64
//}
type student struct {
	Name string
	Score float64
}

//因为student结构体字母为小写，因此我们只能在model 使用
//我们来通过工厂模式来决绝

func NewStudent(n string,s float64)*student  {
	return &student{
		Name : n,
		Score : s,
	}
}
////如果Score字段首字母为小写，则，在其它宝不可以直接调用，我们可以提供一个方法
func (s *student)GetScore()float64  {
	return s.Score //
}
