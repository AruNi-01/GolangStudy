package main

import "fmt"

type struct1 struct {
	i1  int
	f2  float32
	str string
}

func main() {
	s1 := new(struct1) // 返回一个指针
	//s1 := &struct1{}	// 等价于上面的 new

	(*s1).i1 = 1
	s1.f2 = 1.0 // Golang 语法糖：可以直接用 s1 指针 `.`
	s1.str = "str"
	fmt.Println(s1)
}
