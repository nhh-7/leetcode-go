package hot

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	- 不使用map
*	- 交错链表，将新节点插入到旧链表中对应节点的下一节点
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	dummy := &Node{}
	newCur := dummy
	for cur := head; cur != nil; cur = cur.Next {
		tmp := cur.Next
		cur.Next = tmp.Next
		newCur.Next = tmp
		newCur = newCur.Next
	}
	return dummy.Next
}

// 使用map存储新旧链表节点间的映射
func copyRandomList1(head *Node) *Node {
	m := make(map[*Node]*Node)

	var dummy *Node = &Node{}
	newListCur := dummy
	for cur := head; cur != nil; cur = cur.Next {
		newNode := &Node{
			Val: cur.Val,
		}
		newListCur.Next = newNode
		newListCur = newListCur.Next
		m[cur] = newNode
	}

	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Random = m[cur.Random]
	}
	return dummy.Next
}
