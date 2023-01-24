package main

import "fmt"

// 指针作用有限，主要针对函数传入参数进行修改
func main() {
	n := 5
	addptr(&n)
	add(n)
	fmt.Println(n)
	//addptr(&n)
	fmt.Println(n)
}

func add(n int) {
	n += 2
}

func addptr(n *int) {
	*n += 2
}
