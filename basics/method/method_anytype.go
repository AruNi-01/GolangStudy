package main

import "fmt"

type MySlice []int

func (ms MySlice) MySliceSum() int {
	sum := 0
	for _, val := range ms {
		sum += val
	}
	return sum
}

func main() {
	ms := MySlice{1, 2, 3}
	fmt.Println(ms.MySliceSum()) // 6
}
