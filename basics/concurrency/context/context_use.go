package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个过期时间为 1s 的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 向 ctx 中传入 handle 函数
	go handle(ctx, 500*time.Millisecond)
	// select 轮询 ctx.Done() channel
	select {
	case <-ctx.Done():
		fmt.Println("func-main ctx.Err:", ctx.Err())
	}
}

// handle 函数使用 500ms 处理传入的请求
func handle(ctx context.Context, duration time.Duration) {
	// select 轮询 ctx.Done() channel 和 time.After() channel
	select {
	case <-ctx.Done():
		fmt.Println("func-handle ctx.Err:", ctx.Err())
	// 模拟在 duration 时间已完成请求
	case <-time.After(duration):
		fmt.Println("finished: process request with", duration)
	}
}
