package json_learning

import (
	"encoding/json"
	"fmt"
)

// ``里的是json tag json tag主要在struct与json数据转换的过程(Marshal/Unmarshal)中使用。

// 参考https://blog.csdn.net/xz_studying/article/details/106012535

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
		sex:  "male",
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
