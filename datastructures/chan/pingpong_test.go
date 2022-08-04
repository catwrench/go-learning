package _chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Ball struct {
	Hits int64 // 击打次数
}

func TestPingPong(t *testing.T) {
	table := make(chan *Ball)
	tableClose := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		play("A", table, tableClose)
		wg.Done()
	}()
	go func() {
		play("B", table, tableClose)
		wg.Done()
	}()
	// 发球
	table <- new(Ball)
	// 等1秒，取走球
	time.Sleep(time.Second)
	<-table

	tableClose <- 1
	tableClose <- 1
	close(table)
	wg.Wait()
	fmt.Println("结束")
}

func play(name string, table chan *Ball, tableClose chan int) {
	for {
		select {
		case <-tableClose:
			return
		case <-time.After(time.Second * 3):
			return
		case ball := <-table:
			// 接球
			ball.Hits++
			fmt.Println(name, "hits count:", ball.Hits)
			// 打球
			time.Sleep(100 * time.Millisecond)
			table <- ball
		}
	}
}
