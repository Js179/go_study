package basic

import "testing"

var (
	a1 = 1.0
	b1 = 2
	c1 = complex(2, 3)
	d1 = 3 + 2i
)

func Test_Num(t *testing.T) {
	t.Logf("c1 = %.0f", c1)
	// 复数 c1 = (2+3i)
}

func Test_Addition(t *testing.T) {
	// 语法错误，a是浮点数类型，b是整数类型，两者不能直接使用运算符
	// 但可以浮点数变量与整数常数做运算
	//  x := a + b
	// 需要类型转化一致后运算
	x := a1 + float64(b1)
	t.Logf("x type = %T, x = %.2f", x, x)
	// x type = float64, x = 3.00

	y := d1 + c1
	t.Logf("y type = %T, y = %.0f", y, y)
	// y type = complex128, y = (5+5i)
}

func Test_Subtraction(t *testing.T) {
	var x = uint(5 - a1)
	t.Logf("x type = %T, x = %v", x, x)
	// x type = uint, x = 4

	var y = uint(-5 - b1)
	t.Logf("y type = %T, y = %v", y, y)
	// y type = uint, y = 18446744073709551609

	var z = d1 - c1
	t.Logf("z type = %T, z = %v", z, z)
	// z type = complex128, z = (1-1i)
}

func Test_Multiply(t *testing.T) {
	// 浮点数变量可以与整数常数（不能是整数变量）做运算
	var x = a1 * 1
	var y = 4.5 * 2.8
	t.Logf("x type = %T, x = %.2f, x value = %v", x, x, x)
	// x type = float64, x = 1.00, x value = 1

	// %v 输出浮点数时，会省略后面的 0
	t.Logf("y type = %T, y = %v", y, y)
	// y type = float64, y = 12.6
}

func Test_Divide(t *testing.T) {
	x := 5 / 3
	y := 4.5 / 3
	// 此处 (3+1i) 不能写成 (3+i)
	z := (4 + 5i) / (3 + 1i)

	t.Logf("x type = %T, x = %v", x, x)
	// x type = int, x = 1

	t.Logf("y type = %T, y = %v", y, y)
	// y type = float64, y = 1.5

	t.Logf("z type = %T, z = %v", z, z)
	// z type = complex128, z = (1.7+1.1i)
}

func Test_Modulo(t *testing.T) {
	x := 5 % 3
	t.Logf("x type = %T, x = %v", x, x)
	// x type = int, x = 2
}
