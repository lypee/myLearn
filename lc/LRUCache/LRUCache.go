package LRUCache

import "container/list"

// lc 146 LRU缓存

type LRUCache struct {
	lMap  map[int]Node
	lList []Node
	cap   int
	list  list.List
}

type Node struct {
	key   int
	value int

}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lMap:  make(map[int]Node, 0),
		lList: make([]Node, 0, capacity),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.lMap[key]
	if !exist {
		return -1
	}
	index := 0
	for i := range this.lList {
		if this.lList[i] == node {
			break
		}
	}
	this.lList[index], this.lList[0] = this.lList[0], this.lList[index]
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	newNode := Node{
		key:   key,
		value: value,
	}
	oldVal := this.Get(key)
	if oldVal == -1 { //不存在旧node
		if len(this.lList) > this.cap { //map已满
			newList := make([]Node, len(this.lList)+1)
			copy(newList[1:], this.lList)
			newList[0] = newNode
			this.lList = newList
			this.lMap[key] = newNode
			return
		}
		// 未超过cap
		newList := make([]Node, len(this.lList)+1)
		copy(newList[1:], this.lList)
		newList[0] = newNode

		this.lMap[key] = newNode
	} else { //旧node存在,更新lMap
		oldNode := Node{
			key:   key,
			value: oldVal,
		}
		this.lMap[key] = newNode
		idx := 0
		for i := range this.lList {
			if this.lList[i] == oldNode {
				this.lList[i].value = value
				idx = i
				break
			}
		}
		this.lList[idx], this.lList[0] = this.lList[0], this.lList[idx]
	}
	return
}
