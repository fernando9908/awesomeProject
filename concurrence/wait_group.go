package concurrence

import (
	"sync"
)

func HelloGoWait() {
	var wg sync.WaitGroup
	// 开启五个协程
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
