package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

// Speak : Person 自己的方法
func (p Person) Speak() {
	fmt.Println("Person#Speak")
}

// Student 内嵌（相当于继承）了 Person（也就拥有了它的属性和方法）
type Student struct {
	Person
	grade string
	class string
}

func NewStudentWithPerson(person Person, grade string, class string) *Student {
	return &Student{Person: person, grade: grade, class: class}
}
func NewStudent(name string, age int, grade string, class string) *Student {
	p := NewPerson(name, age)
	return NewStudentWithPerson(*p, grade, class)
}

func main() {
	stu := NewStudent("abc", 18, "2020级", "软件工程二班")
	stu.Speak()
}
