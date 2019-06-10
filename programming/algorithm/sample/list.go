package sample

type List struct {
	v interface{}
	next *List
}


func Reverse(lst *List) *List{
	var p, q *List = nil, lst
	for q != nil {
		p, q.next, q = q, p, q.next
	}
	return p
}


