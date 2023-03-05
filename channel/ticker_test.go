package channel

import (
	"testing"
	"time"
)

func Test_Ticker(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case c := <-ticker.C:
				t.Logf("%v", c.Local())
			}
		}
	}()
	/*	2023-03-05 14:58:22.3544164 +0800 CST
		2023-03-05 14:58:22.8603471 +0800 CST
		2023-03-05 14:58:23.3657383 +0800 CST
		2023-03-05 14:58:23.8559753 +0800 CST*/

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
}
