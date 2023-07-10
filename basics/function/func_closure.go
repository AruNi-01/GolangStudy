package main

import "fmt"

func main() {
	var f = adder()
	fmt.Println(f(10)) // 20
	fmt.Println(f(20)) // 40
	fmt.Println(f(30)) // 70

	f1 := adder()
	fmt.Println(f1(40)) // 40
	fmt.Println(f1(50)) // 90
}

/*
闭包：调用 adder() 会返回一个函数，这个函数使用了外部作用域中的 x 变量，
那么变量 x 将会在这个返回函数的生命周期中一直存在。
*/
func adder() func(int) int {
	var x int
	// 返回一个函数
	return func(y int) int {
		x += y
		return x
	}
}

// 下面是使用函数参数的写法
func adder2(x int) func(int) int {
	return func(a int) int {
		x += a
		return x
	}
}

// 闭包的进阶
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func calcTest() {
	// f1, f2共享一个base，因为f1，f2在同一个闭包内
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) // 11(10+1) 9(11-2)
	fmt.Println(f1(3), f2(4)) // 12(9+3) 8(12-4)
	fmt.Println(f1(5), f2(6)) // 13(8+5) 7(13-6)
}
