package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestChannel(t *testing.T) {
	wg.Add(2)
	court := make(chan int)
	go playBall("A",court)
	go playBall("B",court)

	court<-1
	wg.Wait()
}

func playBall(name string, court chan int) {
	defer wg.Done()
	for {
		//ok返回值标识chan是否关闭
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)//关闭chan
			return
		}
		fmt.Printf("Player %s hit %d\n",name, ball)
		ball++
		court <- ball
	}
}
