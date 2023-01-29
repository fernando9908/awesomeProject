package concurrence

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello goroutine : " + fmt.Sprint(i))
}

func HelloGoRountine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	// 休息1s
	time.Sleep(time.Second)
}
