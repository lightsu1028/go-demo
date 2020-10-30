package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

var mutext sync.Mutex
//使用互斥锁进行原子操作
func TestMutex(t *testing.T) {

	wg.Add(2)
	go incrCounter(5000)
	go incrCounter(5000)

	wg.Wait()
	fmt.Printf("%d",counter)
}

func incrCounter(times int) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		mutext.Lock()
		//doSomething
		counter++
		mutext.Unlock()
	}
}
