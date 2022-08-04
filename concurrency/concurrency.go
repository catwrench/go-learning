package concurrency

import (
	"fmt"
	"sync"
)

// SyncWait 并发
type SyncWait struct {
}

func (s *SyncWait) Test() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // 给未匹配的元素预留空间

	wg := &sync.WaitGroup{}
	for _, name := range people {
		wg.Add(1)
		go s.Seek(name, match, wg)
	}
	wg.Wait()

	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		fmt.Println("== 没有待处理的发送操作 ==")
	}
}

// Seek 寻求发送或接收匹配上名称名称的通道,并在完成后通知 等待组.
func (s *SyncWait) Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("从 %s 处收到一条消息： %s.\n", peer, name)
	case match <- name:
		fmt.Println("== 等待其他人接受消息 ==")
	}
	wg.Done()
}
