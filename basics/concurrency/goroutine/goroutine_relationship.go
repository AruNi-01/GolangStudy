package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 协程开始...")
	go func() {
		fmt.Println("父协程开始...")
		go func() {
			fmt.Println("子协程开始...")
			for {
				fmt.Println("子协程 doing...")
			}
		}()
		time.Sleep(time.Microsecond * 20)
		fmt.Println("父协程退出...")
	}()
	time.Sleep(time.Microsecond * 80)
	fmt.Println("main 协程退出...")
}
