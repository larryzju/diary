package main

import "fmt"

func Plus2() func(int)int {
	return func( x int ) int { return x + 2 }
}

func PlusX( v int ) func(int) int {
	return func( x int ) int { return x + v }
}

func main() {
	p2 := Plus2()
	fmt.Println( p2(10) )

	p5 := PlusX(5)
	fmt.Println( p5(10) )
}

