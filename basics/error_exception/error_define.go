package main

import "errors"

// PathError 自定义错误
type PathError struct {
	Op   string
	Path string
	Err  error
}

// String 实现该方法， fmt 打印时使用此格式输出
func (pe PathError) String() string {
	return pe.Op + " " + pe.Path + ": " + pe.Err.Error()
}

// Error 实现该方法，使用 Error() 打印时使用此格式输出
func (pe PathError) Error() string {
	return pe.Op + " " + pe.Path + ": " + pe.Err.Error()
}

func main() {
	pathError := PathError{"myOp", "myPath", errors.New("error")}
	pathError.Err.Error()
}
