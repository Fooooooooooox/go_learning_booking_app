package foooox

import (
	"math"
)

// string代表我要返回的变量是string
func Bestfoooox() string {
	return "hhh"
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// what is argument before function?
// ==> method on types
// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.

type Vertex struct {
	X, Y float64
}

// 这个（v Vertex)表示的是一个type 之后可以使用v.Abs的方法来调用
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
