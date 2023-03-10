package ref

import (
	"reflect"
	"testing"
	"unsafe"
)

type user struct {
	Name string `json:"name,omitempty"`
	age  int
	pro  unsafe.Pointer
}

func Test_Ref_TypeOf(t *testing.T) {
	var u = user{"ZS", 20, nil}
	typeOf := reflect.TypeOf(u)
	t.Logf("type = %v, type name = %s, kind = %s, pkg = %v", typeOf, typeOf.Name(), typeOf.Kind(), typeOf.PkgPath())
	// type = ref.user, type name = user

	field := typeOf.Field(0)
	tag := field.Tag
	t.Logf("tag = %s", tag)
	// tag = json:"name,omitempty"

	value, ok := tag.Lookup("json")
	if ok {
		t.Log(value == tag.Get("json"))
		// name,omitempty
		t.Log(tag.Get("orm"))
		// ""
	}

	name, b := typeOf.FieldByName("age")
	if b {
		t.Logf("name = %s, type = %s", name.Name, name.Type)
		// name = age, type = int
	}

	_, b = typeOf.FieldByName("Age")
	t.Log(b)
	// false

	t.Log(typeOf.NumField())
	// 3

	align := typeOf.Align()
	t.Log(align)
	// 8
}

func TestTypeOf(t *testing.T) {
	var uArr = [1]user{}
	typeOf1 := reflect.TypeOf(uArr)
	t.Logf("type = %v, type name = %s, kind = %s, elem = %v, len = %d", typeOf1, typeOf1.Name(), typeOf1.Kind(), typeOf1.Elem(), typeOf1.Len())
	// type = [1]ref.user, type name = , kind = array, elem = ref.user, len = 1
	// Len() returns an array type's length. It panics if the type's Kind is not Array.
	// Elem() returns a type's element type. It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.

	var uSli = make([]user, 0)
	typeOf3 := reflect.TypeOf(uSli)
	t.Logf("type = %v, type name = %s, kind = %s, elem = %v", typeOf3, typeOf3.Name(), typeOf3.Kind(), typeOf3.Elem())
	// type = []ref.user, type name = , kind = slice, elem = ref.user
	// 切片类型使用Len函数会报错

	var i = 8
	typeOf2 := reflect.TypeOf(i)
	t.Logf("type = %v, type name = %s, kind = %s", typeOf2, typeOf2.Name(), typeOf2.Kind())
}

func TestIndirect(t *testing.T) {
	var u = user{"ZS", 20, nil}
	tmp := reflect.ValueOf(u)
	// 实例化一个新的对象，如果是nil返回0值
	t.Logf("tmp %p, u %p", &tmp, &u)
	// tmp 0xc000008108, u 0xc00004e440

	i := tmp.Interface()
	// 将当前对象类型转换为any
	t.Logf("tmp %p, i %p", &tmp, &i)
	// tmp 0xc000008108, i 0xc00004c2d0

	indirect := reflect.Indirect(tmp)
	// tmp为指针时返回指针类型指向的对象
	// tmp为nil pointer时返回0值
	// tmp不是指针时返回tmp
	t.Log(indirect)

	a := reflect.New(reflect.TypeOf(u)).Elem().Addr().Interface()
	t.Log(a)
}
