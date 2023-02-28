package basic

import (
	"testing"
)

func Test_Variable(t *testing.T) {
	// 变量的多种定义方式
	var a = "xyz"
	b := 1
	var (
		x string
		y int
		z bool
		n any
	)
	t.Logf("a = %s\tb = %d\nx = %s\ty = %v\nz = %+v\tn = %#v", a, b, x, y, z, n)
	/*	a = xyz	  b = 1
		x = 	  y = 0
		z = false	n = <nil>*/
}
