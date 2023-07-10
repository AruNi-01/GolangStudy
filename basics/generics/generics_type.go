package main

type Slice[T int | string] []T

type Map[K int | string, V float32 | float64] map[K]V

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

// Lookup 泛型类型的方法
func (t *Tree[T]) Lookup(x T) *Tree[T] {
	t1 := new(Tree[T])
	return t1
}

func main() {
	// 泛型类型必须实例化后才能使用
	var stringTree Tree[string]
	stringTree.left = nil
	stringTree.right = nil
	stringTree.value = "val"
}
