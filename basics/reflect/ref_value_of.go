package main

import (
	"fmt"
	"reflect"
)

func refValue(x interface{}) {
	valueX := reflect.ValueOf(x)
	kindX := valueX.Kind()
	switch kindX {
	case reflect.Int32:
		fmt.Printf("type is int32, value is: %d\n", int32(valueX.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is: %f\n", float32(valueX.Float()))
	}
}

func main() {
	var a int32 = 1
	var b float32 = 1
	refValue(a) // type is int32, value is: 1
	refValue(b) // type is float32, value is: 1.000000
}
