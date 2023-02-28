package process

import "testing"

var (
	str = "abcd上下左右"
	arr = [...]string{"a", "b", "c", "d", "上", "下", "左", "右"}
	m   = make(map[int]string, 0)
)

func Test_ForStr(t *testing.T) {
	for i := 0; i < len(str); i++ {
		t.Logf("str[%d] = %v， ascii = %v  \n", i, string(str[i]), str[i])
	}
	// 由于中文占3个字符，所以在遍历时3为的字符被分成一个一个的输出，导致乱码
	/*	str[0] = a， ascii = 97
		str[1] = b， ascii = 98
		str[2] = c， ascii = 99
		str[3] = d， ascii = 100
		str[4] = ä， ascii = 228
		str[5] = ¸， ascii = 184
		str[6] = ， ascii = 138
		str[7] = ä， ascii = 228
		str[8] = ¸， ascii = 184
		str[9] = ， ascii = 139
		str[10] = å， ascii = 229
		str[11] = ·， ascii = 183
		str[12] = ¦， ascii = 166
		str[13] = å， ascii = 229
		str[14] = ， ascii = 143
		str[15] = ³， ascii = 179*/

	// 解决方案，使用rune
	for i, v := range []rune(str) {
		t.Logf("index = %d,\t val = %v\n", i, string(v))
	}
	/*	index = 0,	 val = a
		index = 1,	 val = b
		index = 2,	 val = c
		index = 3,	 val = d
		index = 4,	 val = 上
		index = 5,	 val = 下
		index = 6,	 val = 左
		index = 7,	 val = 右*/

	// 无条件死循环
	/*	for {

		}*/

	// 类型C语言的while
	// 因此，golang中没有while，golang只有for类型的多种循环
	i := 0
	for i < 5 {
		i++
	}
}

func Test_ForArr(t *testing.T) {
	// for range, 出参有两个，分别为数组的index，对应的value
	for index, val := range arr {
		t.Logf("index = %d,\t val = %v", index, val)
	}
	/*	index = 0,	 val = a
		index = 1,	 val = b
		index = 2,	 val = c
		index = 3,	 val = d
		index = 4,	 val = 上
		index = 5,	 val = 下
		index = 6,	 val = 左
		index = 7,	 val = 右*/
}

func Test_ForMap(t *testing.T) {
	m[1] = "JS"
	m[2] = "JZ"
	m[3] = "LX"
	m[4] = "YY"
	m[5] = "CC"
	// 出参为key，val
	for k, v := range m {
		t.Logf("key = %v,\t val = %v", k, v)
	}
	/*	key = 1,	 val = JS
		key = 2,	 val = JZ
		key = 3,	 val = LX
		key = 4,	 val = YY
		key = 5,	 val = CC*/
}
