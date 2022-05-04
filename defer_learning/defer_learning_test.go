package defer_learning_test

import (
	"fmt"
	"testing"
)

// 源码中defer的定义
// type _defer struct {
// 	sp uintptr //函数栈指针
// 	pc uintptr //程序计数器
// 	fn *funcval //函数地址
// 	link *_defer //指向自身结构的指针， 用于链接多个defer
// }

// defer通常用来释放函数内部变量 通过defer，我们可以在代码中优雅的关闭/清理代码中所使用的变量
// 什么意思？

// 这个是最常规的用法
// defer 的作用就是注册一个方法的调用，该方法将在程序返回后进行输出。
func TestDe(t *testing.T) {
	defer t.Log("this defer fmt!!")
	t.Log("this is normal fmt!!")
}

//*********************分割线*****************
// 知识补充 值传递和引用传递（这个非常非常重要!
// https://blog.csdn.net/QcloudCommunity/article/details/86505416
// demo1
// test_string试图内部修改一个外部传入的字符串的值
// &是拿到一个变量的内存地址
// *是读取该地址上的变量
// *string 是string的指针地址
func test_string(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n", s, &s)
}

func TestM(t *testing.T) {
	s := "a"
	fmt.Printf("outer: %v, %v\n", s, &s)
	test_string(s)
	fmt.Printf("outer: %v, %v\n", s, &s)
}

func test_string2(s *string) {
	fmt.Printf("inner: %v, %v\n", *s, s)
	*s = "b"
	fmt.Printf("inner: %v, %v\n", *s, s)
}

// 那我改写成指针的形式就可以修改外部的值了
func TestM2(t *testing.T) {
	s := "a"
	fmt.Printf("outer: %v, %v\n", s, &s)
	test_string2(&s)
	fmt.Printf("outer: %v, %v\n", s, &s)
}

/*/
结果如下：
outer: a, 0xc00004e560
inner: a, 0xc00004e580
inner: b, 0xc00004e580
outer: a, 0xc00004e560
得出结论 ==》 golang中的函数参数是值传递 因为函数内部重新开了一个内存空间用来存放s 最终外部变量的值并未被改变
/*/
// 传递指针类型的参数才可以实现在函数内部直接修改内容，如果传递的是值本身的，会有一次拷贝发生（此时函数内外，该变量的地址会发生变化，通过第一个示例可以看出）因此，在函数内部的修改对原外部变量是无效的。

// 但是如果是map结构呢？
func test_map(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}

func TestMap(t *testing.T) {

	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	fmt.Printf("outer: %v, %p\n", m, m)
	test_map(m)
	fmt.Printf("outer: %v, %p\n", m, m)
}

/*/
运行结果：
=== RUN   TestMap
outer: map[a:1 b:2 c:3], 0xc00009a330
inner: map[a:1 b:2 c:3], 0xc00009a330
inner: map[a:11 b:2 c:3], 0xc00009a330
outer: map[a:11 b:2 c:3], 0xc00009a330
？和预料中不一样 外部的值被改变了
==》因为map在golang中是一个比较特别的结构 他不是单单是一些键值对而是一个hashtable 也就是他本身就是指针类型
/*/

// 回到defer
func TestDefer(t *testing.T) {
	a, b := 1, 2
	defer add(a, b)
	a++
	b++
}

func add(a, b int) {
	fmt.Println(a + b)
}

/*/
=== RUN   TestDefer
3
/*/
// 结果是3 ==》 因为defer使用的是值传递 把外部ab的值传给defer之后defercopy了这两个值放到新的内存空间里
