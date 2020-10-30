package concurrency

import (
	"log"
	"os"
	"testing"
	"time"
)

// timeout规定了必须在多少秒内处理完成
const timeout = 10 * time.Second

func TestRunnerClient(t *testing.T) {
	// 为本次执行分配超时时间
	r := New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

}

// createTask返回一个根据id
// 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
