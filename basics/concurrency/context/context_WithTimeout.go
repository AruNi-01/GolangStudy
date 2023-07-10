package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建一个 50ms 超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel() // 通知子 goroutine 结束
	wg.Add(1)

	var ok bool
	var connection any
	go func(ctx context.Context) {
		connection, ok = connectToDB(ctx)
		wg.Done()
	}(ctx)

	// 启动另一个 goroutine 去等待数据库连接完成，否则 main goroutine 会阻塞在 wg.Wait()
	go func() {
		wg.Wait()
		if ok {
			fmt.Println("connect success, connection is:", connection)
		} else {
			fmt.Println("connect fail, connection is:", connection)
		}
	}()

	time.Sleep(time.Second) // 防止 main goroutine 结束太快
}

func connectToDB(ctx context.Context) (connection any, ok bool) {
	fmt.Println("db connecting...")
	time.Sleep(time.Millisecond * 20) // 模拟数据库连接
	select {
	case <-ctx.Done():
		fmt.Println("goroutine is cancel due to", ctx.Err())
		connection = nil
		ok = false
	default:
		fmt.Println("db connected!")
		connection = "my-db-connection"
		ok = true
	}
	return
}
