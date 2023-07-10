package main

import "fmt"

// 只进行发送的 chan
func counter(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch) // 发送完毕 -> close
}

// 从 ch1 中取数据，发送到 ch2 中。即只进行接收的 ch1，只进行发送的 ch2
func squarer(ch1 <-chan int, ch2 chan<- int) {
	for v := range ch1 {
		ch2 <- v * v
	}
	close(ch2)
}

// 打印 chan 中的数据，即只进行接收的 chan
func printer(ch <-chan int) {
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	counter(ch1)
	squarer(ch1, ch2)
	printer(ch2)
}
