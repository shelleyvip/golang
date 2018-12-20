package main

import "fmt"

func main(){
	// 初始化字典有两个参数 一个 键key 和  值 value
	zhou := map[int]string{1:"米露",2:"你哈皮",3:"朝阳",4:"周雪莉",5:"北京"}
//	fmt.Println("zhou=",zhou)
///通过键"key"(下标也就是说）可以改对应的值
	zhou[6] = "❤️" //通过map 底层自动扩容
	fmt.Println("map[1]=",zhou)
}
