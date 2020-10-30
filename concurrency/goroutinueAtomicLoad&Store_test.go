package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	wg2    sync.WaitGroup
	isDone int64
)

/*
*  使用atomic的Store和Load进行并发标记设置和读取
 */
func TestAtomicLoadAndStore(t *testing.T) {

	wg2.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	atomic.StoreInt64(&isDone, 1)
	wg2.Wait()

}

func doWork(name string) {

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&isDone) == 1 {
			fmt.Printf("%s had finished\n", name)
			break
		}
	}
}
