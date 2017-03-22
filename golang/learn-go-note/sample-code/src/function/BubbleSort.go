package main

import "fmt"

// slices is a reference type
// modifies xs in place
func BubbleSort( xs []int ) {
	for i := 0; i < len(xs)-1; i++ {
		for j := i+1; j < len(xs); j++ {
			if xs[i] > xs[j] {
				xs[i], xs[j] = xs[j], xs[i]
			}
		}
	}
}

func main() {
	a := []int{6,4,5,3,2,1}
	fmt.Println( a )
	BubbleSort( a )
	fmt.Println( a )

	b := [...]int{5,7,3,8,1,4}
	c := b[:]
	fmt.Println( b )
	BubbleSort(c)
	fmt.Println( c )
}
