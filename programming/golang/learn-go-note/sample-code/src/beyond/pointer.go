package main

import "fmt"

func main() {
	var p *int
	fmt.Printf( "%v %T\n", p, p )

	var i int
	p = &i
	fmt.Printf( "%v %v %T\n", p, *p, p )
}
