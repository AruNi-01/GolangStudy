package main

import "fmt"

func main() {
	sliInt := []int32{1, 2, 3, 4, 5}
	sliFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sliString := []string{"1", "2", "3", "4", "5"}
	reverseInt := reverseWithGenerics(sliInt)
	reverseFloat := reverseWithGenerics(sliFloat)
	reverseString := reverseWithGenerics(sliString)
	fmt.Printf("reverseInt: %v\nreverseFloat: %v\nreverseString: %v\n", reverseInt, reverseFloat, reverseString)

	/*
		输出：
		reverseInt: [5 4 3 2 1]
		reverseFloat: [5.5 4.4 3.3 2.2 1.1]
		reverseString: [5 4 3 2 1]
	*/

	// 类型参数列表的使用：
	// 类型实例化
	min[int](1, 3)
	min[float64](1.1, 3.3)
	// 类型实例化时，将返回一个实参的类型函数
	IntMinFunc := min[int]
	IntMinFunc(1, 2)
	FloatMinFunc := min[float64]
	FloatMinFunc(1.1, 3.3)

}

func reverseWithGenerics[T any](slice []T) []T {
	l := len(slice)
	newSlice := make([]T, l)
	for i, e := range slice {
		newSlice[l-i-1] = e
	}
	return newSlice
}

// 类型参数列表
func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
