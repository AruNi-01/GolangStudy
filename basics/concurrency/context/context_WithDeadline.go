package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有截止时间(当前时间 + 3s)的 context 和 cancel 函数
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))

	// 启动一个 goroutine 并传入 context
	go task2(ctx)

	// 等待一段时间(5s)后取消 context
	time.Sleep(5 * time.Second)
	cancel()

	// 等待一段时间后退出程序
	time.Sleep(time.Second)
	fmt.Println("main is exited.")
}

func task2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine is canceled due to", ctx.Err())
			return
		default:
			fmt.Println("goroutine is running ~~~")
			time.Sleep(time.Second)
		}
	}
}
