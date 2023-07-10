package main

func main() {

}

// 类型约束字面量，通常外层 interface{} 可省略
func minWithConstraint[T interface{ int | float64 }](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Value 事先定义好的类型约束类型
type Value interface {
	int | float64
}

func minWithConstraint2[T Value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// 在使用类型约束时，如果省略了外层的 interface{} 会引起歧义，那么就不能省略。例如：

// type IntPtrSlice1 [T * int][]T             // 编译报错：T*int ? 数组的长度必须是一个非负整数常量
type IntPtrSlice2[T *int,] []T             // 只有一个类型约束时可以添加 `,`
type IntPtrSlice3[T interface{ *int }] []T // 使用 interface{} 包裹
