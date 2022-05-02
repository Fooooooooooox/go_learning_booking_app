package func_type_learning

import (
	"fmt"
	"reflect"
)

//定义函数类型
type Handler func(name string) int

//实现函数类型方法
func (h Handler) Add(name string) int {
	return h(name) + 10
}

func UseHandler() {
	//注意要成为函数对象必须显式定义handler类型
	// Handler是前面定义的函数类型 my的类型是handler
	var my Handler = func(name string) int {
		return len(name)
	}
	fmt.Println(my("taozs"))
	fmt.Println(reflect.TypeOf(my("taozs")))
	fmt.Println(my.Add("taozs")) //调用函数对象的方法
	fmt.Println(reflect.TypeOf(my("taozs")))
}
