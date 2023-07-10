package main

import "fmt"

func main() {
	// 一维数组
	var arr1 = [5]int{1, 3, 4, 5, 6}
	fmt.Println(arr1) // [1 3 4 5 6]
	var arr2 = [5]int{0: 1, 2: 1}
	fmt.Println(arr2) // [1 0 1 0 0]

	// 二维数组
	var arr3 = [2][4]int{{0, 0}, {1, 1}}
	fmt.Println(arr3) // [[0 0 0 0] [1 1 0 0]]

	// 数组的长度和容量，由于数组大小不可变，所以 cap、len 永远相等
	fmt.Println("arr len: ", len(arr1))        // 5
	fmt.Println("arr cap: ", cap(arr1))        // 5
	fmt.Println("arr3 lenRow: ", len(arr3))    // 2
	fmt.Println("arr3 lenCol: ", len(arr3[0])) // 4

}
