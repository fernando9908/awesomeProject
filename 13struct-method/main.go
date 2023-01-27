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
	fmt.Println(a.checkPassword("haha"))
	a.resetPassword("haha")
	a.resetPasswordPtr("haha")
	fmt.Println(a.checkPassword("haha"))
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u user) resetPassword(password string) {
	// 不带指针不能实现对结构体的修改
	u.password = password
}

//func (u *user) checkPasswordPtr(password string) bool {
//	// 对结构体使用指针能够实现对结构体的值的修改，同时能避免大结构体复制的开销
//	return u.password == password
//}

func (u *user) resetPasswordPtr(password string) {
	// 对结构体使用指针能够实现对结构体的值的修改，同时能避免大结构体复制的开销
	u.password = password
}
