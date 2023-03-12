package ref

import (
	"reflect"
	"testing"
)

var people = People{salary: 200000.2, age: 20, name: "ss", gender: Male}

func Test_No_Param_Method(t *testing.T) {
	t.Logf("people = %+v", people)

	newVal := reflect.ValueOf(people)
	//_ = newVal.Elem()  panic: reflect: call of reflect.Value.Elem on struct Value

	get := newVal.MethodByName("Get")
	get.Call(nil)

	// 当newVal是people的reflect.value时，调用*people的方法会报错
	//salary := newVal.MethodByName("Salary")
	//salary.Call(nil) panic: reflect: call of reflect.Value.Call on zero Value [recovered]

	// 此时创建的newPro可以调用*people的方法
	newPro := reflect.ValueOf(&people)
	_ = newPro.Elem() // pass
	salary := newPro.MethodByName("Salary")
	salary.Call(nil)

	t.Logf("people = %+v", people)

	// 返回any类型的原值
	newVal.Interface()
	// 可以强制类型转换，但转换的类型不完全符合会panic
	_ = newPro.Interface().(*People)
}

func Test_Has_Param_Method(t *testing.T) {
	t.Logf("people = %+v", people)

	newPro := reflect.ValueOf(&people)
	setParam := newPro.MethodByName("SetParam")

	param := []reflect.Value{reflect.ValueOf("ww")}
	setParam.Call(param)

	t.Logf("people = %+v", people)
}

func Test_Type(t *testing.T) {
	typeOf := reflect.TypeOf(people)
	get, _ := typeOf.MethodByName("Get")

	t.Log(get.IsExported()) // true

	t.Logf("type = %v", typeOf.Name())
	// type = People

	// 实例化 typeOf 类型的数据
	newVal := reflect.New(typeOf).Elem()
	t.Logf("%+v", newVal)
	// {age:0 name: gender:0 salary:0}

	if newVal.CanAddr() {
		// Addr 返回 数据v 的指针
		p := newVal.Addr().Interface().(*People)
		t.Log(p)
	}
}

func Test_Ref(t *testing.T) {
	// 类型js的深拷贝数据，无论people是值类型还是指针类型，最终返回reflect.Value类型的值数据
	indirect := reflect.Indirect(reflect.ValueOf(people))
	t.Logf("%+v, indirect p = %p, people p = %p", indirect, &indirect, &people)
	// indirect p = 0xc000008108, people p = 0xce2dc0

	// 创建一个string类型的reflect.slice 切片
	of := reflect.New(reflect.SliceOf(reflect.TypeOf("2"))).Elem()
	// 使用reflect.append增添数据
	of = reflect.Append(of, reflect.ValueOf("23"))
	t.Logf("val = %v, len = %d, cap = %d", of, of.Len(), of.Cap())
	// val = [23], len = 1, cap = 1
}
