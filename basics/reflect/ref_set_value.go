package main

import (
	"fmt"
	"reflect"
)

func ref_set_value(x interface{}) {
	valueX := reflect.ValueOf(x)
	// 在反射中使用 Elem() 获取指针对应的值
	if valueX.Elem().Kind() == reflect.Int64 {
		valueX.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	ref_set_value(&a)     // 只能传递地址，才能修改成功值
	fmt.Println("a: ", a) // a:  200
}
