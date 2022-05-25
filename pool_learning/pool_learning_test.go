package poollearning_test

// pool是go官方提供的池子
// 关于他有三个方法需要掌握的：
// New、Put、Get

// 还补充一个相关知识：GC
// GC就是garbage collector
// runtime.GC()就是用来垃圾回收 释放对象占用的内存的 一般是由系统自动调用

import (
	"fmt"
	"sync"
	"testing"
)

// New是用来定义pool中存放的东西 一个池子只能存放一种类型
var strPool = sync.Pool{
	New: func() interface{} {
		return "test str"
	},
}

func TestPool(t *testing.T) {
	str := strPool.Get()
	fmt.Println(str)
	strPool.Put(str)
}

var intpool = sync.Pool{
	New: func() interface{} {
		return "123"
	},
}

func TestIntpool(t *testing.T) {
	getsth := intpool.Get().(string)
	fmt.Println(getsth)

	// 这里会给你一个提示：argument should be pointer-like to avoid allocations (SA6002)
	// 意思是推荐你用指针类型的数据
	intpool.Put("321")

	getsth2 := intpool.Get().(string)
	fmt.Println(getsth2)
}
