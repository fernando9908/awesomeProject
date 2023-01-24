package main

import "fmt"

func main() {
	res := add(2, 4)
	fmt.Println(res)
	v, ok := exist(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok)
}

func add(a int, b int) int {
	return a + b
}

func exist(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}
