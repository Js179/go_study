package structure

import (
	"os"
	"testing"
)

func Test_Defer(t *testing.T) {
	defer func() {
		t.Log("end1")
	}()
	defer func() {
		t.Log("end2")
	}()
	t.Log("start...")
	t.Log("1,2,3")
	/*	start...
		1,2,3
		end2
		end1*/
	// defer会将待执行的语句放入栈中，在程序结束前再一一执行
	// 而栈是先进后出，所以执行顺序与入栈顺序是反着的
}

func Test_Exit(t *testing.T) {
	defer func() {
		t.Log("end")
	}()
	t.Log("start...")
	// start...
	// 直接结束程序，defer还没执行就结束了
	os.Exit(1)
}
