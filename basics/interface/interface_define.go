package main

import "fmt"

// AnimalInterface 实现了该接口中的所有方法，即视为实现了该接口
type AnimalInterface interface {
	Sing()
	Climb()
}

type Dog struct{}

type Cat struct{}

func (d Dog) Sing() {
	fmt.Println("wang~")
}

func (d Dog) Climb() {

}

func (c Cat) Sing() {
	fmt.Println("miao~")
}

func (c Cat) Climb() {

}

func main() {
	animalSlice := make([]AnimalInterface, 2)
	animalSlice[0] = &Dog{}
	animalSlice[1] = &Cat{}
	for _, animal := range animalSlice {
		animal.Sing()
	}
	/* 输出：
	wang~
	miao~
	*/
}
