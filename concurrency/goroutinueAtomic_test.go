package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	counter int64
	wg1     sync.WaitGroup
)

//测试atmoic包使用
func TestGoRoutinue(t *testing.T) {

	wg1.Add(2)

	go addCounter(5000)
	go addCounter(5000)

	wg1.Wait()
	fmt.Printf("counter:%d", counter)

}

func addCounter(times int) {
	defer wg1.Done()
	for i := 0; i < times; i++ {
		//原子
		atomic.AddInt64(&counter,1)

		//相当于yiel 当前goroutinue退出 放回队列？
		//runtime.Gosched()
	}
}
