package json_learning

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// struct有点像python里的类型 class
// struct里的元素如果加上了json tag可以在将struct转换为json的时候作为json的键

// ``里的是json tag json tag主要在struct与json数据转换的过程(Marshal/Unmarshal)中使用。

// 参考https://blog.csdn.net/xz_studying/article/details/106012535

type Snap struct {
	Number  uint64                      `json:"number"`
	Hash    common.Hash                 `json:"hash"`
	singers map[common.Address]struct{} `json:"signers"`
}

// 不过这里的空struct是用来做啥的呢？

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgn  bool
	sex   string
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func Students() string {
	stu := Stu{
		Name: "ali",
		Age:  20,
		HIgn: true,
		// sex 首字母没有大写 在生成json的时候是无法导出的
		sex: "male",
	}

	cla := new(Class)
	cla.Name = "first class"
	cla.Grade = 3
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
		return "something went wrong"
	}
	return string(jsonStu)

}

// 在json中使用interface 会让json变得特别灵活

type Stu_i struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
}

type Class_i struct {
	Name  string
	Grade int
}

func Students_i() string {
	stu_i := Stu_i{
		Name: "ali",
		Age:  20,
		HIgh: true,
		sex:  "male",
	}

	cla_i := new(Class_i)
	cla_i.Name = "first class"
	cla_i.Grade = 3
	stu_i.Class = cla_i

	jsonStu_i, err := json.Marshal(stu_i)
	if err != nil {
		fmt.Println(err)
		return "something went wrong"
	}
	return string(jsonStu_i)

}
