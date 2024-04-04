package utils

import "testing"

func BenchmarkGetARandomv2(b *testing.B) {
	digit := 9
	// 用于存储随机数的变量

	b.ResetTimer()             // 重置计时器，以便在测试开始时开始计时
	for i := 0; i < b.N; i++ { // 运行b.N次
		GetARandom(digit)
	}

}
func TestGetARandom(t *testing.T) {
	t.Log(GetARandom(5))
}
