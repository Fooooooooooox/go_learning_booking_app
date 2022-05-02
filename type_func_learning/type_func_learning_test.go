package func_type_learning_test

import (
	"fmt"
	"reflect"
	"testing"
)

// func ReturnX(t *testing.T) {
// 	TestX := big_learning.Returnsomething()
// 	t.Log(TestX)
// }

//定义函数类型
type TestHandler func(name string) int

//实现函数类型方法
func (h TestHandler) TestAdd(name string) int {
	return h(name) + 10
}

func TestUsehandler(t *testing.T) {
	//注意要成为函数对象必须显式定义handler类型
	// Handler是前面定义的函数类型 my的类型是handler
	var my TestHandler = func(name string) int {
		return len(name)
	}
	fmt.Println(my("taozs"))
	fmt.Println(reflect.TypeOf(my("taozs")))
	fmt.Println(my.TestAdd("taozs")) //调用函数对象的方法
	fmt.Println(reflect.TypeOf(my("taozs")))
}
