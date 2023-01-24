package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	r1, ok1 := m["unknown"]
	fmt.Println(r1, ok1)
	r2, ok2 := m["two"]
	fmt.Println(r2, ok2)
}
