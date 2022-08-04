package concurrency

import (
	"fmt"
	"sync"
)

// WaitForGroup 等待同步
func WaitForGroup() {
	// 创建同步对象，添加5个等待
	var wg = sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		// 设置闭包用的变量
		n := i
		go func() {
			fmt.Println(n)
			// 告诉同步，已经执行完一个goroutine
			wg.Done()
		}()
	}

	// 等待全部goroutine执行完成
	wg.Wait()
	fmt.Println("执行结束")
}
