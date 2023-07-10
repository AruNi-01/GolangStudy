package main

import "fmt"

func main() {
	sliceDefine()
	sliceCut()

	// 根据数组来初始化 slice
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[1:2])
	sli := arr[:]
	fmt.Println(sli) // [1 2 3 4 5]

}

func sliceDefine() {
	var sli1 []int // nil slice
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli1), cap(sli1), sli1)
	// len=0 cap=0 slice=[]

	var sli2 = []int{} // null slice
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli2), cap(sli2), sli2)
	// len=0 cap=0 slice=[]

	var sli3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli3), cap(sli3), sli3)
	// len=5 cap=5 slice=[1 2 3 4 5]

	sli4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli4), cap(sli4), sli4)
	// len=5 cap=5 slice=[1 2 3 4 5]

	// use make() to define slice
	var sli5 = make([]int, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli5), cap(sli5), sli5)
	// len=5 cap=5 slice=[0 0 0 0 0]

	sli6 := make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli6), cap(sli6), sli6)
	// len=5 cap=8 slice=[0 0 0 0 0]
}

func sliceCut() {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%d\n", len(sli), cap(sli), sli)
	/*
		通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
		注意：[] 区间 左闭右开
	*/

	// 切指定位置
	fmt.Println("sli[1] == ", sli[1])
	// 从头切到尾
	fmt.Println("sli[:] == ", sli[:])
	// 从 1(包含)切到最尾部
	fmt.Println("sli[1:] == ", sli[1:])
	// 从最开头切到 4(不包含)
	fmt.Println("sli[:4] == ", sli[:4])

	// 子切片从 0(包含)到 3(不包含) —> index 0,1,2
	fmt.Println("sli[0:3] == ", sli[0:3])
	fmt.Printf("len=%d cap=%d slic=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	// 子切片从 0 到 3(不包含)，4 控制切片的 cap (4-1)
	fmt.Println("sli[0:3:4] == ", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4])

}
