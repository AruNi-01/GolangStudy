package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "requestID", "123456")
	// 在另一个 goroutine 函数中使用 Context 对象
	go doSomething(ctx)

	time.Sleep(time.Millisecond)
}

func doSomething(ctx context.Context) {
	// 从 Context 对象中获取请求 ID
	requestID := ctx.Value("requestID").(string)
	fmt.Println("Request ID:", requestID)
}

// 输出：Request ID: 123456
