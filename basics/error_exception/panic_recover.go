package main

import "fmt"

func badCall() {
	// 引发一个 panic，如果不进行捕获则会使程序崩溃
	panic("this is a badCall, will be panic")
}

func test() {
	// 使用 recover 捕获 panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panicing: %s\n", err)
		}
	}()
	badCall()
	fmt.Println("badCall after")
}

func main() {
	fmt.Println("test() call")
	test()
	fmt.Println("test() completed")

	/* 输出：
	test() call
	panicing: this is a badCall, will be panic
	test() completed
	*/
}
