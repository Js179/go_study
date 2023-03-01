package structure

import (
	"testing"
	"time"
)

func Test_Generics(t *testing.T) {
	// golang泛型使用 [K comparable, V any]
	var m LockMap[string, any]
	m.Put("123", time.Now())
	m.Put("234", true)
	m.Put("345", 2+4i)

	if v, ok := m.Get("345"); ok {
		t.Logf("key[345] value = %v", v)
	}

	t.Logf("m len = %d", m.Len())

	m.Delete("2")
}
