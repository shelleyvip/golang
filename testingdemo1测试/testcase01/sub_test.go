package main

import "testing"

//
func TestGetSub(t *testing.T)  {
	//调用
	res := getSub(10,20)
	if res != 30{
		t.Fatalf("getSub执行错误,期望=%v,实际是%v\n",30,res)
	}
}
