package basic

import "testing"

func Test_Map(t *testing.T) {
	// 使用make初始化，size可以指定，也可以不指定
	// map1 := make(map[int]string)
	map1 := make(map[int]string, 3)
	map1[1] = "JS"
	map1[2] = "JZ"

	t.Logf("map1 = %v,\t, len = %d, map has not cap value", map1, len(map1))
	// map1 = map[1:JS 2:JZ],	, len = 2, map has not cap value

	val, ok := map1[3]
	/*	key不存在时，获取val并不会报错，而是返回val的类型0值(例如：int类型返回0，bool类型返回false)
		只有通过第二个返回参数ok(bool)来判断key是否存在*/
	t.Logf("key[3] is existed? %v", ok)
	// key[3] is existed? false
	t.Logf("key[3], value = %v", val)
	// key[3], value =

	if _, existed := map1[4]; existed {
		t.Log("key[4] is existed!")
	} else {
		t.Log("key[4] is not existed!")
	}
	// key[4] is not existed!
}
