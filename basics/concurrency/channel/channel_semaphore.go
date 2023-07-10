package main

import (
	"fmt"
	"sort"
	"time"
)

func doTask(ch chan int) {
	ch <- task() // task() 任务完成后，向 chan 中发送信号
}

func task() int {
	fmt.Println("doing task")
	time.Sleep(time.Second * 3)
	fmt.Println("task done")
	return 1
}

func main() {
	ch := make(chan int)
	go doTask(ch)
	// ......
	fmt.Println("main wait ~~")
	<-ch                           // 阻塞读取 chan
	fmt.Println("after operation") // 后续操作...

	/*
		main wait ~~
		doing task
		task done
		after operation
	*/
}

// 匿名函数版：无需返回值入 chan
func anonymousFunc() {
	ch := make(chan int)
	go func() {
		task()
		ch <- 1
	}()

	<-ch
}

// 等待两个协程完成任务
func waitTowGoroutine() {
	done := make(chan bool)
	// 定义一个做排序的函数，调用完 sort() 后将一个数据塞入 chan
	doSort := func(s []int) {
		sort.Ints(s)
		done <- true
	}
	sli := []int{2, 1, 8, 7, 9}
	i := 2
	go doSort(sli[:i])
	go doSort(sli[i:])
	fmt.Println("maybe not sorted absolutely: ", sli)
	<-done
	<-done
	fmt.Println("sorted absolutely: ", sli)

	/*
		在信号量到达之前输出 sli，可能会出现还没全部排好序的情况
		maybe not sorted absolutely:  [1 2 8 7 9]
		sorted absolutely:  [1 2 7 8 9]
	*/
}

// 等待 N 个协程完成任务
const N = 20

type Empty interface{}

var empty Empty

func doSomething(i int, xi float64) float64 { return 0 }

func waitNGoroutine() {
	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N)

	// range data, let goroutine run the task
	for i, xi := range data {
		/*
			注意：这里需要拷贝一份 i，xi 传入匿名函数，再让 goroutine 执行，否则就会形成闭包，
			这样外部 i/xi 改变时，会影响到 goroutine 的值，所以需要和外部变量隔离。
			而 res 变量不用拷贝一份给 goroutine，因为它是共享的，保证 i 不受影响即可。
		*/
		go func(i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		}(i, xi)
	}

	// wait for goroutine to finish
	for i := 0; i < N; i++ {
		<-sem
	}
}

// 带缓冲的 channel 实现 PV，然后实现互斥
// type Empty interface {}
type semaphore chan Empty

/* 对信号量进行操作 */
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// 实现一个互斥
/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
