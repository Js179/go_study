package ref

import (
	"reflect"
	"testing"
	"unsafe"
)

type user struct {
	Name string
	age  int
	pro  unsafe.Pointer
}

func Test_Ref(t *testing.T) {
	var u = user{"zs", 23, nil}
	of := reflect.TypeOf(u)
	t.Logf("kind = %v, type = %v, name = %v", of.Kind(), of, of.Name())
	t.Logf("val = %v, pkg = %v", reflect.ValueOf(u), of.PkgPath())
}
