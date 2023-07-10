package main

import (
	"fmt"
	"time"
)

func simpleUse() {
	ch := make(chan int, 5)
	// channel 只能使用 <- 进行操作
	ch <- 10
	ch <- 20
	var x1 int
	// 从 ch 中读取一个数据，与 10 进行对比
	if <-ch != 10 {
		x1 = <-ch
	}
	x2 := <-ch
	close(ch)
	fmt.Printf("x1: %d, x2: %d\n", x1, x2)
	// x1: 0, x2: 20
}

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch chan int) {
	var v int
	//var ok bool
	//for {
	//	if v, ok = <-ch; !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}

	for {
		v = <-ch
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 10)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}
