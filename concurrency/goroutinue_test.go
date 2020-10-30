package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestUsingGoroutinue(t*testing.T){
	//设置运行时只有一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutinues")

	//使用go关键字开启协程
	//打应小写字母a~z
	go func(){
		defer wg.Done()

		for char:='a';char<'a'+26;char++{
			fmt.Printf("%c ",char)
		}
		fmt.Printf("\n")
	}()

	//打印大写字母A~Z
	go func(){
		defer wg.Done()
		for char:='A';char<'A'+26;char++{
			fmt.Printf("%c ",char)
		}
		fmt.Printf("\n")
	}()

	fmt.Println("Waiting Finished")
	wg.Wait()

	fmt.Println("program terminated")


	/*Start Goroutinues
	Waiting Finished
	A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	a b c d e f g h i j k l m n o p q r s t u v w x y z
	program terminated*/
}
