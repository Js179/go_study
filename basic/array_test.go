package basic

import "testing"

func change(arr [3]int) {
	arr[1] = 2
}

func changePr(arr *[3]int) {
	arr[1] = 2
}

func changeSlice(sli []string) {
	sli[0] = "2"
}

func Test_Array(t *testing.T) {
	arr := [3]int{}
	t.Logf("arr = %v,\t len = %d,\t cap = %d", arr, len(arr), cap(arr))
	// arr = [0 0 0],	 len = 3,	 cap = 3

	change(arr)
	t.Logf("arr = %v,\t len = %d,\t cap = %d", arr, len(arr), cap(arr))
	// arr = [0 0 0],	 len = 3,	 cap = 3
	// 数组作为参数时，使用的是副本，不会改变原数组数据

	changePr(&arr)
	t.Logf("arr = %v,\t len = %d,\t cap = %d", arr, len(arr), cap(arr))
	// arr = [0 2 0],	 len = 3,	 cap = 3

	// change的参数是[3]int类型，[4]int类型不能作为其参数
	// arr4 := [4]int{}
	// change(arr4)
	// 初始化自适应长度数组
	arr3 := [...]int{}
	t.Logf("arr type = %T,\t arr3 type = %T", arr, arr3)
	// arr type = [3]int,	 arr3 type = [0]int
}

func Test_Slice(t *testing.T) {
	var sl []int
	t.Logf("make slice sl = %v", sl)
	// make slice sl = []

	/* 推荐创建方式 */
	sli1, sli2 := make([]int, 0), make([]string, 1)
	t.Logf("sli1 = %#v,\t len = %d,\t cap = %d\n", sli1, len(sli1), cap(sli1))
	// sli1 = []int{},	 len = 0,	 cap = 0

	t.Logf("sli2 = %#v,\t len = %d,\t cap = %d\n", sli2, len(sli2), cap(sli2))
	// sli2 = []string{""},	 len = 1,	 cap = 1

	changeSlice(sli2)
	// 切片slice作为函数参数时，传递的是指针，毕竟切片的底层就是数组指针
	t.Logf("sli2 = %#v,\t len = %d,\t cap = %d\n", sli2, len(sli2), cap(sli2))
	// sli2 = []string{"2"},	 len = 1,	 cap = 1
}

func Test_Slice2(t *testing.T) {
	sli := make([]int, 5, 6)
	t.Logf("sli = %v,\t len = %d,\t cap = %d\n", sli, len(sli), cap(sli))
	// sli = [0 0 0 0 0],	 len = 5,	 cap = 6

	// 填充数据
	sli = append(sli, []int{1, 23, 4, 5}...)
	t.Logf("sli = %v,\t len = %d,\t cap = %d\n", sli, len(sli), cap(sli))
	/*	sli = [0 0 0 0 0 1 23 4 5],	 len = 9,	 cap = 12
		切片初始化了5个数据 0
		使用append填充数据后，容量cap不足，会进行自增
		自增逻辑为：当前容量*2；如果当前容量>1024,则自增25%*/

	/*	切[n:m:c]
		获取新切片，切片区间[n,m),第3个参数计算新切片的cap
		新切片len = m - n
		新切片cap = c - n
		n省略时为0
		m省略时为原切片的len
		c省略时为原切片的cap*/
	sli1 := sli[1:]
	t.Logf("sli1 = %v,\t len = %d,\t cap = %d\n", sli1, len(sli1), cap(sli1))
	// sli1 = [0 0 0 0 1 23 4 5],	 len = 8,	 cap = 11

	sli2 := sli[:6]
	t.Logf("sli2 = %v,\t len = %d,\t cap = %d\n", sli2, len(sli2), cap(sli2))
	// sli2 = [0 0 0 0 0 1],	 len = 6,	 cap = 12

	sli3 := sli1[3:5:6]
	t.Logf("sli3 = %v,\t len = %d,\t cap = %d\n", sli3, len(sli3), cap(sli3))
	// sli2 = [0 1],	 len = 2,	 cap = 3
}
