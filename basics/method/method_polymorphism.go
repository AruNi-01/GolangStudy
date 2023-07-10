package main

import "fmt"

type Animal interface {
	Say()
}

type Dog struct{}

type Cat struct{}

func (d Dog) Say() {
	fmt.Println("wang~")
}

func (c Cat) Say() {
	fmt.Println("miao~")
}

func main() {
	animalSlice := make([]Animal, 2)
	animalSlice[0] = &Dog{}
	animalSlice[1] = &Cat{}
	for _, animal := range animalSlice {
		animal.Say()
	}
	/* 输出：
	wang~
	miao~
	*/
}
