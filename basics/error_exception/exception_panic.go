package main

import (
	"fmt"
	"log"
)

func usePanic() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

// protect 函数调用函数参数 g 来保护调用者，防止其从 g 中抛出的运行时 panic
func protect(g func()) {
	// 用 defer 做最后的 panic 捕获
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() // 可能会发生 runtime error
}

func main() {

}
