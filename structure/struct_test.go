package structure

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type People struct {
	name string
	Age  uint
}

type Teacher struct {
	Name     string `json:"name"` // tag 标签
	Age      uint   `json:"age,string"`
	WorkTime uint8  `json:"workTime,omitempty"` // omitempty 序列化时忽略0值或空值
}

func (p People) SetName(name string) {
	p.name = name
}

func (p *People) SetAge(age uint) {
	p.Age = age
}

func Test_Struct(t *testing.T) {
	p := People{}
	p.SetName("JS")
	// SetName的receiver是结构体p，而不是其指针，不会修改原结构体数据
	t.Logf("p = %+v", p)
	// p = {name: age:0}

	(&p).SetAge(20)
	// 结构体可以直接修改属性值
	p.name = "JS"
	t.Logf("p = %+v", p)
	// p = {name:JS age:20}
}

func Test_Struct_JSON(t *testing.T) {
	teacher := Teacher{
		Age:      23,
		Name:     "ZW",
		WorkTime: 3,
	}

	bytes, err := json.Marshal(teacher)
	// json 解析成功
	if err == nil {
		t.Logf("struct json value = %s", string(bytes))
		// {"name":"ZW","age":"23","workTime":3} 此时age序列化为字符串
	}

	// new创建结构体指针
	tea := new(Teacher)
	err = json.Unmarshal(bytes, tea)
	if err == nil {
		t.Logf("struct value = %+v", *tea)
		// struct value = {Name:ZW Age:23 WorkTime:3}
	}

	p := People{
		Age:  20,
		name: "WS",
	}
	pJSON, err := json.Marshal(p)
	if err == nil {
		t.Logf("struct json value = %s", string(pJSON))
		// struct json value = {"Age":20}
		// 由于People结构体中的name属性是小写字母开头，对外不可见
	}

	zs := Teacher{
		Age:  20,
		Name: "ZS",
	}

	zJSON, err := json.Marshal(zs)
	if err == nil {
		t.Logf("struct json value = %s", string(zJSON))
		// struct json value = {"name":"ZS","age":"20"} WorkTime为0值(空值),已被忽略
	}
}

func Test_JSONITER(t *testing.T) {
	// jsoniter用法与encoding/json完全一致
	var jsonUtil = jsoniter.ConfigCompatibleWithStandardLibrary

	teacher := Teacher{
		Age:      25,
		Name:     "XR",
		WorkTime: 3,
	}

	bytes, err := jsonUtil.Marshal(teacher)
	if err == nil {
		t.Logf("struct json value = %s", string(bytes))
		// struct json value = {"name":"XR","age":"25","workTime":3}
	}

	// 用于http连接与socket连接的读取与写入，或者文件读取

	// 将teacher的数据写入到write流中
	/*	var write io.Writer
		err := jsonUtil.NewEncoder(write).Encode(teacher)*/

	// 将read流中的数据反序列化到p结构体中
	/*	var read io.Reader
		var p People
		err := jsonUtil.NewDecoder(read).Decode(&p)*/
}
