package concurrence

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			// 将生产的数字i发送到src通道channel
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	// 能够保证数字的输出顺序。即并发安全
	for i := range dest {
		println(i)
	}
}
