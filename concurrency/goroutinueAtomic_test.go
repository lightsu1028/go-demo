package concurrency

import (
	"runtime"
	"sync"
	"testing"
)

var (
	wg      sync.WaitGroup
	counter int64
)

func TestGoRoutinue(t *testing.T) {

	wg.Add(2)

}

func addCounter(times int) {
	for i := 0; i < times; i++ {
		counter++
		//相当于yiel 当前goroutinue退出 返回
		runtime.Gosched()
	}
}
