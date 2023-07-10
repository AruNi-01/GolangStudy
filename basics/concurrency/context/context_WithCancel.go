package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个可取消的 context 和 cancel 函数
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个 goroutine 并传入 context
	go task(ctx)

	// 等待一段时间后取消 context
	time.Sleep(3 * time.Second)
	cancel()

	// 等待一段时间后退出程序
	time.Sleep(1 * time.Second)
	fmt.Println("main is exited.")
}

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine is cancelled, it will return soon.")
			return
		default:
			fmt.Println("goroutine is running ~~~")
			time.Sleep(time.Second)
		}
	}
}
