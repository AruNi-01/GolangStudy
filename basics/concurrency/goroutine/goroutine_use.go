package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello~")
}

func main() { // 开启一个主 goroutine 去执行main函数
	go hello() // 开启另一个 goroutine 去执行hello函数
	fmt.Println("main~")
	// 睡 1s，别让 main goroutine 结束，否则该 goroutine 下的所有 goroutine 都会销毁
	time.Sleep(time.Second) // 更优雅的方式是使用 WaitGroup（等待-唤醒）

	// 一般 go 是配合匿名函数来执行的
	go func() {
		fmt.Println("hello~")
	}()
}
