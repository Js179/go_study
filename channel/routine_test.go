package channel

import (
	"sync"
	"testing"
	"time"
)

func Test_Routine(t *testing.T) {
	// 使用go关键字开启协程
	go func() {
		t.Log("go runtime 1")
	}()

	go func() {
		t.Log("go runtime 2")
	}()

	// 如果不sleep等待，主main函数会直接结束，终止其他协程，可能会没有输出
	time.Sleep(time.Second)
}

func Test_WaitGroup(t *testing.T) {
	// 协程同步
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		t.Log("go runtime 1")
		wait.Done()
	}()

	go func() {
		t.Log("go runtime 2")
		wait.Done()
	}()

	// 等待其他协程执行结束
	wait.Wait()
}
