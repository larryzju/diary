package main

import "fmt"

func Map( f func(int)int, xz []int ) []int {
	res := make( []int, len(xz) )
	for i, v := range xz {
		res[i] = f(v)
	}
	return res
}

func main() {
	fmt.Println( Map( func(x int)int{return x*x}, []int{1,2,3,4,5,6} ) )
}
