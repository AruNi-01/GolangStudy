package main

import (
	"fmt"
	"reflect"
)

func refType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", t)
}

// kind 是底层的类型
func type_name_and_kind(x interface{}) {
	typeX := reflect.TypeOf(x)
	fmt.Printf("type(typeName): %v, typeKind: %v\n", typeX.Name(), typeX.Kind())
}

type MyInt int

func main() {
	var a float32 = 1.0
	refType(a) // float32
	var b int32 = 1
	refType(b) // int32

	var ptr *int32
	var r rune // int32 别名
	var myInt MyInt
	type str struct{}
	// 注：数组、slice、指针、map 等，typeName 都为空
	type_name_and_kind(ptr)   // type(typeName): , typeKind: ptr
	type_name_and_kind(r)     // type(typeName): int32, typeKind: int32
	type_name_and_kind(myInt) // type(typeName): MyInt, typeKind: int
	type_name_and_kind(str{}) // type(typeName): str, typeKind: struct

}
