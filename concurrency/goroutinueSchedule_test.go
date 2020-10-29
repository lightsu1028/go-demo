package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup

//测试调度器为防止某个goroutinue长时间占用处理器，进行切换goroutinue
func TestGoroutinueSchedule(t *testing.T) {
	//逻辑处理器为1
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("start!!!")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("wait to finish")
	wg.Wait()
	fmt.Println("progam terminated")


}

func printPrime(prefix string) {
	defer wg.Done()

	next:
	for i := 2; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%2 == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, i)
	}
}
