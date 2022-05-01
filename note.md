## how to star?

go mod init github.com/Fooooooooooox/foooox
go mod init github.com/Fooooooooooox/interface_learning


go run main.go


## loop

there are while, do while, for each in other languages

but in go 

you only have "for" for loop

1. infinite loop

for{
    
}

2. finite loop

 
// range is for making iterable stuff

for index, booking := range bookings {

}

//but we don't need index actually, so we can write in this way:

for _, booking := range bookings {

}




## unused variables

in go, if you define a variable but don't use it it throws an error

if you want to define an unused variable: use _


## export function

if you want your function to be used in other file, you can capitalize the first letter of the function name, then it'll be automatically exported.

and you can do the same with variables

you capitalize the first letter and it'll be automatically exported


## can i have multiple main.go?

go和其他的编程语言不同 只可以有一个main作为入口文件

这会让测试变得很不方便 

解决方法：编写可测试的go代码（参考：https://blog.csdn.net/demored/article/details/123983265，https://codeplayer.vip/p/j7tek

文件名必须以xxx_test.go命名；

方法名称必须是Test[^a-z]开头，而且 Test 的后缀部分第一个字符必须大写；

方法参数必须是 t *testing.T。

## about package

package is folder.

package name is folder name.

package path is folder path.

## how to recognize your packages?
https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project


## how to print stuff when testing?
-v
go test -v /Users/foooox/go_learning_booking_app/go_learning_booking_app/interface_learning/interface_learning_test.go

## export fun 

如果你要使用你在package里自定义的fun 要把首字母大写（自动export fun）
