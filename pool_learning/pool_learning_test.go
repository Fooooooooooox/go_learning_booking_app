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

	"github.com/holiman/uint256"
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

type Stack struct {
	data []uint256.Int
}

func TestStack(t *testing.T) {
	fmt.Println(&Stack{data: make([]uint256.Int, 0, 16)}) //&{[]}
	var S Stack
	S.data = make([]uint256.Int, 0, 16)
	// 0是长度 16是容量 如果没有指定容量，长度=容量
	// go中slice的长度指的是他所包含的元素的个数 容量是他从第一个元素开始数到末尾的个数 长度通过len()来获取 容量通过cap()来获取
	// 什么意思？
	// 试一试
	s := []int{2, 3, 4, 5, 9}
	fmt.Println(s)      //[2 3 4 5 9]
	fmt.Println(len(s)) //5
	fmt.Println(cap(s)) //5

	fmt.Println("-----")
	// 截取slice 让他变为0长度
	s = s[:0]
	fmt.Println(s)      //[]
	fmt.Println(len(s)) //0
	fmt.Println(cap(s)) //5

	// 现在懂了：
	// len是slice当前的长度
	// cap是slice的最大长度
	// Slice: The size specifies the length. The capacity of the slice is
	// equal to its length. A second integer argument may be provided to
	// specify a different capacity; it must be no smaller than the
	// length. For example, make([]int, 0, 10) allocates an underlying array
	// of size 10 and returns a slice of length 0 and capacity 10 that is
	// backed by this underlying array.
	// 第一个参数是len 第二个参数是cap
	// cap要大于len 比如make([]int, 0, 10)指派了一个隐藏的array然后返回了这个array的一个长度为0的slice

	// len和cap有不同的值是因为这个数组会被切割
	// 用make的方法来创建slice
	s2 := make([]uint64, 3, 4)
	s2[0] = 1
	s2[2] = 3
	fmt.Println(s2)
	// 那怎么扩展这个slice的长度len呢？
	s2 = s2[:4]
	fmt.Println(s2)
	s2[3] = 4

}
