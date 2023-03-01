package structure

import "testing"

func Test_Error(t *testing.T) {
	defer func() {
		t.Log("end")
	}()
	// 抛出异常，程序异常终止，但会保证defer运行完
	panic("error")
}

func Test_Recover(t *testing.T) {
	defer func() {
		if a := recover(); a != nil {
			t.Logf("recover success, val = %v", a)
		}
	}()
	panic(0)
	// recover success, val = 0 程序正常结束
	// 利用defer来捕获panic，防止程序异常终止
}
