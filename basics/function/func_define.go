package main

func main() {

}

func func1() int {
	return 1
}

func func2(a, b int) (int, error) {
	return 1, nil
}

func func3(a ...int) {
	for _, val := range a {
		if val == 1 {

		}
	}
}

type binOp func(a, b int) int
