package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	// 没有初始化值的字段会初始化成空值
	fmt.Println(a)
	fmt.Println(checkPassword(a, "haha"))
	fmt.Println(checkPassword(a, "1024"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPasswordPtr(u *user, password string) bool {
	// 对结构体使用指针能够实现对结构体的值的修改，同时能避免大结构体复制的开销
	return u.password == password
}
