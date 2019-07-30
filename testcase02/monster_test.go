package monster

import (
	"testing"
)

//创建一个测试用例
func TestStore(t *testing.T)  { //*testing.T 是测试的固定写法

//先创建一个Monster实例
    monster := &Monster{
    	Name : "红孩儿",
    	Age : 88,
      Skill : "必杀技 吐火",
	}
    res := monster.Store()
    if !res{
    	t.Fatalf("monster.Store()错误 希望为=%v 实际为%v",true,res)
	}
    t.Logf("monster.Store() 测试成功") //正确的日志
}

func TestReStore(t *testing.T) {
	//先创建一个Monster的实例   不需要指定字段的值
	var monster= &Monster{}
	res := monster.ReStore()
		if !res {
			t.Fatalf("monster.Store()错误 希望为=%v 实际为%v", true, res)
		}
		//进一步判断
		if monster.Name != "红孩儿" {
			t.Fatalf("monster.Store()错误 希望为=%v 实际为=%v", "红孩儿", monster.Name)
		}
		t.Logf("monster.Re	Store() 测试成功") //正确的日志
	}
	//t.Logf("monster.Re	Store() 测试成功") //正确的日志
