package channel

import (
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {
	now := time.Now()
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	t.Logf("it takes %v", time.Since(now))
	// it takes 2.0130685s
}

func Test_Timer2(t *testing.T) {
	now := time.Now()
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
	}()

	// 停止
	timer.Stop()
	t.Logf("it takes %v", time.Since(now))
	// it takes 0s timer被停止了
}
