package list

type Element interface {
}

type Node interface {
	Element() Element
}

type List interface {
	Search(Element) Node
	InsertAfter(Node, Element) Node
	Delete(Node) Element
}
