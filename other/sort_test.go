package other

import (
	"sort"
	"testing"
)

func Test_Sort(t *testing.T) {
	arr := []int{2, 5, 1, 5, 6, 3}
	sort.Ints(arr)
	t.Log(arr)
	// [1 2 3 5 5 6]

	strAres := []string{"arr", "hello", "world", "make"}
	sort.Strings(strAres)
	t.Log(strAres)
	// [arr hello make world]
}

// 实现 sort.Interface 接口的 Len、Less 和 Swap 方法
type strArr []string

// Len、Swap的实现基本一致
func (a strArr) Len() int {
	return len(a)
}
func (a strArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 自定义排序逻辑
func (a strArr) Less(i, j int) bool {
	// 首字母降序
	if x, y := len(a[i]), len(a[j]); x > 0 && y > 0 {
		return a[i][0] > a[j][0]
	} else if x == 0 {
		return false
	}
	return true
}

func Test_Sort_By_Func(t *testing.T) {
	strAres := strArr{"arr", "", "hello", "world", "make", ""}
	sort.Sort(strAres)
	t.Logf("%#v", strAres)
	// other.strArr{"world", "make", "hello", "arr", "", ""}
}
