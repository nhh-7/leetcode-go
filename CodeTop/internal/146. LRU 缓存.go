package internal

type CacheNode struct {
	prev, next *CacheNode
	key, value int
}

type LRUCache struct {
	capacity   int
	head, tail *CacheNode
	keyToNode  map[int]*CacheNode
}

func _Constructor(capacity int) LRUCache {
	head, tail := &CacheNode{}, &CacheNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity:  capacity,
		head:      head,
		tail:      tail,
		keyToNode: make(map[int]*CacheNode),
	}
}

func (lc *LRUCache) RemoveNode(node *CacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(lc.keyToNode, node.key)
}

func (lc *LRUCache) PushFront(node *CacheNode) {
	node.prev = lc.head
	node.next = lc.head.next

	node.prev.next = node
	node.next.prev = node
	lc.keyToNode[node.key] = node
}

func (lc *LRUCache) Get(key int) int {
	if _, ok := lc.keyToNode[key]; !ok {
		return -1
	}
	node := lc.keyToNode[key]
	lc.RemoveNode(node)
	lc.PushFront(node)
	return node.value
}

func (lc *LRUCache) Put(key int, value int) {
	if _, ok := lc.keyToNode[key]; !ok {
		if len(lc.keyToNode) == lc.capacity {
			lc.RemoveNode(lc.tail.prev)
		}
		newNode := &CacheNode{key: key, value: value}
		lc.PushFront(newNode)
	} else {
		node := lc.keyToNode[key]
		node.value = value
		lc.RemoveNode(node)
		lc.PushFront(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
