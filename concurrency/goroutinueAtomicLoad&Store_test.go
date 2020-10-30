package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg2    sync.WaitGroup
	isDone int64
)

func TestAtomicLoadAndStore() {

	wg2.Add(2)

	doWork("A")
	doWork("B")

	wg2.Wait()
}

func doWork(name string) {

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&isDone) == 1 {
			fmt.Printf("%s had finished", name)
			break
		}
	}
}
