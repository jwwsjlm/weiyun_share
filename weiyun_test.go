package weiyun

import (
	"testing"
)

func TestWeiyun_Json(t *testing.T) {
	t.Log("测试开始")
	//k为https://share.weiyun.com/******的最后一串字符
	w := NewWeiyun("*******")

	for i := 0; i < 5; i++ {
		json, err := w.WeiyunJson()
		if err != nil {
			t.Error(err)
		}
		t.Log(json)
	}
	t.Log("测试完成")
}
func TestWeiyun_Text(t *testing.T) {
	t.Log("测试开始")
	//k为https://share.weiyun.com/******的最后一串字符
	w := NewWeiyun("******")

	for i := 0; i < 5; i++ {
		json, err := w.WeiyunText()
		if err != nil {
			t.Error(err)
		}
		t.Log(json)
	}
	t.Log("测试完成")

}
