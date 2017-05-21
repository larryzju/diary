package main

import "fmt"

func max( l []int ) (max int) {
	max = l[0]
	for _, v := range l {
		if max < v {
			max = v
		}
	}
	return
}

func main() {
	fmt.Println( max( []int{1,2,3,4,-1,2,9} ) )
}
