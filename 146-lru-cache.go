package leetcode

import "fmt"

/**
146. LRU缓存机制
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。



进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:
*/

//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ )
//cache.Put(1, 1);
//cache.Put(2, 2);
//cache.Get(1);       // 返回  1
//cache.Put(3, 3);    // 该操作会使得密钥 2 作废
//cache.Get(2);       // 返回 -1 (未找到)
//cache.Put(4, 4);    // 该操作会使得密钥 1 作废
//cache.Get(1);       // 返回 -1 (未找到)
//cache.Get(3);       // 返回  3
//cache.Get(4);       // 返回  4
type LRUCache struct {
	cap    int
	header *LruCacheNode
	tail   *LruCacheNode
	m      map[int]*LruCacheNode
}

type LruCacheNode struct {
	key   int
	value int
	pre   *LruCacheNode
	next  *LruCacheNode
}

func (l *LruCacheNode) remove() *LruCacheNode {
	l.pre.next, l.next.pre = l.next, l.pre
	return l
}

func (l *LruCacheNode) Print(str string) {
	result := []int{}
	for node := l; node != nil; node = node.next {
		result = append(result, node.value)
	}
	fmt.Print(str, result, "\n")
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &LruCacheNode{},
		tail:   &LruCacheNode{},
		m:      make(map[int]*LruCacheNode, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (l *LRUCache) Get(key int) int {
	if node, exists := l.m[key]; !exists {
		return -1
	} else {
		defer l.header.Print("Get: ")
		l.putHeader(node.remove())
		return node.value
	}
}

func (l *LRUCache) Put(key int, value int) {
	defer l.header.Print("PUT: ")

	if v, exists := l.m[key]; exists {
		v.value = value
		l.putHeader(v.remove())
		return
	}
	if l.cap <= len(l.m) {
		delete(l.m, l.tail.pre.key)
		l.tail.pre.remove()
	}
	newTailNode := &LruCacheNode{
		key:   key,
		value: value,
		pre:   l.tail,
	}
	l.m[key] = newTailNode
	l.putHeader(newTailNode)
	return
}

//l.header.next 更新
//node.pre 更新
//node.next 更新
//
//l.header.next.pre 更新

func (l *LRUCache) putHeader(node *LruCacheNode) {
	l.header.next.pre, node.next = node, l.header.next
	l.header.next = node
	node.pre = l.header
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Gut(key);
 * obj.Put(key,value);
 */

/**
map 双向链表 哨兵节点
变化:
cap ->
	初始化使用;
	判断容量 使用;
header ->
	put 超出 cap 变化
tail ->
	put & k 不存在
m ->
	put & k 存在 ->
		更新 v
	put & k 不存在 ->
		len(m) < cap ->
			写入 node
		len(m) >= cap
			移除tail node
			写入 新 node
总结:
	容量够不够 & m[k] 存不存
	m[k] 存在 则容量肯定够 -> m[k] 不存在 容量够不够问题
拆分:
	put
		存在
			是否更新到头结点
			更新value
		不存在
			不超出容量
				头结点是否更新为最新的node
			超出容量
				移除尾结点
				头结点是否更新为最新的node
	get
		不存在
			是否返回-1
		存在
			是否返回正确值
			是否更新到头结点
基础模块:
	是否更新到头结点 ->
		l.header.next 更新
		node.pre 更新
		node.next 更新
		l.header.next.pre 更新
*/
