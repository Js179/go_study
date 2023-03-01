package structure

import (
	"fmt"
	"testing"
)

type P interface {
	Call() bool
}

type Student struct {
	Age   uint
	Name  string
	Tel   string
	Email string
}

func Register(age uint, name, tel, email string) *Student {
	return &Student{
		Age:   age,
		Name:  name,
		Tel:   tel,
		Email: email,
	}
}

// 实现interface的函数
func (s *Student) Call() bool {
	fmt.Printf("call %s or send to %s", s.Tel, s.Email)
	return true
}

func (s *Student) String() string {
	return fmt.Sprintf("Age = %d,\tName = %s,\tTel=%s,\tEmail=%v",
		s.Age, s.Name, s.Tel, s.Email)
}

func Test_Method(t *testing.T) {
	// 实现P接口的全部函数，可以定义Student结构体(类似java的继承)
	var p P = Register(20, "ZS", "13422244422", "js131415@hotmail.com")
	// 类型强转
	t.Log(p.(*Student).String())

	p.Call()
}
