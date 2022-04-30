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

In newer versions (post 1.13) of Go, you don't need to set environment variables like GOPATH, GOBIN, etc.

You also need to have a go.mod file at the project root. This will make the directory a Go module. This is also where the .git/ is located. This means that only one go.mod is needed per repository. Inside the project root you could do a go mod init remote-repo.com/username/repository

I installed Go using Homebrew on macOS so GOROOT is /opt/homebrew/Cellar/go/1.17.5/libexec. This location contains the standard library and runtimes for Go.

test and run commands are run in the format go COMMAND package_path/xxx. Without specifying the package_path ./ and just running go COMMAND xxx, the compiler assumes that the module xxx is located in GOROOT, and throws error package xxx is not in GOROOT (path/to/GOROOT/src/xxx) because it doesn't exist.

This behavior is expected because the package we are working with is not part of the Go SDK, i.e., not in GOROOT. The package we are working with will either end up in the go workspace or in the current working directory. Running go install compiles and puts an executable binary in $GOBIN (a.k.a $GOPATH/bin - here $GOPATH is the Go workspace). Running go build from inside a package compiles and puts an execuatble in that directory.

You haven't listed the files inside the server/ package and which file has the main function, so I'll emulate 3 workflows of a calculator each demonstrating more complexity. The last workflow is similar to your directory structure.

Directory Structure
Version 1:
Getting started with packages

Basic functionality

calculatorv1
├── go.mod                      <- go mod init github.com/yourname/calculatorv1
└── basic/
    ├── add.go
    ├── add_test.go
    ├── main.go
    ├── multiply.go
    └── multiply_test.go
Version 2:
More functionality

Multiple packages

calculatorv2
├── go.mod                      <- go mod init github.com/yourname/calculatorv2
├── main.go
└── basic/
│   ├── add.go
│   ├── add_test.go
│   ├── multiply.go
│   └── multiply_test.go
└─── advanced/
     ├── square.go
     └── square_test.go

In newer versions (post 1.13) of Go, you don't need to set environment variables like GOPATH, GOBIN, etc.

You also need to have a go.mod file at the project root. This will make the directory a Go module. This is also where the .git/ is located. This means that only one go.mod is needed per repository. Inside the project root you could do a go mod init remote-repo.com/username/repository

I installed Go using Homebrew on macOS so GOROOT is /opt/homebrew/Cellar/go/1.17.5/libexec. This location contains the standard library and runtimes for Go.

test and run commands are run in the format go COMMAND package_path/xxx. Without specifying the package_path ./ and just running go COMMAND xxx, the compiler assumes that the module xxx is located in GOROOT, and throws error package xxx is not in GOROOT (path/to/GOROOT/src/xxx) because it doesn't exist.

This behavior is expected because the package we are working with is not part of the Go SDK, i.e., not in GOROOT. The package we are working with will either end up in the go workspace or in the current working directory. Running go install compiles and puts an executable binary in $GOBIN (a.k.a $GOPATH/bin - here $GOPATH is the Go workspace). Running go build from inside a package compiles and puts an execuatble in that directory.

You haven't listed the files inside the server/ package and which file has the main function, so I'll emulate 3 workflows of a calculator each demonstrating more complexity. The last workflow is similar to your directory structure.

Directory Structure
Version 1:
Getting started with packages

Basic functionality

calculatorv1
├── go.mod                      <- go mod init github.com/yourname/calculatorv1
└── basic/
    ├── add.go
    ├── add_test.go
    ├── main.go
    ├── multiply.go
    └── multiply_test.go
Version 2:
More functionality

Multiple packages

calculatorv2
├── go.mod                      <- go mod init github.com/yourname/calculatorv2
├── main.go
└── basic/
│   ├── add.go
│   ├── add_test.go
│   ├── multiply.go
│   └── multiply_test.go
└─── advanced/
     ├── square.go
     └── square_test.go
Version 3:
Even more functionality

Nested packages

calculatorv3
├── go.mod                      <- go mod init github.com/yourname/calculatorv3
├── main.go
└── basic/
│   ├── add.go
│   ├── add_test.go
│   ├── multiply.go
│   └── multiply_test.go
└─── advanced/
     ├── square.go
     ├── square_test.go
     └── scientific/
         ├── declog.go
         └── declog_test.go
