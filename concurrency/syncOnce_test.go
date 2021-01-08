package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {

}

var singletonInstance *Singleton
var once sync.Once


func TestSyncOnce(t*testing.T){
	wg.Add(5)
	for i:=0;i<5;i++{
		go func() {
			defer wg.Done()
			instance := getSingletonInstance()
			fmt.Println(unsafe.Pointer(instance))
		}()
	}
	wg.Wait()
	fmt.Println("end!")
}

//单例延时加载
func getSingletonInstance() *Singleton{

	once.Do(func(){
		fmt.Println("create instance")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}
