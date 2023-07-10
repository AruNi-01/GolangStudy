package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}

// Exam Student 的方法，注意方法名大写，否则在反射中获取不到
func (s Student) Exam() {
	fmt.Println(s.Name, "do exam...")
}

func main() {
	stu := Student{
		Name:  "AruNi",
		Score: 90.5,
	}
	printFiled(stu)
	printMethod(stu)

	// 如果是使用 new 的方式实例化 Student，则返回的是指针，所以要么传递值进去，要么在反射中通过 Elem() 获取对应的值
	stuPtr := &Student{} // & 的方式等同于 new()
	printFiled(*stuPtr)
	printMethod(*stuPtr)
}

func printFiled(x interface{}) {
	t := reflect.TypeOf(x)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("fieldType: %s", field.Type)
		fmt.Printf("，fieldName: %s", field.Name)
		fmt.Printf("，fieldTag: %s\n", field.Tag.Get("json"))
	}

	// 通过字段名获取字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("fieldType: %s", scoreField.Type)
		fmt.Printf("，fieldName: %s", scoreField.Name)
		fmt.Printf("，fieldIndex: %d", scoreField.Index)
		fmt.Printf("，fieldTag: %s\n", scoreField.Tag.Get("json"))
	}

}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("methodType: %s", v.Method(i).Type())
		fmt.Printf("，methodName: %s\n", t.Method(i).Name)

		// 调用方法，通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args) // 无参数也可以直接传 nil 进去
	}
}
