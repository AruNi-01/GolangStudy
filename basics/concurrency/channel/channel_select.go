package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(time.Millisecond)
}

/* 生产者 */
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 1
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 100
	}
}

/* 消费者 */
func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("received on channel 2: %d\n", v)
		}
	}
}
