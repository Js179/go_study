package process

import "testing"

func Test_IF_ELSE(t *testing.T) {
	m[0] = "JS"
	if val, ok := m[2]; ok {
		t.Logf("key[2] is existed, val = %v", val)
	} else {
		t.Log("key[2] is not existed")
	}
	// key[2] is not existed

	it := 2
	for ; ; it++ {
		if it > 5 {
			t.Log("it is greater than 5")
			break
		} else if it == 4 {
			continue
		} else if it >= 3 {
			t.Log("it is greater than or equal 3")
		} else {
			t.Log("it is less than 3")
		}
	}
	/*	it is less than 3
		it is greater than or equal 3
		it is greater than or equal 3
		it is greater than 5*/
}
