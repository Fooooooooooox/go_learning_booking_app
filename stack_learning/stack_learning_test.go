package stack_learning_test

import (
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/stack_learning"
	"github.com/holiman/uint256"
)

// go知识点：
// 在go中 用new()方法生成的都是指向一个值的指针
// new的用法示例：
// var p *int
// p = new(int)
// *p = 666
// fmt.Println("*p = ", *p)

func TestStack(t *testing.T) {
	var stack = stack_learning.Newstack()
	var data = stack.Data()
	var c *uint256.Int
	t.Log(data)
	c = new(uint256.Int).SetUint64(8)
	stack.Push(c)
	t.Log(stack.Data())
	c = new(uint256.Int).SetUint64(100000000)
	stack.Push(c)
	t.Log(stack.Data())
	c = new(uint256.Int).SetOne()
	stack.Push(c)
	t.Log(stack.Data())
	t.Log("here is thr peek():")
	t.Log(stack.Peek())
	var pop uint256.Int
	pop = stack.Pop()
	t.Log("here is the ret of pop:")
	t.Log(pop)
	t.Log("here is the final result:")
	t.Log(stack.Data())

	// 试一下add指令
	// 直接拿了add指令中的exec函数 改一下简化了

	x, y := stack.Pop(), stack.Peek()
	t.Log("after this assignment the stack becomes:")
	t.Log(stack.Data())
	y.Add(&x, y)

	t.Log("x:", x)
	t.Log("y:", y)
	t.Log(stack.Data())

}
