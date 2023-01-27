package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 第2个参数代表进制，第3个参数代表精度
	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n)

	// 传0代表自动推测
	n, _ = strconv.ParseInt("0X1000", 0, 64)
	fmt.Println(n)

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)

	n2, err := strconv.Atoi("AAa")
	fmt.Println(n2, err)
}
