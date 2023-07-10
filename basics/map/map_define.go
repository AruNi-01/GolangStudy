package main

import "fmt"

func main() {
	test01()
}

func test01() {
	var map1 = map[string]int{"k1": 1, "k2": 2}
	fmt.Println(map1["k1"])

	map2 := make(map[string]int)
	map2["k1"] = 1
	map2["k2"] = 2

	// function 也可作为 value
	funcMap := map[int]func(i int) int{
		1: func(i int) int {
			return i + 1
		},
		2: func(i int) int {
			return i + 2
		},
		3: func(i int) int {
			return i + 3
		},
	}
	fmt.Println(funcMap)
}
