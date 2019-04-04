package main

import "fmt"

func intNum(intChan chan int)  {
	for i:=1;i<8000;i++{
		intChan <- i
	}
	close(intChan)
}
func PrimeNum (intChan chan int,primeChan chan int,exitChan chan bool)  {
	var flag bool
	for{
		num,ok := <-intChan
		if !ok{
			break
		}


		flag = true
		for i :=2;i<num;i++{
			if num %i == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan <- num
		}

	}
	fmt.Println("有一个PrimeNum 协程因为取不到数据退出了")
	exitChan <- true

}
func main() {
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)
	exitChan := make(chan bool,4)

	go intNum(intChan)


	go func() {
		for i :=0;i<4;i++{
			go PrimeNum(intChan,primeChan,exitChan)
		}
		for i :=0;i<4;i++{
			<- exitChan
		}
		close(primeChan)
	}()

	for{
		result,ok := <- primeChan
		if !ok{
			break
		}else {
			fmt.Printf("素数=%d\n",result)
		}

	}



}