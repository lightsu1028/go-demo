package concurrency

import (
	"fmt"
	"log"
	"runtime/debug"
	"testing"
)

func TestSelect(t *testing.T) {

	ints := make(chan int, 10)
	ints <- 100
	ints <- 40
	ints <- 23
	for {
		select {
		case v := <-ints:
			go RecoverWrap(func(){
				fmt.Println(v)
				return
			})
		default:
			fmt.Println("select end")
		}
	}
}

func RecoverWrap(f func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("SYSTEM ACTION PANIC: %v, stack: %v\n", r, string(debug.Stack()))
			}
		}()
		f()
	}
}