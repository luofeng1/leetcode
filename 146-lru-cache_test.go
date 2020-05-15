package leetcode

import "testing"

func TestLRUCache_Get(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 { // 返回  1
		t.Errorf("1. get err")
	}
	cache.Put(3, 3)         // 该操作会使得密钥 2 作废
	if cache.Get(2) != -1 { // 返回 -1 (未找到)
		t.Errorf("2. get err")
	}
	cache.Put(4, 4)         // 该操作会使得密钥 1 作废
	if cache.Get(1) != -1 { // 返回 -1 (未找到)
		t.Errorf("3. get err")
	}
	if cache.Get(3) != 3 { // 返回  3
		t.Errorf("4. get err")
	}
	if cache.Get(4) != 4 { // 返回  4
		t.Errorf("5. get err")
	}
}
