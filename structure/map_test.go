package structure

import (
	"sync"
	"testing"
)

func Test_Map(t *testing.T) {
	var work sync.WaitGroup
	work.Add(2)

	m := make(map[int]string, 0)

	go func() {
		defer work.Done()
		m[1] = "2"
		m[2] = "22"

		t.Logf("go1 map = %v", m)
		delete(m, 1)
		t.Logf("after delete go1 map = %v", m)
	}()

	go func() {
		defer work.Done()
		m[1] = "23"
		m[2] = "24"

		t.Logf("go2 map = %v", m)
		delete(m, 1)
		t.Logf("after delete go2 map = %v", m)
	}()
	// 多协程操作map存在并发写入问题
	// fatal error: concurrent map writes

	work.Wait()
}

func Test_LockMap(t *testing.T) {
	var work sync.WaitGroup
	work.Add(3)

	var m LockMap[int, string]

	go func() {
		defer work.Done()
		m.Put(1, "21")
		m.Put(2, "22")

		t.Logf("go1 map = %v", m.data)
		m.Delete(1)
		t.Logf("after delete go1 map = %v", m.data)
	}()

	go func() {
		defer work.Done()
		m.Put(1, "23")
		m.Put(2, "24")

		t.Logf("go2 map = %v", m.data)
		m.Delete(1)
		t.Logf("after delete go2 map = %v", m.data)
	}()

	go func() {
		defer work.Done()
		m.Put(1, "25")
		m.Put(2, "26")

		t.Logf("go3 map = %v", m.data)
		m.Delete(1)
		t.Logf("after delete go3 map = %v", m.data)
	}()

	work.Wait()
	// lock map可以避免并发写入问题
}
