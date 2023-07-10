package main

import "fmt"

type Stu struct {
	name string
	age  int
}

func NewStu(name string, age int) *Stu {
	return &Stu{name: name, age: age}
}

// Student的方法：shangKe()
// (s Stu) s 相当于 this
func (s Stu) shangKe() {
	fmt.Println(s.name, "在上课")
}

// 指针类型的接受者，由于指针是指向内存地址的，所以可以改变其值
func (s *Stu) setName(newName string) {
	s.name = newName
}

func main() {
	stu := NewStu("AruNi", 21)
	stu.shangKe()
	stu.setName("Aaryn")
	stu.shangKe()

	/* 输出：
	AruNi 在上课
	Aaryn 在上课
	*/
}
