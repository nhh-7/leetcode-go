package hot

type CacheNode struct {
	key, value int
	prev, next *CacheNode
}

type LRUCache struct {
	head, tail *CacheNode
	capacity   int
	keyToNode  map[int]*CacheNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &CacheNode{}, &CacheNode{}
	head.next = tail
	tail.prev = head
	lc := LRUCache{
		head:      head,
		tail:      tail,
		capacity:  capacity,
		keyToNode: make(map[int]*CacheNode),
	}
	return lc
}

func (lc *LRUCache) remove(x *CacheNode) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

func (lc *LRUCache) pushFront(x *CacheNode) {
	x.prev = lc.head
	x.next = lc.head.next

	x.prev.next = x
	x.next.prev = x
}

func (lc *LRUCache) Get(key int) int {
	node, ok := lc.keyToNode[key]
	if !ok {
		return -1
	}
	lc.remove(node)
	lc.pushFront(node)
	return node.value
}

func (lc *LRUCache) Put(key int, value int) {
	node, ok := lc.keyToNode[key]
	if ok {
		node.value = value
		lc.remove(node)
		lc.pushFront(node)
		return
	}
	if len(lc.keyToNode) >= lc.capacity {
		r := lc.tail.prev
		lc.remove(r)
		delete(lc.keyToNode, r.key)
	}
	newNode := &CacheNode{
		key:   key,
		value: value,
	}
	lc.pushFront(newNode)
	lc.keyToNode[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
