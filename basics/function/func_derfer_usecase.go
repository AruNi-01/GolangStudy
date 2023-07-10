package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer EnterExitFunc()()
	fmt.Println("main running...")

	/*
		输出：
		main.main enter, line - 10
		main running...
		main.main exit; time-consuming: 38.752µs
	*/
}

func EnterExitFunc() func() {
	pc, _, lineNo, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	startTime := time.Now()
	fmt.Println(funcName, "enter, line -", lineNo)
	return func() {
		fmt.Println(funcName, "exit;", "time-consuming:", time.Since(startTime))
	}
}
