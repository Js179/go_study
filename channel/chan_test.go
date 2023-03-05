package channel

import (
	"log"
	"testing"
	"time"
)

func Test_Chan_NoBuffer(t *testing.T) {
	// 无缓冲通道
	c := make(chan int)

	// 默认发送和接收操作是阻塞的
	// 使用协程并发通道的读写操作
	go func() {
		c <- 2
	}()

	// 当通道读时，协程的写操作被唤醒
	t.Log(<-c)
	// 因此，无缓冲的通道如果不使用并发处理读写操作则会报错
}

func Test_Chan_Buffer(t *testing.T) {
	// 缓存通道
	c := make(chan string, 2)

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中
	c <- "msg"
	c <- "data"

	t.Log(<-c, <-c)
}

func worker(done chan bool) {
	log.Printf("work starting")
	time.Sleep(time.Second)
	log.Printf("work finish")
	done <- true
}

func Test_Sync(t *testing.T) {
	var done = make(chan bool)

	go worker(done)

	// 利用通道的堵塞特性实现同步等待运行
	<-done
	t.Log("end")
}

func run[T any](c chan T, v T) {
	time.Sleep(time.Second)
	c <- v
}

func Test_Select(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go run(c1, "msg")
	go run(c2, "data")
	go run(c1, "ok")
	go run(c2, "status")

	for i := 0; i < 4; i++ {
		// select 通道选择器
		select {
		case msg := <-c1:
			t.Logf("chan c1 val = %v", msg)
		case msg := <-c2:
			t.Logf("chan c2 val = %v", msg)
		}
	}
}

func Test_After(t *testing.T) {
	c1 := make(chan string)

	go run(c1, "ok")

	end := true
	for end {
		select {
		case msg := <-c1:
			t.Logf("chan c1 val = %v", msg)
		case <-time.After(2 * time.Second):
			end = false
		}
	}
	// 在select中，超过2s后会触发第二个case选项，从而结束循环
	t.Log("超时结束")
	/*	=== RUN   Test_After
		chan_test.go:86: chan c1 val = ok
		chan_test.go:91: 超时结束
		--- PASS: Test_After (3.01s)*/
}

func Test_Chan_Close(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string, 2)

	c2 <- "msg3"
	c2 <- "msg2"
	// 关闭c2通道
	close(c2)
	go run(c1, "finish")

	t.Logf("c1 val = %v", <-c1)

	// 缓存通道close后仍然可以读取
	for v := range c2 {
		t.Logf("c2 val = %v", v)
	}
	/*	c1 val = finish
		c2 val = msg3
		c2 val = msg2*/
}
