package main

import (
	"fmt"
	"time"
)

/* goroutine 池：控制 goroutine 的数量，去执行多个任务 */

// worker：从 jobs 中获取任务，执行完毕后将结果写入 results
func worker(id int, jobs <-chan int, results chan<- int) {
	// 循环 jobs 获取任务
	for job := range jobs {
		fmt.Printf("worker(goroutine)-%d start job: %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500) // 模拟执行了 5ms
		fmt.Printf("worker(goroutine)-%d stop job: %d\n", id, job)
	}
	// results 通道接收完后关闭，避免下面从中取结果输出时一直等待，造成死锁
	close(results)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启 3 个 goroutine 去执行下面的 5 个任务，一个 goroutine 对应一个 worker
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 向 jobs 通道中发送 5 个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	// 发送完任务后记得关闭，避免上面从 jobs 中获取任务时一直等待，发生死锁
	close(jobs)

	// 从 results 通道中输出结果
	for result := range results {
		fmt.Println(result)
	}
}
