package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {

	hash := New(3, func(key []byte) uint32 {
		// hash算法依据key的数值
		// 使用自定义的简单的hash算法验证，只处理数字的key
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// 添加3次，总共9个点
	hash.Add("6", "4", "2")
	// 9个点形成一个环
	//06 16 26
	//04 14 24
	//02 12 22

	// key对应的节点的值
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s should have yielded %s", k, v)
		}
	}

	// 再增加3个点 08 18 28
	hash.Add("8")

	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s should have yielded %s", k, v)
		}
	}
}
