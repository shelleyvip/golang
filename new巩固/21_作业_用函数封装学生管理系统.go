package main
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name string
	Id int
}

func add(args []string) error {
	fmt.Println("call add")
	fmt.Println("args",args)
	//name := args[0]
	//id := args[1]
  return nil
}

func list(args []string)error {
	return errors.New("unimplemention")
}
func main() {
		actionmap := map[string]func([]string)error{"add":add,}
		f := bufio.NewReader(os.Stdin)
		for{
			//标识符/提示符
			fmt.Print(">")
			//先读取一行 按换行来分隔
			line, _ := f.ReadString('\n')
			//👇去掉两端的空格和换行
			line = strings.TrimSpace(line)
			//👇按空格分隔字符 得到字符串列表
			args := strings.Fields(line)
			if len(args) == 0{
				continue
			}
			//获取命令和参数列表
			cmd := args[0]
			args = args[1:]

			//获取命令所对应的函数
			actionfuanc := actionmap[cmd]
			//如果没找到相应的函数👇 actionfuanc == nil 那么就提示一下 fmt.Println("bad cmd",cmd) 继续执行continue
			if actionfuanc == nil{
				fmt.Println("bad cmd",cmd)
				continue
			}
			fmt.Println("funcd cmdfundc")
			continue

			err := actionfuanc(args)
			if err != nil{
				fmt.Printf("execute action %s error :%s\n",cmd,err)
				continue
			}

		}

	}