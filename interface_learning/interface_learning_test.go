package interface_learning_test

import (
	"fmt"
	"hash"
	"math"
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/interface_learning"
)

func TestInterface(t *testing.T) {
	t.Log("hhhhh")
	TestAreas := interface_learning.GetAreas()
	t.Log(TestAreas)
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func GetArea(s Shape) float64 {
	return s.area()
}

func GetAreas() []float64 {
	c1 := Circle{4.5}
	r1 := Rect{1, 2}
	Shapes := []Shape{c1, r1}
	var areas []float64
	for _, shape := range Shapes {
		areas = append(areas, GetArea(shape))
	}
	return areas
}

// 如何查看一个interface是否实现了某个method：
// 这里要注意的是 如果你检查的是golang它自带的一些函数是不会报错的 比如int64
// 但是如果你要检查一个自定义的接口方法是否存在 直接输入method的名字是不可行的 会抛出undefined的错误 你需要先定义这个method
func f(v interface{}) {
	type readerHash interface {
		hash.Hash
		Read([]byte) (int, error)
	}
	// 这个写法应该是ok为真的时候执行 否则不执行
	if _, ok := v.(readerHash); ok {
		fmt.Println("area")
	}
	// 重写一个
	_, ok := v.(readerHash)
	fmt.Println("ok = ", ok)

	// 所以如果你要判断有没有实现某个interface method：这样：
	if _, ok := v.(Shape); ok {
		fmt.Println("Shape")
	}
	// 重写一个
	_, ok2 := v.(Shape)
	fmt.Println("ok2 = ", ok2)
}

func TestOk(t *testing.T) {
	c1 := Circle{4.5}
	r1 := Rect{1, 2}
	Shapes := []Shape{c1, r1}
	for _, shape := range Shapes {
		f(shape)
	}
}
