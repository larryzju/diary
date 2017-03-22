package main

import "fmt"

func fibnacci( n int ) []uint64 {
	xs := make( []uint64, n )
	for s,b,i := uint64(0),uint64(1),0; i < n; i,s,b = i+1,b,s+b {
		xs[i] = s
	}
	return xs
}

func main() {
	fmt.Println( fibnacci( 20 ) )
}
