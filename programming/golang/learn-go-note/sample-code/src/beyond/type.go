package main

import "fmt"

type Point struct{
	x float64
	y float64
}

func main() {
	p := &Point{x:10}
	fmt.Printf( "%v %v %T\n", p, *p, p )
}
