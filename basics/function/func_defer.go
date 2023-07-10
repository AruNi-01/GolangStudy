package main

import "fmt"

func main() {
	test01()
	test02()
}

func test01() {
	i := 0
	defer fmt.Println(i)
	i++
}

func test02() {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	i++
}
