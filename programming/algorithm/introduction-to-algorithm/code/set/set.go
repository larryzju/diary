type Key int64

type Element interface{
	Key() Key
}

type Set interface {
	Search(k Key) Element
	Insert(x Element)
	Delete(x Element) bool
	Minimum() Element
	Maximum() Element
	Successor(x Element) Element
	Predecessor(x Element) Element
}
