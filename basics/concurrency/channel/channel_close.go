package main

import "fmt"

// 判断 chan 是否关闭方式一：使用多值返回模式
func channelIsCloseByMultiRet(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Printf("v: %#v, ok: %#v\n", v, ok)
	}
	/*
		v: 1, ok: true
		v: 2, ok: true
		channel closed
	*/
}

// 判断 chan 是否关闭方式二：使用 for-range 来遍历 chan
func channelIsCloseByForRange(ch chan int) {
	// 循环将一直阻塞，直到 channel 被关闭
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("channel closed")
	/*
		1
		2
		channel closed
	*/
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	//channelIsCloseByMultiRet(ch)
	channelIsCloseByForRange(ch)
}
