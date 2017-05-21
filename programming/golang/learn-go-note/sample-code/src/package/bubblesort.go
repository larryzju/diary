// unfinished
package main

import "fmt"

type Comparable interface{}

func bubblesort( n []Comparable ) {
	for i := 0; i < len(n)-1; i++ {
		for j := i+1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func main() {
	s := []int{1,2,3,4,-1,2}
	bubblesort( s )
	fmt.Println( s )
}
