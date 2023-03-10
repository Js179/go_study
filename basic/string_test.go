package basic

import (
	"log"
	"strings"
	"testing"
	"time"
	"unsafe"
)

var str = "i am very hungry"

func safeEqual(a, b string) bool {
	now := time.Now()
	defer func() {
		log.Printf("the func takes %v", time.Since(now))
	}()
	if len(a) != len(b) {
		return false
	}
	var equal int = 0
	for i := 0; i < len(a); i++ {
		equal |= int(a[i]) ^ int(b[i])
	}
	return equal == 0
}

func Test_String(t *testing.T) {
	t.Logf("%s %T %#v", str, str, str)
	// i am hungry string "i am hungry"
}

func Test_Util(t *testing.T) {
	t.Logf("str prefix = 'i'? %v", strings.HasPrefix(str, "i"))
	// str prefix = 'i'? true

	t.Logf("str suffix = 'ry'? %v", strings.HasSuffix(str, "ry"))
	// str suffix = 'ry'? true

	tmp := strings.Clone(str)
	t.Logf("tmp = equal str ? %v", strings.EqualFold(str, tmp))
	// tmp = equal str ? true

	t.Logf("first 'y' index = ? %d", strings.Index(str, "y"))
	// first 'y' index = ? 8
	t.Logf("last 'y' index = ? %d", strings.LastIndex(str, "y"))
	// last 'y' index = ? 15

	res := strings.Split(str, " ")
	t.Logf("res type = %T, len = %d, value = %#v", res, len(res), res)
	// res type = []string, len = 4, value = []string{"i", "am", "very", "hungry"}

	// 定长分割
	resN := strings.SplitN(str, " ", 3)
	t.Logf("resN type = %T, len = %d, value = %#v", resN, len(resN), resN)
	// resN type = []string, len = 3, value = []string{"i", "am", "very hungry"}

	join := strings.Join(res, ",")
	t.Logf("join type = %T, len = %d, value = %#v", join, len(join), join)
	// join type = string, len = 16, value = "i,am,very,hungry"
}

func Test_Builder(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("a")
	builder.Write([]byte("b"))
	builder.WriteRune([]rune("中")[0])

	t.Logf("str value = %s", builder.String())
	// str value = ab中

	t.Logf("builder len = %d", builder.Len())
	// builder len = 5  a、b等各占1个字节 汉字'中'占3个字节

	t.Logf("before grow = %d\t", builder.Cap())
	// before grow = 8
	builder.Grow(3)
	t.Logf("after grow = %d\t", builder.Cap())
	// after grow = 8
	// 此时增长的n+len()<=原cap()[3+5<=8]，cap不会进行扩容，依然是8

	builder.Grow(10)
	t.Logf("after grow = %d\n", builder.Cap())
	// after = 26 计算: before * 2 + n 即 26 = 8 * 2 + 10

	// 清空
	builder.Reset()
	t.Logf("str value = %s,\t len = %d,\t cap = %d", builder.String(), builder.Len(), builder.Cap())
	// str value = ,	 len = 0,	 cap = 0
}

func Test_SafeEqual(t *testing.T) {
	var str1, str2 = "ses中问222222222222222222222也也", "ses中问222222222222222222222也也"
	t.Logf("val = %v", safeEqual(str1, str2))
}

func TestRune(t *testing.T) {
	var a rune = 2
	t.Logf("size = %d", unsafe.Sizeof(a))

	var str = "Go语言"
	t.Logf("size  = %d, len = %d", unsafe.Sizeof(str), len(str))

	arr := []rune(str)

	t.Logf("size  = %d, len = %d", unsafe.Sizeof(arr), len(arr))
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestStringSlice(t *testing.T) {
	var a, b []string
	a = append(a, "1,", "2", "3")
	b = append(b, "1,", "2", "31")

	equal := StringSliceEqual(a, b)
	t.Log(equal)
}
