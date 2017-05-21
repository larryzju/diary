package main

import(
   "fmt"
   "container/list"
)

func main() {
	lst := list.New()
	fmt.Println( lst )

	lst.PushBack( 1 )
	lst.PushBack( 2 )
	lst.PushBack( 4 )
	fmt.Println( lst )

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Printf( "%v\n", e.Value )
	}
}
