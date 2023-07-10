package main

import "fmt"

type Any interface{}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
	// ...
}

// 格式化 Car 结构体的输出
func (c Car) String() string {
	return fmt.Sprintf("{%s, %s, %d}", c.Model, c.Manufacturer, c.BuildYear)
}

// Cars 切片类型来存储 Car
type Cars []*Car

// Process 根据传入进来的函数处理 car
func (cs Cars) Process(f func(car *Car)) {
	// 获取所有的 car，执行传进来的函数
	for _, car := range cs {
		f(car)
	}
}

// FindAll 寻找出符合条件的 Car，条件函数为 f（需传入）
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	/* 调用 Process，传去一个一个函数。这里的整体调用逻辑为：
		1. 外部传入一个条件函数；
		2. 调用 Process，传入的函数为下面的实现，即执行条件函数后讲符合条件的 car 放入结果；
		3. 在 Process 中遍历 Cars，执行下面的实现，然后调用条件函数，判断每一个 car

	   所以调用循序是：FindAll -> Process -> Process 的参数（函数）-> FindAll 的参数（函数）

	   外部环境 cars 和下面的函数形成一个闭包，所以在执行 Process 函数时，可以对 cars 进行 append
	*/
	cs.Process(func(c *Car) {
		// 调用传入的函数，符合则加入到 cars 中
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})

	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	/* 输出：
	AllCars:  [{Fiesta, Ford, 2008} {XL 450, BMW, 2011} {D600, Mercedes, 2009} {X 800, BMW, 2008}]
	New BMWs:  [{XL 450, BMW, 2011}]
	*/
}
