package process

import "testing"

func Test_Switch(t *testing.T) {
	var i = 20

	/*	golang的switch中的每个case不需要再使用break
		可以使用fallthrough继续执行下面的case
	*/
	switch i {
	case 1:
		t.Log("i val = 1")
		//fallthrough
	case 20:
		t.Log("i val  = 20")
	default:
		t.Logf("nil")
	}
	// i val  = 20
}

func Test_Switch_Type(t *testing.T) {
	var k any = float32(2.2)

	switch k.(type) {
	case int:
		t.Log("k is a int")
	case float64:
		t.Log("k is a float64")
	default:
		t.Logf("k type is %T", k)
	}
	// k type is float32
}
