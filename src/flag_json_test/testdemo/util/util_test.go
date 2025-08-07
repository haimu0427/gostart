package main

import (
	"testing" // 导入测试框架
)

// 编写测试用例
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper(10) 错误，期望值=%v, 实际值=%v", 55, res)
		return
	}
	t.Logf("addUpper(10) 正确，期望值=%v, 实际值=%v", 55, res)
}
func TestHello(t *testing.T) {
	t.Log("Hello, World!")
}
