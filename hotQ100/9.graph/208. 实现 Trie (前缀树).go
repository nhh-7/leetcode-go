package graph

/*
type Node struct {
	next  [26]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.next[c] == nil {
			cur.next[c] = &Node{}
		}
		cur = cur.next[c]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, c := range prefix {
		c -= 'a'
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true
} */

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Node struct {
	isEnd bool
	next  [26]*Node
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.next[c] == nil {
			cur.next[c] = &Node{}
		}
		cur = cur.next[c]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isEnd == true
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, c := range prefix {
		c -= 'a'
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
