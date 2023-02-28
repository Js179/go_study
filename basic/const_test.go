package basic

import (
	"testing"
)

const (
	a, b = iota, iota
	c, d       // 未赋值，严格相仿上一行
	x    = "a" // 赋值，打破要求
	y          // 同时，下面的变量与上一行赋值的变量数据相同
	z
	e = iota // iota理解为行数，从0开始，此时e在第6行，值为5
)

func Test_Const(t *testing.T) {
	t.Logf("a = %d\tb = %d\nc = %d\td = %d\nx = %s\ty = %s\nz = %s\te = %d", a, b, c, d, x, y, z, e)
	/*	a = 0	b = 0
		c = 1	d = 1
		x = a	y = a
		z = a	e = 5*/
}
