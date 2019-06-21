package list

type DoubleLinkedListNode struct {
	value Element
	next  *DoubleLinkedListNode
	prev  *DoubleLinkedListNode
}

func (n *DoubleLinkedListNode) Element() Element {
	return n.value
}

type DoubleLinkedList struct {
	Head *DoubleLinkedListNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	head := &DoubleLinkedListNode{value: nil, next: nil, prev: nil}
	return &DoubleLinkedList{Head: head}
}

func (lst *DoubleLinkedList) Search(e Element) Node {
	p := lst.Head.next
	for p != nil && p.value != e {
		p = p.next
	}
	return p
}

func (lst *DoubleLinkedList) InsertAfter(node Node, e Element) Node {
	n := node.(*DoubleLinkedListNode)
	en := &DoubleLinkedListNode{e, n.next, n}
	n.next = en
	return en
}

func (lst *DoubleLinkedListNode) Delete(node Node) Element {
	n := node.(*DoubleLinkedListNode)
	n.prev.next = n.next
	n.next.prev = n.prev
	return n.value
}
